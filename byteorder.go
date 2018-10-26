package byteorder

import (
	"encoding/binary"
	"unsafe"
)

// Get detects and returns the byte order of runtime system.
// Return type: binary.ByteOrder interface.
// Possible values: binary.LittleEndian and binary.BigEndian.
// See https://godoc.org/encoding/binary#pkg-variables for more information.
func Get() binary.ByteOrder {
	const size int = int(unsafe.Sizeof(0))
	var (
		i int = 1
	)

	bytesBuf := (*[size]byte)(unsafe.Pointer(&i))
	if bytesBuf[0] == 1 {
		return binary.LittleEndian
	}
	return binary.BigEndian
}
