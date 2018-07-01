package cryptosdk_wrapper

import (
	"testing"
	"fmt"
)

func TestBase58(t *testing.T) {

	testString := "Hello this is a nice string"

	encoded := CryptoBase58Encode([]byte(testString))
	decoded := CryptoBase58Decode(encoded)

	//byteArray = uintptr("hello")
	//base58Result := CryptoBase58Decode(string("41AM49hyv2AcCKGybYeqPprmJ"))
	fmt.Printf("Result: encoded: %v decoded: %v\n", encoded, decoded)
	//fmt.Println(base58Result.Swigcptr())
	//RegisterFailHandler(Fail)
	//RunSpecs(t, "PublicApi Service Spec Suite")
}

//var _ = Describe("Base58", func() {
//	It("should Encode Base58 correctly", func() {
//	})
//})
