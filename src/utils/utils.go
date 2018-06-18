package utils

import "time"

func HasLetters(input string) bool {
	for i := 0; i < len(input); i++ {
		if input[i] > '9' {
			return true
		}
	}
	return false
}

func RoundDuration(d, r time.Duration) time.Duration {
	if m := d % r; m+m < r {
		d = d - m
	} else {
		d = d + r - m
	}
	return d
}
