package cryptosdk_wrapper

import (
	"testing"
	"fmt"
	"unsafe"
	"encoding/binary"
)

func TestBase58(t *testing.T) {

	var u uintptr = 42
	size := unsafe.Sizeof(u)
	b := make([]byte, size)
	switch size {
	case 4:
		binary.LittleEndian.PutUint32(b, uint32(u))
	case 8:
		binary.LittleEndian.PutUint64(b, uint64(u))
	default:
		panic(fmt.Sprintf("unknown uintptr size: %v", size))
	}

	//byteArray = uintptr("hello")
	base58Result := CryptoBase58Encode(u)
	fmt.Printf("Result: %v", base58Result)
	//RegisterFailHandler(Fail)
	//RunSpecs(t, "PublicApi Service Spec Suite")
}

//var _ = Describe("Base58", func() {
//	It("should Encode Base58 correctly", func() {
//	})
//})
