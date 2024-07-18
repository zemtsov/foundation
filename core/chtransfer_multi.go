package core

import (
	"encoding/json"
	"fmt"
	"strings"
	"unicode/utf8"

	"github.com/anoideaopen/foundation/core/cctransfer"
	"github.com/anoideaopen/foundation/core/types"
	"github.com/anoideaopen/foundation/core/types/big"
	pb "github.com/anoideaopen/foundation/proto"
	"github.com/golang/protobuf/proto"
)

type TransferItem struct {
	Token  string   `json:"token"`
	Amount *big.Int `json:"amount"`
}

// TxChannelMultiTransferByCustomer - transaction initiating multi transfer between channels.
// The owner of tokens signs the transaction. Tokens are transferred to themselves.
// After the checks, a multi transfer record is created and the user's balances are reduced.
func (bc *BaseContract) TxChannelMultiTransferByCustomer(
	sender *types.Sender,
	idTransfer string,
	to string,
	items []TransferItem,
) (string, error) {
	return bc.createCCMultiTransferFrom(idTransfer, to, sender.Address(), items)
}

// TxChannelMultiTransferByAdmin - transaction initiating multi transfer between channels.
// Signed by the channel admin (site). The tokens are transferred from idUser to the same user.
// After the checks, a multi transfer record is created and the user's balances are reduced.
func (bc *BaseContract) TxChannelMultiTransferByAdmin(
	sender *types.Sender,
	idTransfer string,
	to string,
	idUser *types.Address,
	items []TransferItem,
) (string, error) {
	// Checks
	if !bc.config.IsAdminSet() {
		return "", cctransfer.ErrAdminNotSet
	}

	if admin, err := types.AddrFromBase58Check(bc.config.GetAdmin().GetAddress()); err == nil {
		if !sender.Equal(admin) {
			return "", cctransfer.ErrUnauthorisedNotAdmin
		}
	} else {
		return "", fmt.Errorf("creating admin address: %w", err)
	}

	if sender.Equal(idUser) {
		return "", cctransfer.ErrInvalidIDUser
	}

	return bc.createCCMultiTransferFrom(idTransfer, to, idUser, items)
}

func (bc *BaseContract) createCCMultiTransferFrom(
	idTransfer string,
	to string,
	idUser *types.Address,
	items []TransferItem,
) (string, error) {
	if strings.EqualFold(bc.config.GetSymbol(), to) {
		return "", cctransfer.ErrInvalidChannel
	}

	if len(items) == 0 {
		return "", cctransfer.ErrEmptyTransferItems
	}

	symbol := tokenSymbol(items[0].Token)
	if !strings.EqualFold(bc.config.GetSymbol(), symbol) && !strings.EqualFold(to, symbol) {
		return "", cctransfer.ErrInvalidToken
	}

	var transferItems []*pb.CCTransferItem
	for _, item := range items {
		if symbol != tokenSymbol(item.Token) {
			return "", cctransfer.ErrInvalidToken
		}
		transferItems = append(transferItems, &pb.CCTransferItem{Token: item.Token, Amount: item.Amount.Bytes()})
	}

	stub := bc.GetStub()

	if _, err := cctransfer.LoadCCFromMultiTransfer(stub, idTransfer); err == nil {
		return "", cctransfer.ErrIDMultiTransferExist
	}

	ts, err := stub.GetTxTimestamp()
	if err != nil {
		return "", err
	}

	tr := &pb.CCMultiTransfer{
		Id:               idTransfer,
		From:             bc.config.GetSymbol(),
		To:               to,
		Items:            transferItems,
		User:             idUser.Bytes(),
		ForwardDirection: strings.EqualFold(bc.config.GetSymbol(), symbol),
		TimeAsNanos:      ts.AsTime().UnixNano(),
	}

	if err = cctransfer.SaveCCFromMultiTransfer(stub, tr); err != nil {
		return "", err
	}

	for _, item := range items {
		if err = bc.ccTransferChangeBalance(
			CreateFrom,
			tr.GetForwardDirection(),
			idUser,
			item.Amount,
			tr.GetFrom(),
			tr.GetTo(),
			item.Token,
		); err != nil {
			return "", err
		}
	}

	return stub.GetTxID(), nil
}

