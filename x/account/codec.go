package account

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

// RegisterCodec registers concrete types on the Amino codec
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgSetName{}, "account/SetName", nil)
	cdc.RegisterConcrete(MsgBuyName{}, "account/BuyName", nil)
}
