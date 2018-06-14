package terminal

import (
	"fmt"

	"time"

	"encoding/hex"

	"github.com/bokub/vanity-eth-cli/src/vanity"
	"github.com/ethereum/go-ethereum/crypto"
)

const ResetLine = "\r\033[K"

func DisplayResult(result *vanity.VanityOutput, attempts int64, elapsed time.Duration) {
	fmt.Printf("\nAddress found in %s after %d attempts, \n", elapsed.Round(time.Second), attempts)
	fmt.Printf("Address:     %s\n", result.Wallet.Address.Hex())
	fmt.Printf("Private key: 0x%s\n", hex.EncodeToString(crypto.FromECDSA(result.Wallet.PrivateKey)))
}

func DisplaySpeed(speed int64) {
	fmt.Printf("%sSpeed: %d hashes / second", ResetLine, speed)
}
