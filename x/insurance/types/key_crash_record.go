package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// CrashRecordKeyPrefix is the prefix to retrieve all CrashRecord
	CrashRecordKeyPrefix = "CrashRecord/value/"
)

// CrashRecordKey returns the store key to retrieve a CrashRecord from the index fields
func CrashRecordKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
