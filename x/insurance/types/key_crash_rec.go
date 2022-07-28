package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// CrashRecKeyPrefix is the prefix to retrieve all CrashRec
	CrashRecKeyPrefix = "CrashRec/value/"
)

// CrashRecKey returns the store key to retrieve a CrashRec from the index fields
func CrashRecKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
