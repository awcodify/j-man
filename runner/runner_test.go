package runner

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWrapOptions(t *testing.T) {
	o := Options{
		ScriptPath:     "./scripts/example.jmx",
		ResultFileName: "./result/test.csv",
		Users:          1,
		RampUp:         1,
		Duration:       1,
	}

	actual := o.WrapOptions()
	expected := []string{
		"-nongui",
		"--testfile",
		"./scripts/example.jmx",
		"--logfile",
		"results/log.csv",
		"-Jusers=1",
		"-JrampUp=1",
		"-Jduration=1",
	}

	// For now, since ResultFileName is dynamic based on time, we cannot assert it
	// Need to mock time or do better approach, so we can assert like this:
	//	assert.Equal(t, actual, expected)

	assert.Equal(t, actual[2], expected[2]) //expect ScriptPath
	assert.Equal(t, actual[5], expected[5]) //expect Users
	assert.Equal(t, actual[6], expected[6]) //expect RampUp
	assert.Equal(t, actual[7], expected[7]) //expect Duration

	// Now, change the file type to not to use .jmx or .csv
	o.ResultFileName = "./result/hello.j-man"

	assert.PanicsWithValue(t, "Should use .jmx or .csv file", func() { o.WrapOptions() })
}

func TestRun(t *testing.T) {
	err := Run("echo", "hello", "world")

	assert.Nil(t, err)

	err = Run("thisis", "should", "be", "error")
	expectedError := `exec: "thisis": executable file not found in $PATH`

	if assert.NotNil(t, err) {
		assert.Equal(t, expectedError, err.Error())
	}
}
