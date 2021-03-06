package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ixofoundation/ixo-cosmos/x/ixo/sovrin"
)

func NewCreateBondMsg(bondDoc BondDoc, bondDid sovrin.SovrinDid) CreateBondMsg {
	return CreateBondMsg{
		SignBytes: "",
		TxHash:    "",
		SenderDid: "",
		BondDid:   bondDid.Did,
		PubKey:    bondDid.VerifyKey,
		Data:      bondDoc,
	}
}

func NewUpdateBondStatusMsg(txHash string, senderDid string, updateBondStatusDoc UpdateBondStatusDoc, projectDid sovrin.SovrinDid) UpdateBondStatusMsg {
	return UpdateBondStatusMsg{
		SignBytes: "",
		TxHash:    txHash,
		SenderDid: senderDid,
		BondDid:   projectDid.Did,
		Data:      updateBondStatusDoc,
	}
}

func CheckNotEmpty(value string, name string) (valid bool, err sdk.Error) {
	if len(value) == 0 {
		return false, sdk.ErrUnknownRequest(name + " is empty.")
	} else {
		return true, nil
	}
}
