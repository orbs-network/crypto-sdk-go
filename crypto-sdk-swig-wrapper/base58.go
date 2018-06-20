package crypto_sdk_swig_wrapper

import (
	"github.com/orbs-network/crypto-sdk-go/crypto-sdk/lib"
)

func Base58Encode(input *[]byte) string {
	return cryptosdk.Base58Encode(nil)
}
