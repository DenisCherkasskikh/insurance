package keeper

import (
	"github.com/DenisCherkasskikh/insurance/x/insurance/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetCrashRec set a specific crashRec in the store from its index
func (k Keeper) SetCrashRec(ctx sdk.Context, crashRec types.CrashRec) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CrashRecKeyPrefix))
	b := k.cdc.MustMarshal(&crashRec)
	store.Set(types.CrashRecKey(
		crashRec.Index,
	), b)
}

// GetCrashRec returns a crashRec from its index
func (k Keeper) GetCrashRec(
	ctx sdk.Context,
	index string,

) (val types.CrashRec, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CrashRecKeyPrefix))

	b := store.Get(types.CrashRecKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveCrashRec removes a crashRec from the store
func (k Keeper) RemoveCrashRec(
	ctx sdk.Context,
	index string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CrashRecKeyPrefix))
	store.Delete(types.CrashRecKey(
		index,
	))
}

// GetAllCrashRec returns all crashRec
func (k Keeper) GetAllCrashRec(ctx sdk.Context) (list []types.CrashRec) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CrashRecKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.CrashRec
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
