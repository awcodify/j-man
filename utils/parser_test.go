package utils

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestParseTime(t *testing.T) {
	actual := ParseTime("11")
	assert.Equal(t, time.Unix(11, 0), actual)
}

func TestParseInt(t *testing.T) {
	actual := ParseInt("10")
	assert.Equal(t, int64(10), actual)
}

func TestParseBool(t *testing.T) {
	actual := ParseBool("true")
	assert.Equal(t, true, actual)

	actual = ParseBool("false")
	assert.Equal(t, false, actual)
}
