package keeper

import (
	"github.com/DenisCherkasskikh/insurance/x/insurance/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetCrashRecord set a specific crashRecord in the store from its index
func (k Keeper) SetCrashRecord(ctx sdk.Context, crashRecord types.CrashRecord) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CrashRecordKeyPrefix))
	b := k.cdc.MustMarshal(&crashRecord)
	store.Set(types.CrashRecordKey(
		crashRecord.Index,
	), b)
}

// GetCrashRecord returns a crashRecord from its index
func (k Keeper) GetCrashRecord(
	ctx sdk.Context,
	index string,

) (val types.CrashRecord, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CrashRecordKeyPrefix))

	b := store.Get(types.CrashRecordKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveCrashRecord removes a crashRecord from the store
func (k Keeper) RemoveCrashRecord(
	ctx sdk.Context,
	index string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CrashRecordKeyPrefix))
	store.Delete(types.CrashRecordKey(
		index,
	))
}

// GetAllCrashRecord returns all crashRecord
func (k Keeper) GetAllCrashRecord(ctx sdk.Context) (list []types.CrashRecord) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CrashRecordKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.CrashRecord
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
