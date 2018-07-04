package cryptosdk_test

import (
	"github.com/orbs-network/crypto-sdk-go/cryptosdk"
	"fmt"
	"encoding/hex"
	"testing"
)

func TestEd25519key(t *testing.T) {
	fmt.Println("===> TestEd25519key <===")
	inst := cryptosdk.ED25519KeyNew()
	defer inst.ED25519KeyFree()
	fmt.Printf("Instance: %T %v\n", inst, inst)
	fmt.Printf("Instance has priv key: %v\n", inst.ED25519KeyHasPrivateKey())
	publicKey := inst.ED25519KeyGetPublicKey()
	publicKeyStr := hex.EncodeToString(publicKey)
	fmt.Printf("Public Key: %T %v\n", publicKey, publicKey)
	fmt.Printf("Public Key str: %T %v\n", publicKeyStr, publicKeyStr)
}