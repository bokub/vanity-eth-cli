package main

import (
	"runtime"

	"strings"

	"time"

	"github.com/bokub/vanity-eth-cli/src/terminal"
	"github.com/bokub/vanity-eth-cli/src/utils"
	"github.com/bokub/vanity-eth-cli/src/vanity"
)

const Input = "Ab0"
const Checksum = true

func main() {
	cpus := runtime.NumCPU()
	ch := make(chan vanity.VanityOutput, cpus)

	input, checksum := getInput()

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

// TODO get from cli args
func getInput() (string, bool) {
	input := Input
	checksum := Checksum && utils.HasLetters(Input)
	if !checksum {
		input = strings.ToLower(input)
	}
	return input, checksum
}
