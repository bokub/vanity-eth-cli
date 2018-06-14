package utils

func HasLetters(input string) bool {
	for i := 0; i < len(input); i++ {
		if input[i] > '9' {
			return true
		}
	}
	return false
}
