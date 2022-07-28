package keeper

import (
	"context"
	"strconv"

	"github.com/DenisCherkasskikh/insurance/x/insurance/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) AddRecord(goCtx context.Context, msg *types.MsgAddRecord) (*types.MsgAddRecordResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	nextRec, found := k.Keeper.GetNextRec(ctx)
	if !found {
		panic("Next record not found")
	}
	newNum := strconv.FormatUint(nextRec.RecNum, 10)
	crashRecord := types.CrashRec{
		Index:        newNum,
		Brand:        msg.Brand,
		Model:        msg.Model,
		Year:         msg.Year,
		Owner:        msg.Owner,
		LicensePlate: msg.LicensePlate,
		VinNumber:    msg.VinNumber,
		Odometer:     msg.Odometer,
		Side:         msg.Side,
		Part:         msg.Part,
		Payout:       msg.Payout,
	}
	k.Keeper.SetCrashRec(ctx, crashRecord)
	nextRec.RecNum++
	k.Keeper.SetNextRec(ctx, nextRec)

	return &types.MsgAddRecordResponse{
		RecNum: newNum,
	}, nil
}
