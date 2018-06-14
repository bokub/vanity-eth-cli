package vanity

import (
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/common"
)

type VanityOutput struct {
	Success  bool
	Attempts int64
	Wallet   *Wallet
}

type Wallet struct {
	PrivateKey *ecdsa.PrivateKey
	Address    common.Address
}
