package vanity

import (
	"encoding/hex"

	"strings"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/sha3"
)

const step = 500

func GetVanityWallet(input string, checksum bool, ch chan VanityOutput) {
	res := VanityOutput{
		Attempts: 0,
	}
	for {
		res.Attempts++
		wallet := getRandomWallet()
		if wallet.isValidVanityAddress(input, checksum) {
			res.Success = true
			res.Wallet = wallet
			ch <- res
			return
		}
		if res.Attempts >= step {
			ch <- res
			res.Attempts = 0
		}
	}
}

func getRandomWallet() *Wallet {
	private, _ := crypto.GenerateKey()
	address := crypto.PubkeyToAddress(private.PublicKey)
	return &Wallet{
		PrivateKey: private,
		Address:    address,
	}
}

func (w *Wallet) isValidVanityAddress(input string, checksum bool) bool {
	hexAddress := hex.EncodeToString(w.Address[:])
	if !checksum {
		return input == hexAddress[:len(input)]
	}
	if strings.ToLower(input) != hexAddress[:len(input)] {
		return false
	}

	sha := sha3.NewKeccak256()
	sha.Write([]byte(hexAddress)) // nolint: errcheck
	hash := sha.Sum(nil)

	for i := 0; i < len(input); i++ {
		if input[i] <= '9' {
			continue
		}
		hashByte := hash[i/2]
		if i%2 == 0 {
			hashByte = hashByte >> 4
		} else {
			hashByte &= 0xf
		}
		if (input[i] <= 'F' && hashByte <= 7) || (input[i] > 'F' && hashByte > 7) {
			return false
		}
	}

	return true
}
