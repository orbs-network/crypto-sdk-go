package cryptosdk_test

import (
	"github.com/orbs-network/crypto-sdk-go/cryptosdk"
	"fmt"
	"encoding/hex"
	"testing"
)

// This is just a POC, it still doesn't actually test anything
func TestEd25519key(t *testing.T) {
	fmt.Println("===> TestEd25519key <===")

	ed25519Key := cryptosdk.ED25519KeyNew()
	defer ed25519Key.ED25519KeyFree()
	publicKey := hex.EncodeToString(ed25519Key.GetPublicKey())
	hasPrivateKey := ed25519Key.HasPrivateKey()

	fmt.Printf("Public Key: %v\n", publicKey)
	fmt.Printf("Instance has priv key: %v\n", hasPrivateKey)
}
