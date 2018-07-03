package main

// #include "cbase58.h"
// #include "ced25519key.h"
import "C"
import (
	"fmt"
	"unsafe"
)

type GoED25519Key struct {
	inst C.CED25519Key
}

func main() {

	//bindCpp()
	ed25519keyTest()
	ed25519keyTest()

}

func New() GoED25519Key {
	var ret GoED25519Key
	ret.inst = C.CED25519KeyInit()
	return ret
}
func (ed GoED25519Key) Free() {
	C.CED25519KeyFree(ed.inst)
}
func (ed GoED25519Key) HasPrivateKey() bool {
	return C.CED25519KeyHasPrivateKey(ed.inst) == true
}

func (ed GoED25519Key) GetPublicKey() []byte  {

	var len C.int
	res := unsafe.Pointer(C.CED25519KeyGetPublicKey(ed.inst, &len))
	return C.GoBytes(res, len)
}

func ed25519keyTest() {
	inst := New()
	fmt.Printf("Instance: %T %v\n", inst, inst)
	fmt.Printf("Instance has priv key: %v\n", inst.HasPrivateKey())
	publicKey := inst.GetPublicKey()
	fmt.Printf("Public Key: %T %v\n", publicKey, publicKey)
	inst.Free()
}

/*
func bindC() {
	input := "Hello world"
	fmt.Printf("Input: %v\n", input)
	cstr := C.CString(input)
	encodeRes := C.Base58Encode(cstr)
	encodedStr := C.GoString(encodeRes)
	fmt.Printf("Encoded: %v\n", encodedStr)
	decodeRes := C.Base58Decode(C.CString(encodedStr))
	decodeStr := C.GoString(decodeRes)
	fmt.Printf("Decoded: %v\n", decodeStr)
}
*/
//func bindCpp()  {
//	input := "Hello world"
//	fmt.Printf("Input: %v\n", input)
//	cstr := C.CString(input)
//	encodeRes := C._Base58Encode(cstr)
//	encodedStr := C.GoString(encodeRes)
//	fmt.Printf("Encoded: %v\n", encodedStr)
//	decodeRes := C._Base58Decode(C.CString(encodedStr))
//	decodeStr := C.GoString(decodeRes)
//	fmt.Printf("Decoded: %v\n", decodeStr)
//
//}
