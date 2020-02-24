package utils

import (
	"strconv"
	"time"
)

// ParseTime to parse string into time
func ParseTime(text string) time.Time {
	t := ParseInt(text)
	return time.Unix(t, 0)
}

// ParseInt to parse string into int
func ParseInt(text string) int64 {
	i, err := strconv.ParseInt(text, 10, 64)
	if err != nil {
		panic(err)
	}
	return i
}

// ParseBool to parse string into boolean
func ParseBool(text string) bool {
	b, err := strconv.ParseBool(text)
	if err != nil {
		panic(err)
	}
	return b
}
