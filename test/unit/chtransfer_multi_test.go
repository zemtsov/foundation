package unit

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/anoideaopen/foundation/core"
	"github.com/anoideaopen/foundation/core/types/big"
	"github.com/anoideaopen/foundation/mock"
	"github.com/anoideaopen/foundation/token"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestMultiTransferByCustomer(t *testing.T) {
	ledger := mock.NewLedger(t)
	owner := ledger.NewWallet()

	ccConfig := makeBaseTokenConfig("CC Token", "CC", 8,
		owner.Address(), "", "", "", nil)
	initMsg := ledger.NewCC("cc", &token.BaseToken{}, ccConfig)
	require.Empty(t, initMsg)

	vtConfig := makeBaseTokenConfig("VT Token", "VT", 8,
		owner.Address(), "", "", "", nil)
	initMsg = ledger.NewCC("vt", &token.BaseToken{}, vtConfig)
	require.Empty(t, initMsg)

	user1 := ledger.NewWallet()
	user1.AddTokenBalance("cc", "CC_1", 1000)
	user1.AddTokenBalance("cc", "CC_2", 2000)

	id := uuid.NewString()

	items := []core.TransferItem{
		{
			Token:  "CC_1",
			Amount: new(big.Int).SetInt64(100),
		},
		{
			Token:  "CC_2",
			Amount: new(big.Int).SetInt64(200),
		},
	}

	itemsJSON, err := json.Marshal(items)
	require.NoError(t, err)

	_ = user1.SignedInvoke("cc", "channelMultiTransferByCustomer", id, "VT", string(itemsJSON))
	cct := user1.Invoke("cc", "channelMultiTransferFrom", id)

	_, _, err = user1.RawChTransferInvokeWithBatch("vt", "createCCMultiTransferTo", cct)
	require.NoError(t, err)
	ledger.WaitChMultiTransferTo("vt", id, time.Second*5)
	_ = user1.Invoke("vt", "channelMultiTransferTo", id)

	_, _, err = user1.RawChTransferInvoke("cc", "commitCCMultiTransferFrom", id)
	require.NoError(t, err)

	_, _, err = user1.RawChTransferInvoke("vt", "deleteCCMultiTransferTo", id)
	require.NoError(t, err)

	_, _, err = user1.RawChTransferInvoke("cc", "deleteCCMultiTransferFrom", id)
	require.NoError(t, err)

	err = user1.InvokeWithError("cc", "channelMultiTransferFrom", id)
	require.Error(t, err)
	err = user1.InvokeWithError("vt", "channelMultiTransferTo", id)
	require.Error(t, err)

	//user1.BalanceShouldBe("cc", 550)
	user1.AllowedBalanceShouldBe("vt", "CC_1", 100)
	user1.AllowedBalanceShouldBe("vt", "CC_2", 200)
	//user1.CheckGivenBalanceShouldBe("vt", "VT", 0)
	//user1.CheckGivenBalanceShouldBe("vt", "CC_1", 0)
	//user1.CheckGivenBalanceShouldBe("vt", "CC_2", 0)
	//user1.CheckGivenBalanceShouldBe("cc", "CC_1", 0)
	//user1.CheckGivenBalanceShouldBe("cc", "VT", 450)
}
