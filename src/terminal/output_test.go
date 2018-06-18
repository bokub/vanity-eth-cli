package terminal

import (
	"bytes"
	"io"
	"os"
	"testing"

	"time"

	"github.com/bokub/vanity-eth-cli/src/vanity"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/assert"
)

func TestDisplayResult(t *testing.T) {
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	pr, _ := crypto.HexToECDSA("07fd70b0276d5676688414df3ac3305580adbd28c0fa057109e9c21b774951ee")
	DisplayResult(&vanity.VanityOutput{Wallet: &vanity.Wallet{PrivateKey: pr, Address: crypto.PubkeyToAddress(pr.PublicKey)}},
		999, time.Second)

	outC := make(chan string)
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		outC <- buf.String()
	}()

	w.Close()
	os.Stdout = oldStdout
	out := <-outC

	assert.Contains(t, out, "Address found in 1s after 999 attempts")
	assert.Contains(t, out, "Address:     0x0000017CE99c34DA5208e04B17782226b4a655D2")
	assert.Contains(t, out, "Private key: 0x07fd70b0276d5676688414df3ac3305580adbd28c0fa057109e9c21b774951ee")
}

func TestDisplaySpeed(t *testing.T) {
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	DisplaySpeed(7777)
	outC := make(chan string)
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		outC <- buf.String()
	}()

	w.Close()
	os.Stdout = oldStdout
	out := <-outC

	assert.Contains(t, out, "Speed: 7777 hashes / second")
}
