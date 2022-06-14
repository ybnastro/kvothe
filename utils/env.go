package utils

import (
	"strconv"
	"strings"
)

// GetInt is a func to get int from a string
func GetInt(input string) int {
	i, err := strconv.Atoi(input)
	if err != nil {
		return 0
	}

	return i
}

// GetInt64 is a func to get int64 from a string
func GetInt64(input string) int64 {
	i, err := strconv.ParseInt(input, 10, 64)
	if err != nil {
		return 0
	}
	return i
}

// GetBool is a func to get bool from a string
func GetBool(input string) bool {
	if strings.EqualFold(input, "true") {
		return true
	}
	if strings.EqualFold(input, "false") {
		return false
	}
	return false
}
