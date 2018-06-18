package vanity

import (
	"testing"

	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/assert"
)

var Wallet0xabcd8 = walletFromPrivateKey("59daf052610c858344b708581f3020a22960b60f707fe0af3caa8905ff0f2788")
var Wallet0xABcd1 = walletFromPrivateKey("635c74f541eba7d40bebedf656cf00b2d4eeb0ead5c69aa1e81ee15824f1df5d")
var Wallet0x000001 = walletFromPrivateKey("07fd70b0276d5676688414df3ac3305580adbd28c0fa057109e9c21b774951ee")

func TestIsValidVanityWallet(t *testing.T) {
	w := Wallet0xabcd8
	assert.Equal(t, "0xabcd8", w.Address.Hex()[:7])
	assert.True(t, w.isValidVanityAddress("abcd8", true))
	assert.True(t, w.isValidVanityAddress("abcd8", false))
	assert.False(t, w.isValidVanityAddress("abCd8", true))

	w = Wallet0xABcd1
	assert.Equal(t, "0xABcd1", w.Address.Hex()[:7])
	assert.True(t, w.isValidVanityAddress("ABcd1", true))
	assert.False(t, w.isValidVanityAddress("ABCD1", true))
	assert.True(t, w.isValidVanityAddress("abcd1", false))

	w = Wallet0x000001
	assert.Equal(t, "0x000001", w.Address.Hex()[:8])
	assert.True(t, w.isValidVanityAddress("000001", true))
	assert.True(t, w.isValidVanityAddress("000001", false))
	assert.False(t, w.isValidVanityAddress("00000a", false))
}

func TestGetVanityWallet(t *testing.T) {
	ch := make(chan VanityOutput, 1)
	go GetVanityWallet("a", true, ch)
	res := <-ch
	assert.True(t, res.Success)
	assert.Equal(t, "0xa", res.Wallet.Address.Hex()[:3])

	go GetVanityWallet("1", false, ch)
	res = <-ch
	assert.True(t, res.Success)
	assert.Equal(t, "0x1", res.Wallet.Address.Hex()[:3])

	go GetVanityWallet("DeadBeef1234", true, ch)
	res = <-ch
	assert.False(t, res.Success)
}

func BenchmarkGetRandomWallet(b *testing.B) {
	var w *Wallet
	for n := 0; n < b.N; n++ {
		w = getRandomWallet()
	}
	_ = w
}

func BenchmarkGenerateKey(b *testing.B) {
	var private *ecdsa.PrivateKey
	for n := 0; n < b.N; n++ {
		private, _ = crypto.GenerateKey()
	}
	_ = private
}

func BenchmarkPrivateToPublic(b *testing.B) {
	var address common.Address
	private, _ := crypto.GenerateKey()

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		address = crypto.PubkeyToAddress(private.PublicKey)

	}
	_ = address
}

func BenchmarkToCaseSensitive(b *testing.B) {
	w := Wallet0xABcd1
	for n := 0; n < b.N; n++ {
		w.Address.Hex()
	}
}

// Generic benchmark
func benchVanity(wallet *Wallet, input string, checksum bool, b *testing.B) {
	for n := 0; n < b.N; n++ {
		wallet.isValidVanityAddress(input, checksum)
	}
}

func BenchmarkInvalidCharsSensitive(b *testing.B)   { benchVanity(Wallet0xABcd1, "DeadBeef", true, b) }
func BenchmarkInvalidCaseSensitive(b *testing.B)    { benchVanity(Wallet0xABcd1, "ABcD1", true, b) }
func BenchmarkInvalidCharsInsensitive(b *testing.B) { benchVanity(Wallet0xABcd1, "deadbeef", false, b) }
func BenchmarkValidCharsSensitive(b *testing.B)     { benchVanity(Wallet0xABcd1, "ABcd1", true, b) }
func BenchmarkValidCharsInsensitive(b *testing.B)   { benchVanity(Wallet0xABcd1, "abccd1", false, b) }

func walletFromPrivateKey(privateKey string) *Wallet {
	pr, _ := crypto.HexToECDSA(privateKey)
	return &Wallet{
		PrivateKey: pr,
		Address:    crypto.PubkeyToAddress(pr.PublicKey),
	}
}
