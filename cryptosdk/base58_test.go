package cryptosdk_test

import (
	"github.com/orbs-network/crypto-sdk-go/cryptosdk"
	"fmt"
	"testing"
)

func TestBase58(t *testing.T) {
	fmt.Println("===> TestBase58 <===")
	plaintext := "Hello I am a string!"
	encoded := cryptosdk.Base58Encode([]byte(plaintext), len(plaintext))
	fmt.Printf("Encoded: %T %v\n", encoded, encoded)
	decoded := cryptosdk.Base58Decode(encoded)
	decodedStr := string(decoded)
	fmt.Printf("Decoded: %T %v\n", decodedStr, decodedStr)

	if plaintext != decodedStr {
		t.Errorf("Plaintext and decoded strings are different: Plaintext=%v Decoded=%v", plaintext, decodedStr)
	}

}
