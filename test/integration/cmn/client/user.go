package client

import (
	"crypto/ecdsa"
	"crypto/ed25519"
	"crypto/rand"
	"errors"
	"fmt"
	"strings"

	"github.com/anoideaopen/foundation/core/eth"
	pbfound "github.com/anoideaopen/foundation/proto"
	"github.com/btcsuite/btcutil/base58"
	"golang.org/x/crypto/sha3"
)

type UserFoundation struct {
	PrivateKeyBytes    []byte
	PublicKeyBytes     []byte
	PublicKeyType      string
	PublicKeyBase58    string
	AddressBase58Check string
	UserID             string
}

func NewUserFoundation(keyType string) *UserFoundation {
	var (
		privateKeyBytes []byte
		publicKeyBytes  []byte
		err             error
	)

	switch keyType {
	case pbfound.KeyType_ed25519.String():
		publicKeyBytes, privateKeyBytes, err = ed25519.GenerateKey(rand.Reader)
		if err != nil {
			return &UserFoundation{}
		}

	case pbfound.KeyType_secp256k1.String():
		var privateKey *ecdsa.PrivateKey
		privateKey, err = eth.NewKey()
		if err != nil {
			return &UserFoundation{}
		}
		privateKeyBytes = eth.PrivateKeyBytes(privateKey)
		publicKeyBytes = eth.PublicKeyBytes(&privateKey.PublicKey)

	default:
		return &UserFoundation{}
	}

	publicKeyBase58 := base58.Encode(publicKeyBytes)
	hash := sha3.Sum256(publicKeyBytes)
	addressBase58Check := base58.CheckEncode(hash[1:], hash[0])

	return &UserFoundation{
		PrivateKeyBytes:    privateKeyBytes,
		PublicKeyBytes:     publicKeyBytes,
		PublicKeyType:      keyType,
		PublicKeyBase58:    publicKeyBase58,
		AddressBase58Check: addressBase58Check,
		UserID:             "testuser",
	}
}

func UserFoundationFromPrivateKey(privateKey ed25519.PrivateKey) (*UserFoundation, error) {
	publicKey, ok := privateKey.Public().(ed25519.PublicKey)
	if !ok {
		return nil, errors.New("type requireion failed")
	}

	publicKeyBase58 := base58.Encode(publicKey)
	hash := sha3.Sum256(publicKey)
	addressBase58Check := base58.CheckEncode(hash[1:], hash[0])

	return &UserFoundation{
		PrivateKeyBytes:    privateKey,
		PublicKeyBytes:     publicKey,
		PublicKeyBase58:    publicKeyBase58,
		AddressBase58Check: addressBase58Check,
		UserID:             "testuser",
		PublicKeyType:      pbfound.KeyType_ed25519.String(),
	}, nil
}

func UserFoundationFromBase58CheckPrivateKey(base58Check string) (*UserFoundation, error) {
	decode, ver, err := base58.CheckDecode(base58Check)
	if err != nil {
		return nil, fmt.Errorf("check decode: %w", err)
	}
	privateKey := ed25519.PrivateKey(append([]byte{ver}, decode...))

	return UserFoundationFromPrivateKey(privateKey)
}

func (u *UserFoundation) Sign(args ...string) (publicKeyBase58 string, signMsg []byte, err error) {
	publicKeyBase58 = u.PublicKeyBase58
	msg := make([]string, 0, len(args)+1)
	msg = append(msg, args...)
	msg = append(msg, publicKeyBase58)

	message := []byte(strings.Join(msg, ""))

	if signMsg, err = u.sign(message); err != nil {
		return "", nil, err
	}

	return
}

func (u *UserFoundation) sign(message []byte) (signMsg []byte, err error) {
	switch u.PublicKeyType {
	case pbfound.KeyType_ed25519.String():
		hash := sha3.Sum256(message)
		signMsg = signMessageEd25519(u.PrivateKeyBytes, hash[:])
		err = verifyEd25519(u.PublicKeyBytes, hash[:], signMsg)
		if err != nil {
			return nil, err
		}

	case pbfound.KeyType_secp256k1.String():
		sha256 := sha3.Sum256(message)
		hash := eth.Hash(sha256[:])
		signMsg = signMessageSecp256k1(u.PrivateKeyBytes, hash)
		err = verifySecp256k1(u.PublicKeyBytes, hash, signMsg)
		if err != nil {
			return nil, err
		}

	default:
		return nil, errors.New("unknown key type")
	}

	return
}

func (u *UserFoundation) SetUserID(id string) {
	if len(id) != 0 {
		u.UserID = id
	}
}

// MultiSign - added multi sign
func MultiSign(args []string, users ...*UserFoundation) (publicKeysBase58 []string, signMsgs [][]byte, err error) {
	msg := make([]string, 0, len(args)+len(users))
	msg = append(msg, args...)
	for _, i := range users {
		msg = append(msg, i.PublicKeyBase58)
		publicKeysBase58 = append(publicKeysBase58, i.PublicKeyBase58)
	}

	bytesToSign := sha3.Sum256([]byte(strings.Join(msg, "")))

	for _, i := range users {
		var sMsg []byte
		if sMsg, err = i.sign(bytesToSign[:]); err != nil {
			return nil, nil, err
		}
		signMsgs = append(signMsgs, sMsg)
	}

	return
}

// signMessageEd25519 - sign arguments with private key in ed25519
func signMessageEd25519(privateKey ed25519.PrivateKey, msgToSign []byte) []byte {
	sig := ed25519.Sign(privateKey, msgToSign)
	return sig
}

// verifyEd25519 - verify publicKey in ed25519 with message and signed message
func verifyEd25519(publicKey []byte, bytesToSign []byte, sMsg []byte) error {
	if !ed25519.Verify(publicKey, bytesToSign, sMsg) {
		err := errors.New("valid signature rejected")
		return err
	}
	return nil
}

// signMessageSecp256k1 - signs a message with private key in secp256k1
func signMessageSecp256k1(privateKeyBytes []byte, msgToSign []byte) []byte {
	privateKey, err := eth.PrivateKeyFromBytes(privateKeyBytes)
	if err != nil {
		return nil
	}
	sig, err := eth.Sign(msgToSign, privateKey)
	if err != nil {
		return nil
	}
	return sig
}

// verifySecp256k1 - verify publicKey in secp256k1 with message and signed message
func verifySecp256k1(publicKeyBytes []byte, message []byte, sig []byte) error {
	if !eth.Verify(publicKeyBytes, message, sig) {
		return errors.New("secp256k1 signature rejected")
	}

	return nil
}
