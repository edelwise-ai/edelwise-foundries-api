package utils

import "regexp"

// This file will be used to store helper functions that will be used in multiple packages
// For example, to check email input format with regex

func CheckFormat(input string) bool {
	if !regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`).MatchString(input) {
		return false
	}
	return true
}
