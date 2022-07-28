package keeper

import (
	"github.com/DenisCherkasskikh/insurance/x/insurance/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetNextRec set nextRec in the store
func (k Keeper) SetNextRec(ctx sdk.Context, nextRec types.NextRec) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NextRecKey))
	b := k.cdc.MustMarshal(&nextRec)
	store.Set([]byte{0}, b)
}

// GetNextRec returns nextRec
func (k Keeper) GetNextRec(ctx sdk.Context) (val types.NextRec, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NextRecKey))

	b := store.Get([]byte{0})
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveNextRec removes nextRec from the store
func (k Keeper) RemoveNextRec(ctx sdk.Context) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.NextRecKey))
	store.Delete([]byte{0})
}
