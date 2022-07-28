package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/DenisCherkasskikh/insurance/testutil/keeper"
	"github.com/DenisCherkasskikh/insurance/x/insurance"
	"github.com/DenisCherkasskikh/insurance/x/insurance/keeper"
	"github.com/DenisCherkasskikh/insurance/x/insurance/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func setupMsgServerAddRecord(t testing.TB) (types.MsgServer, keeper.Keeper, context.Context) {
	k, ctx := keepertest.InsuranceKeeper(t)
	insurance.InitGenesis(ctx, *k, *types.DefaultGenesis())
	return keeper.NewMsgServerImpl(*k), *k, sdk.WrapSDKContext(ctx)
}

func TestAddRecord(t *testing.T) {
	msgServer, _, context := setupMsgServerAddRecord(t)
	createResponse, err := msgServer.AddRecord(context, &types.MsgAddRecord{
		Brand:        "Porsche",
		Model:        "911",
		Owner:        "Jordan",
		LicensePlate: "ABC1234",
		VinNumber:    "WPO12345",
		Odometer:     "20000",
		Side:         "front-left",
		Part:         "headlight",
	})
	require.Nil(t, err)
	require.EqualValues(t, types.MsgAddRecordResponse{
		RecNum: "1",
	}, *createResponse)
}

func TestAdd1RecordGetAll(t *testing.T) {
	msgServer, keeper, context := setupMsgServerAddRecord(t)
	msgServer.AddRecord(context, &types.MsgAddRecord{
		Brand:        "Porsche",
		Model:        "911",
		Year:         "1989",
		Owner:        "Jordan",
		LicensePlate: "ABC1234",
		VinNumber:    "WPO12345",
		Odometer:     "20000",
		Side:         "left",
		Part:         "headlight",
	})
	records := keeper.GetAllCrashRec(sdk.UnwrapSDKContext(context))
	require.Len(t, records, 1)
	require.EqualValues(t, types.CrashRec{
		Index:        "1",
		Brand:        "Porsche",
		Model:        "911",
		Year:         "1989",
		Owner:        "Jordan",
		LicensePlate: "ABC1234",
		VinNumber:    "WPO12345",
		Odometer:     "20000",
		Side:         "left",
		Part:         "headlight",
	}, records[0])
}

func TestAdd3Records(t *testing.T) {
	msgServer, _, context := setupMsgServerAddRecord(t)
	msgServer.AddRecord(context, &types.MsgAddRecord{
		Brand:        "Porsche",
		Model:        "911",
		Year:         "1989",
		Owner:        "Jordan",
		LicensePlate: "ABC1234",
		VinNumber:    "WPO12345",
		Odometer:     "20000",
		Side:         "left",
		Part:         "headlight",
	})
	createResponse2, err2 := msgServer.AddRecord(context, &types.MsgAddRecord{
		Brand:        "Audi",
		Model:        "A8",
		Year:         "2020",
		Owner:        "Angela",
		LicensePlate: "DEF5678",
		VinNumber:    "WAUZ12345",
		Odometer:     "2000",
		Side:         "front-left",
		Part:         "door",
	})
	require.Nil(t, err2)
	require.EqualValues(t, types.MsgAddRecordResponse{
		RecNum: "2",
	}, *createResponse2)
	createResponse3, err3 := msgServer.AddRecord(context, &types.MsgAddRecord{
		Brand:        "Jaguar",
		Model:        "XF",
		Year:         "2017",
		Owner:        "Michael",
		LicensePlate: "GHI2345",
		VinNumber:    "SAJ3456",
		Odometer:     "20000",
		Side:         "right",
		Part:         "mirror",
	})
	require.Nil(t, err3)
	require.EqualValues(t, types.MsgAddRecordResponse{
		RecNum: "3",
	}, *createResponse3)
}

func TestAdd3RecordsGetAll(t *testing.T) {
	msgServer, keeper, context := setupMsgServerAddRecord(t)
	msgServer.AddRecord(context, &types.MsgAddRecord{
		Brand:        "Porsche",
		Model:        "911",
		Year:         "1989",
		Owner:        "Jordan",
		LicensePlate: "ABC1234",
		VinNumber:    "WPO12345",
		Odometer:     "20000",
		Side:         "left",
		Part:         "headlight",
	})
	msgServer.AddRecord(context, &types.MsgAddRecord{
		Brand:        "Audi",
		Model:        "A8",
		Year:         "2020",
		Owner:        "Angela",
		LicensePlate: "DEF5678",
		VinNumber:    "WAUZ12345",
		Odometer:     "2000",
		Side:         "front-left",
		Part:         "door",
	})
	msgServer.AddRecord(context, &types.MsgAddRecord{
		Brand:        "Jaguar",
		Model:        "XF",
		Year:         "2017",
		Owner:        "Michael",
		LicensePlate: "GHI2345",
		VinNumber:    "SAJ3456",
		Odometer:     "20000",
		Side:         "right",
		Part:         "mirror",
	})
	records := keeper.GetAllCrashRec(sdk.UnwrapSDKContext(context))
	require.Len(t, records, 3)
	require.EqualValues(t, types.CrashRec{
		Index:        "1",
		Brand:        "Porsche",
		Model:        "911",
		Year:         "1989",
		Owner:        "Jordan",
		LicensePlate: "ABC1234",
		VinNumber:    "WPO12345",
		Odometer:     "20000",
		Side:         "left",
		Part:         "headlight",
	}, records[0])
	require.EqualValues(t, types.CrashRec{
		Index:        "2",
		Brand:        "Audi",
		Model:        "A8",
		Year:         "2020",
		Owner:        "Angela",
		LicensePlate: "DEF5678",
		VinNumber:    "WAUZ12345",
		Odometer:     "2000",
		Side:         "front-left",
		Part:         "door",
	}, records[1])
	require.EqualValues(t, types.CrashRec{
		Index:        "3",
		Brand:        "Jaguar",
		Model:        "XF",
		Year:         "2017",
		Owner:        "Michael",
		LicensePlate: "GHI2345",
		VinNumber:    "SAJ3456",
		Odometer:     "20000",
		Side:         "right",
		Part:         "mirror",
	}, records[2])
}

func TestAdd1RecordHasSaved(t *testing.T) {
	msgServer, keeper, context := setupMsgServerAddRecord(t)
	msgServer.AddRecord(context, &types.MsgAddRecord{
		Brand:        "Porsche",
		Model:        "911",
		Year:         "1989",
		Owner:        "Jordan",
		LicensePlate: "ABC1234",
		VinNumber:    "WPO12345",
		Odometer:     "20000",
		Side:         "left",
		Part:         "headlight",
	})
	nextRec, found := keeper.GetNextRec(sdk.UnwrapSDKContext(context))
	require.True(t, found)
	require.EqualValues(t, types.NextRec{
		RecNum: 2,
	}, nextRec)
	record1, found1 := keeper.GetCrashRec(sdk.UnwrapSDKContext(context), "1")
	require.True(t, found1)
	require.EqualValues(t, types.CrashRec{
		Index:        "1",
		Brand:        "Porsche",
		Model:        "911",
		Year:         "1989",
		Owner:        "Jordan",
		LicensePlate: "ABC1234",
		VinNumber:    "WPO12345",
		Odometer:     "20000",
		Side:         "left",
		Part:         "headlight",
	}, record1)
}