func (bc *BaseContract) TxCreateCCMultiTransferTo(dataIn string) (string, error) {
	var tr pb.CCMultiTransfer
	if err := proto.Unmarshal([]byte(dataIn), &tr); err != nil {
		if err = json.Unmarshal([]byte(dataIn), &tr); err != nil {
			return "", err
		}
	}

	stub := bc.GetStub()

	if _, err := cctransfer.LoadCCToMultiTransfer(stub, tr.GetId()); err == nil {
		return "", cctransfer.ErrIDMultiTransferExist
	}

	if !strings.EqualFold(bc.ContractConfig().GetSymbol(), tr.GetFrom()) && !strings.EqualFold(bc.ContractConfig().GetSymbol(), tr.GetTo()) {
		return "", cctransfer.ErrInvalidChannel
	}

	if strings.EqualFold(tr.GetFrom(), tr.GetTo()) {
		return "", cctransfer.ErrInvalidChannel
	}

	if len(tr.GetItems()) == 0 {
		return "", cctransfer.ErrEmptyTransferItems
	}

	symbol := tokenSymbol(tr.GetItems()[0].GetToken())

	if !strings.EqualFold(tr.GetFrom(), symbol) && !strings.EqualFold(tr.GetTo(), symbol) {
		return "", cctransfer.ErrInvalidToken
	}

	if strings.EqualFold(tr.GetFrom(), symbol) != tr.GetForwardDirection() {
		return "", cctransfer.ErrInvalidToken
	}

	for _, item := range tr.GetItems() {
		if symbol != tokenSymbol(item.GetToken()) {
			return "", cctransfer.ErrInvalidToken
		}
	}

	tr.IsCommit = true
	if err := cctransfer.SaveCCToMultiTransfer(bc.GetStub(), &tr); err != nil {
		return "", err
	}

	for _, item := range tr.GetItems() {
		err := bc.ccTransferChangeBalance(
			CreateTo,
			tr.GetForwardDirection(),
			types.AddrFromBytes(tr.GetUser()),
			new(big.Int).SetBytes(item.GetAmount()),
			tr.GetFrom(),
			tr.GetTo(),
			item.GetToken(),
		)
		if err != nil {
			return "", err
		}

	}

	return bc.GetStub().GetTxID(), nil
}

// TxCancelCCMultiTransferFrom - transaction cancels (deletes) the multi transfer record in the From channel
// returns balances to the user. If the service cannot create a response part in the "To" channel
// within some timeout, it is required to cancel the transfer.
// After TxChannelMultiTransferByAdmin or TxChannelMultiTransferByCustomer
// This transaction is sent only by the channel-transfer service with a "robot" certificate
func (bc *BaseContract) TxCancelCCMultiTransferFrom(id string) error {
	// see if it's already gone
	tr, err := cctransfer.LoadCCFromMultiTransfer(bc.GetStub(), id)
	if err != nil {
		return cctransfer.ErrNotFound
	}

	// If it's already committed, it's a mistake.
	if tr.GetIsCommit() {
		return cctransfer.ErrTransferCommit
	}

	// rebalancing
	for _, item := range tr.GetItems() {
		err = bc.ccTransferChangeBalance(
			CancelFrom,
			tr.GetForwardDirection(),
			types.AddrFromBytes(tr.GetUser()),
			new(big.Int).SetBytes(item.GetAmount()),
			tr.GetFrom(),
			tr.GetTo(),
			item.GetToken(),
		)
		if err != nil {
			return err
		}
	}

	return cctransfer.DelCCFromTransfer(bc.GetStub(), id)
}

