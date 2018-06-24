package terminal

import (
	"bufio"
	"fmt"
	"regexp"
	"strings"
)

var re = regexp.MustCompile("^[0-9a-fA-F]*$")

func ReadString(reader *bufio.Reader) (string, error) {
	for {
		fmt.Println("Enter an address prefix of your choice: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Print("Cannot read terminal input. Exiting program.\nError: " + err.Error())
			return input, err
		}
		input = strings.TrimSpace(input)
		if validatePrefix(input) {
			return input, nil
		}
	}
}

func ReadBool(reader *bufio.Reader) (bool, error) {
	fmt.Println("Is your prefix case-sensitive? [Y/n]")
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Print("Cannot read terminal input. Exiting program.\nError: " + err.Error())
		return false, err
	}
	input = strings.ToLower(strings.TrimSpace(input))
	return input != "n", nil
}

func validatePrefix(input string) bool {
	if !re.MatchString(input) {
		fmt.Println("Your prefix can only contains numers and letters from A to F")
		return false
	}
	if len(input) > 40 {
		fmt.Println("An ETH address is only 40 characters long")
		return false
	}
	return true
}
