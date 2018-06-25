package cryptosdk_wrapper

import (
	"github.com/orbs-network/crypto-sdk-go/crypto-sdk/lib"
)

//type Vec lib.Std_vector_Sl_uint8_t_Sg_

func CryptoBase58Encode(ptr uintptr) string {

	return lib.Base58Encode(ptr)
}