// NBTxCommitCCMultiTransferFrom - transaction writes the commit flag in the multi transfer in the From channel.
// Executed after successful creation of a mating part in the channel To (TxCreateCCMTransferTo)
// This transaction is sent only by the channel-transfer service with a "robot" certificate
func (bc *BaseContract) NBTxCommitCCMultiTransferFrom(id string) error {
	// see if it's already gone
	tr, err := cctransfer.LoadCCFromMultiTransfer(bc.GetStub(), id)
	if err != nil {
		return cctransfer.ErrNotFound
	}

	// if it's already committed, it's an error
	if tr.GetIsCommit() {
		return cctransfer.ErrTransferCommit
	}

	tr.IsCommit = true
	return cctransfer.SaveCCFromMultiTransfer(bc.GetStub(), tr)
}

// NBTxDeleteCCMultiTransferFrom - transaction deletes the multi transfer record in the channel From.
// Performed after successful removal in the channel To (NBTxDeleteCCMultiTransferTo)
// This transaction is sent only by the channel-transfer service with a "robot" certificate
func (bc *BaseContract) NBTxDeleteCCMultiTransferFrom(id string) error {
	// see if it's already gone
	tr, err := cctransfer.LoadCCFromMultiTransfer(bc.GetStub(), id)
	if err != nil {
		return cctransfer.ErrNotFound
	}

	// if it's not committed, it's an error
	if !tr.GetIsCommit() {
		return cctransfer.ErrTransferNotCommit
	}

	return cctransfer.DelCCFromMultiTransfer(bc.GetStub(), id)
}

// NBTxDeleteCCMultiTransferTo - transaction deletes multi transfer record in channel To.
// Executed after a successful commit in the From channel (NBTxCommitCCTransferFrom)
// This transaction is sent only by the channel-transfer service with a "robot" certificate
func (bc *BaseContract) NBTxDeleteCCMultiTransferTo(id string) error {
	// Let's check if it's not already
	tr, err := cctransfer.LoadCCToMultiTransfer(bc.GetStub(), id)
	if err != nil {
		return cctransfer.ErrNotFound
	}

	// if it is not committed, error
	if !tr.GetIsCommit() {
		return cctransfer.ErrTransferNotCommit
	}

	return cctransfer.DelCCToMultiTransfer(bc.GetStub(), id)
}

// QueryChannelMultiTransferFrom - receiving a multi transfer record from the channel From
func (bc *BaseContract) QueryChannelMultiTransferFrom(id string) (*pb.CCMultiTransfer, error) {
	tr, err := cctransfer.LoadCCFromMultiTransfer(bc.GetStub(), id)
	if err != nil {
		return nil, err
	}

	return tr, nil
}

// QueryChannelMultiTransferTo - receiving a multi transfer record from the channel From
func (bc *BaseContract) QueryChannelMultiTransferTo(id string) (*pb.CCMultiTransfer, error) {
	tr, err := cctransfer.LoadCCToMultiTransfer(bc.GetStub(), id)
	if err != nil {
		return nil, err
	}

	return tr, nil
}

// QueryChannelMultiTransfersFrom - getting all multi transfer records from the channel From
// You can receive them in parts (chunks)
func (bc *BaseContract) QueryChannelMultiTransfersFrom(pageSize int64, bookmark string) (*pb.CCMultiTransfers, error) {
	if pageSize <= 0 {
		return nil, cctransfer.ErrPageSizeLessOrEqZero
	}

	prefix := cctransfer.CCFromMultiTransfers()
	startKey, endKey := prefix, prefix+string(utf8.MaxRune)

	if bookmark != "" && !strings.HasPrefix(bookmark, prefix) {
		return nil, cctransfer.ErrInvalidBookmark
	}

	trs, err := cctransfer.LoadCCFromMultiTransfers(bc.GetStub(), startKey, endKey, bookmark, int32(pageSize))
	if err != nil {
		return nil, err
	}

	return trs, nil
}
