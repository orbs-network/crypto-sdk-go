package cryptosdk

// #include "ced25519key.h"
import "C"
import (
	"fmt"
	"unsafe"
)

type GoED25519Key struct {
	inst C.CED25519Key
}

func ED25519KeyNew() GoED25519Key {
	var ret GoED25519Key
	ret.inst = C.CED25519KeyInit()
	return ret
}
func (ed GoED25519Key) ED25519KeyFree() {
	C.CED25519KeyFree(ed.inst)
}
func (ed GoED25519Key) ED25519KeyHasPrivateKey() bool {
	return C.CED25519KeyHasPrivateKey(ed.inst) == true
}

func (ed GoED25519Key) ED25519KeyGetPublicKey() []byte {

	var len C.int
	res := unsafe.Pointer(C.CED25519KeyGetPublicKey(ed.inst, &len))
	fmt.Printf("Res: %T %v len=%d\n", res, res, len)
	return C.GoBytes(res, len)
}
