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

	expected := [][]string{[]string{"This is a test string"}}
	actual := Collect(tmpfn)

	assert.Equal(t, expected, actual)
}
