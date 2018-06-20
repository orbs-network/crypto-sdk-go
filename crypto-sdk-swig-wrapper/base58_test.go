package crypto_sdk_swig_wrapper_test

import (
	"fmt"
	. "github.com/onsi/ginkgo"
	"github.com/orbs-network/crypto-sdk-go/crypto-sdk-swig-wrapper"
)

var _ = Describe("Base58", func() {
	It("should Encode Base58 correctly", func() {
		var byteArray []byte
		byteArray = []byte("hello")
		base58Result := crypto_sdk_swig_wrapper.Base58Encode(&byteArray)
		fmt.Println(base58Result)
	})
})
