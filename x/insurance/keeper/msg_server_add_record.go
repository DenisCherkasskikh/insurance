package keeper

import (
	"context"

	"github.com/DenisCherkasskikh/insurance/x/insurance/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) AddRecord(goCtx context.Context, msg *types.MsgAddRecord) (*types.MsgAddRecordResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	storedRecord := types.CrashRecord{
		Index:        msg.VinNumber,
		Brand:        msg.Brand,
		Model:        msg.Model,
		Owner:        msg.Owner,
		LicencePlate: msg.LicencePlate,
		VinNumber:    msg.VinNumber,
		Mileage:      msg.Mileage,
		Side:         msg.Side,
		Part:         msg.Part,
		Payout:       msg.Payout,
	}

	k.Keeper.SetCrashRecord(ctx, storedRecord)

	number := msg.VinNumber
	return &types.MsgAddRecordResponse{
		Vin: number,
	}, nil
}
