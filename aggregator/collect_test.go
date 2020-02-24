package aggregator

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

// To avoid test using File, we seperated the method for parsing CSV to string
func Test_parseCSVToStrings(t *testing.T) {
	content := []byte("This is a test string")
	dir, err := ioutil.TempDir("", "example")
	if err != nil {
		log.Fatal(err)
	}

	defer os.RemoveAll(dir) // clean up

	tmpfn := filepath.Join(dir, "tmpfile")
	if err := ioutil.WriteFile(tmpfn, content, 0666); err != nil {
		log.Fatal(err)
	}

	collector := Collector{Source: tmpfn}
	expected := [][]string{[]string{"This is a test string"}}
	actual := collector.Collect()

	assert.Equal(t, expected, actual)
}

func TestToResult(t *testing.T) {
	raw := [][]string{
		[]string{"100", "100", "label", "200", "OK",
			"threadName", "json", "true", "", "99",
			"100", "GroupThreads", "AllThreads", "url",
			"14", "15", "16"},
	}
	collector := Collector{Raw: raw}

	expected := []Result{
		Result{
			Elapsed:         100.0,
			Label:           "label",
			ResponseCode:    200,
			ResponseMessage: "OK",
			ThreadName:      "threadName",
			DataType:        "json",
			Success:         true,
			FailureMessage:  "",
			Bytes:           99,
			SentBytes:       100,
			GroupThreads:    "GroupThreads",
			AllThreads:      "AllThreads",
			URL:             "url",
			Latency:         14,
			IdleTime:        15,
			Connect:         16,
		}}
	actual := collector.ToResult()

	assert.Equal(t, expected, actual.Summary)
}
