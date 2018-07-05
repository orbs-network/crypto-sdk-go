package cryptosdk

// #include "ced25519key.h"
import "C"
import (
	"fmt"
	"unsafe"
)

type ED25519Key struct {
	inst C.CED25519Key
}

func ED25519KeyNew() ED25519Key {
	var ret ED25519Key
	ret.inst = C.CED25519KeyInit()
	return ret
}
func (ed ED25519Key) ED25519KeyFree() {
	C.CED25519KeyFree(ed.inst)
}
func (ed ED25519Key) HasPrivateKey() bool {
	return C.CED25519KeyHasPrivateKey(ed.inst) == true
}

func (ed ED25519Key) GetPublicKey() []byte {

	//var len C.int
	var len C.int
	C.CED25519KeyGetPublicKey(ed.inst, nil, &len)
	buf := make([]byte, int(len))
	bufaddr := (*C.char)(unsafe.Pointer(&buf[0]))
	fmt.Printf("Buf: len=%d, bufaddr=%p %p ed.inst=%v\n", len, buf, &buf[0], ed.inst)
	C.CED25519KeyGetPublicKey(ed.inst, bufaddr, &len)
	fmt.Printf("Res: %T %v len=%d\n", buf, buf, int(len))
	return buf
}
