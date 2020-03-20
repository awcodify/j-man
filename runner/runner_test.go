package runner

import (
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWrapOptions(t *testing.T) {
	options := Options{
		ScriptPath:     "./scripts/example.jmx",
		ResultFilePath: "./result/test.csv",
		Users:          1,
		RampUp:         1,
		Duration:       1,
	}

	actual, _ := options.WrapOptions()
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

	// For now, since ResultFilePath is dynamic based on time, we cannot assert it
	// Need to mock time or do better approach, so we can assert like this:
	//	assert.Equal(t, actual, expected)

	assert.Equal(t, actual[2], expected[2]) //expect ScriptPath
	assert.Equal(t, actual[5], expected[5]) //expect Users
	assert.Equal(t, actual[6], expected[6]) //expect RampUp
	assert.Equal(t, actual[7], expected[7]) //expect Duration

	// Now, change the file type to not to use .jmx or .csv
	options.ResultFilePath = "./result/hello.j-man"

	expectedError := "Should use .jmx or .csv file"

	_, err := options.WrapOptions()
	assert.Equal(t, expectedError, err.Error())
}

func TestRun(t *testing.T) {
	options := Options{
		ScriptPath:     "./scripts/example.jmx",
		ResultFilePath: "./result/test.csv",
		Users:          1,
		RampUp:         1,
		Duration:       1,
	}

	t.Run("Command successfully run", func(t *testing.T) {
		resultFilePath, err := Run("echo", options)

		assert.Nil(t, err)
		assert.NotNil(t, resultFilePath)
	})

	t.Run("Failed to run command", func(t *testing.T) {
		_, err := Run("wrong", options)

		expectedError := `exec: "wrong": executable file not found in $PATH`
		if runtime.GOOS == "windows" {
			expectedError = `exec: "wrong": executable file not found in %PATH%`
		}
		assert.Equal(t, expectedError, err.Error())
	})

	t.Run("Failed to wrap options", func(t *testing.T) {
		options.ResultFilePath = "test.failed"
		_, err := Run("wrong", options)

		assert.Equal(t, "Should use .jmx or .csv file", err.Error())
	})
}
