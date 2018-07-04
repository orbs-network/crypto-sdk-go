package cryptosdk

// #include "cbase58.h"
import "C"

func Base58Encode(plaintext []byte, len int) string {
	res := C.Base58Encode(C.CString(string(plaintext)), C.int(len))
	return C.GoString(res)
}

func Base58Decode(encoded string) []byte {
	res := C.Base58Decode(C.CString(encoded))
	return []byte(C.GoString(res))
}
