package types

import (
	"fmt"
)

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		NextRec: &NextRec{
			RecNum: uint64(DefaultIndex),
		},
		CrashRecList: []CrashRec{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in crashRec
	crashRecIndexMap := make(map[string]struct{})

	for _, elem := range gs.CrashRecList {
		index := string(CrashRecKey(elem.Index))
		if _, ok := crashRecIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for crashRec")
		}
		crashRecIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
