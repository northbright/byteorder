# byteorder

[![Build Status](https://travis-ci.org/northbright/byteorder.svg?branch=master)](https://travis-ci.org/northbright/byteorder)
[![GoDoc](https://godoc.org/github.com/northbright/byteorder?status.svg)](https://godoc.org/github.com/northbright/byteorder)

byteorder is a [Golang](https://golang.org) package which detects byte order of runtime system.

## Example
```
package main

import (
        "encoding/binary"
        "log"

        "github.com/northbright/byteorder"
)

func main() {
        order := byteorder.Get()

        log.Printf("Byte Order: ")
        if order == binary.LittleEndian {
                log.Printf("Little Endian")
        } else {
                log.Printf("Big Endian")
        }
}

```

## Documentation
* [API References](https://godoc.org/github.com/northbright/byteorder)

## License
* [MIT License](./LICENSE)
