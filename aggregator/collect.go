package aggregator

import (
	"encoding/csv"
	"os"

	"github.com/awcodify/j-man/utils"
)

// Collect is for collecting the CSV result from JMeter
func Collect(filePath string) [][]string {
	file, err := os.Open(filePath)
	utils.DieIf(err)
	defer file.Close()

	rows, err := csv.NewReader(file).ReadAll()
	utils.DieIf(err)

	return rows
}
