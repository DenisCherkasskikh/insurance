package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgAddRecord = "add_record"

var _ sdk.Msg = &MsgAddRecord{}

func NewMsgAddRecord(creator string, brand string, model string, year string, owner string, licensePlate string, vinNumber string, odometer string, side string, part string, payout string) *MsgAddRecord {
	return &MsgAddRecord{
		Creator:      creator,
		Brand:        brand,
		Model:        model,
		Year:         year,
		Owner:        owner,
		LicensePlate: licensePlate,
		VinNumber:    vinNumber,
		Odometer:     odometer,
		Side:         side,
		Part:         part,
		Payout:       payout,
	}
}

func (msg *MsgAddRecord) Route() string {
	return RouterKey
}

func (msg *MsgAddRecord) Type() string {
	return TypeMsgAddRecord
}

func (msg *MsgAddRecord) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgAddRecord) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgAddRecord) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
