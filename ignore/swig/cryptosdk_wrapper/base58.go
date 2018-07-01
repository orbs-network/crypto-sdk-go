package cryptosdk_wrapper

import (
	"github.com/orbs-network/crypto-sdk-go/crypto-sdk/lib"
)

//type Vec lib.Std_vector_Sl_uint8_t_Sg_

func CryptoBase58Encode(bytes []byte) string {

	var bv lib.ByteVector
	for i, ch := range bytes  {
		bv.Set(i, lib.SwigcptrUint8_t(ch))
	}


	return lib.Base58Encode(bv)
}

func CryptoBase58Decode(input string) []byte {

	bv := lib.Base58Decode(input)
	bytes := make([]byte, bv.Size())
	var i int
	length := int(bv.Size())
	var p uint8
	for i  = 0; i<length; i++  {
		p = uint8(bv.Get(i))
		bytes[i] = p
	}

	return bytes
}
