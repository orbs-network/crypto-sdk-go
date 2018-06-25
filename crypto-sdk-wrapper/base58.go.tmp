package crypto_sdk_wrapper

import (
	"github.com/orbs-network/crypto-sdk-go/crypto-sdk/lib"
)

// #cgo CFLAGS: -I${SRCDIR}/crypto-sdk/lib
// #cgo LDFLAGS: -L${SRCDIR}/crypto-sdk/lib

// #include "base58.h"
import "C"

func Base58Encode(input *[]byte) string {
	return cryptosdk.Base58Encode(input)
}