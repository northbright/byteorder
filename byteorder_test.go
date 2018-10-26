package byteorder_test

import (
	"encoding/binary"
	"log"

	"github.com/northbright/byteorder"
)

func ExampleGet() {
	order := byteorder.Get()

	log.Printf("Byte Order: ")
	if order == binary.LittleEndian {
		log.Printf("Little Endian")
	} else {
		log.Printf("Big Endian")
	}

	// Output:
}
