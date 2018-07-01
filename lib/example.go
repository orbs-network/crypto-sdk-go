package main

// #cgo CXXFLAGS: -std=c++11
// #cgo CPPFLAGS: -I${SRCDIR}/../native/headers
// #cgo LDFLAGS:  -L${SRCDIR}/../native/mac -lcryptosdk
// #include "cbase58.h"
import "C"
import (
	"fmt"
)

func main() {
	fmt.Println("Hello world")
}
