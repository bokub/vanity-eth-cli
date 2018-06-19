package main

import (
	"runtime"

	"strings"

	"time"

	"os"

	"github.com/bokub/vanity-eth-cli/src/terminal"
	"github.com/bokub/vanity-eth-cli/src/utils"
	"github.com/bokub/vanity-eth-cli/src/vanity"
)

const Checksum = true

func main() {
	input, checksum := getInput()

	cpus := runtime.NumCPU()
	ch := make(chan vanity.VanityOutput, cpus)

	for i := 0; i < cpus; i++ {
		go vanity.GetVanityWallet(input, checksum, ch)
	}
	start := time.Now()
	attempts := int64(0)
	for result := range ch {
		attempts += result.Attempts

		terminal.DisplaySpeed(int64(float64(attempts) / (time.Since(start).Seconds())))

		if result.Success {
			terminal.DisplayResult(&result, attempts, time.Since(start))
			break
		}
	}
}

func getInput() (string, bool) {
	input, err := terminal.ReadString()
	if err != nil {
		os.Exit(1)
	}
	checksum := Checksum && utils.HasLetters(input)
	if !checksum {
		input = strings.ToLower(input)
	}
	return input, checksum
}
