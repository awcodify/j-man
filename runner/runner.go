package runner

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

type extension struct {
	Name  string
	Valid bool
}

var allowedExtensions = map[string]bool{".jmx": true, ".csv": true}

// Runner is jmeter manager
type Runner struct {
	Jmeter  string
	Options []string
}

// Options need to wrap the jmeter options
type Options struct {
	ScriptPath     string
	ResultFilePath string
	Users          int64
	RampUp         int64
	Duration       int64
}

// WrapOptions is to wrap default JMeter command options
func (o *Options) WrapOptions() ([]string, error) {
	users := fmt.Sprintf("-Jusers=%v", o.Users)
	rampUp := fmt.Sprintf("-JrampUp=%v", o.RampUp)
	duration := fmt.Sprintf("-Jduration=%v", o.Duration)
	fileName, err := resultFileName(o.ResultFilePath)
	if err != nil {
		return []string{}, err
	}

	return []string{
		"-nongui",
		"--testfile",
		o.ScriptPath,
		"--logfile",
		fileName,
		users,
		rampUp,
		duration,
	}, nil
}

// Run will execute the jmeter
func Run(command string, o Options) (string, error) {
	options, err := o.WrapOptions()
	if err != nil {
		return "", err
	}
	cmd := exec.Command(command, options...)

	var stdBuffer bytes.Buffer
	mw := io.MultiWriter(os.Stdout, &stdBuffer)

	cmd.Stdout = mw
	cmd.Stderr = mw

	// Execute the command
	return options[4], cmd.Run()
}

func resultFileName(filename string) (string, error) {
	extension := filepath.Ext(filename)
	if !allowedExtensions[extension] {
		return "", fmt.Errorf("Should use .jmx or .csv file")
	}
	name := filename[0 : len(filename)-len(extension)]
	return fmt.Sprintf("%s_%s.csv", name, time.Now().Format("20060102150405")), nil
}
