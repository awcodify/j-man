package aggregator

import (
	"encoding/csv"
	"os"

	"github.com/awcodify/j-man/utils"
)

// Collector will collect data from CSV result
type Collector struct {
	Source  string
	Raw     [][]string
	Summary []Result
}

// Result from JMeter csv file
type Result struct {
	ID              int64
	RoundID         int64
	Elapsed         int64
	Label           string
	ResponseCode    int64
	ResponseMessage string
	ThreadName      string
	DataType        string
	Success         bool
	FailureMessage  string
	Bytes           int64
	SentBytes       int64
	GroupThreads    string
	AllThreads      string
	URL             string
	Latency         int64
	IdleTime        int64
	Connect         int64
}

// Collect is for collecting the CSV result from JMeter
func Collect(resultFilePath string) (collector Collector) {
	file, err := os.Open(resultFilePath)
	utils.DieIf(err)
	defer file.Close()

	rows, err := csv.NewReader(file).ReadAll()
	utils.DieIf(err)

	collector.Raw = rows

	return collector
}

// ToResult will convert csv to Result
func (c Collector) ToResult() Collector {
	// we will skip first row because first row is a Header row
	coll := c.Raw[1:]
	result := make([]Result, 0, len(coll))

	for _, line := range coll {

		data := Result{
			Elapsed:         utils.ParseInt(line[1]),
			Label:           line[2],
			ResponseCode:    utils.ParseInt(line[3]),
			ResponseMessage: line[4],
			ThreadName:      line[5],
			DataType:        line[6],
			Success:         utils.ParseBool(line[7]),
			FailureMessage:  line[8],
			Bytes:           utils.ParseInt(line[9]),
			SentBytes:       utils.ParseInt(line[10]),
			GroupThreads:    line[11],
			AllThreads:      line[12],
			URL:             line[13],
			Latency:         utils.ParseInt(line[14]),
			IdleTime:        utils.ParseInt(line[15]),
			Connect:         utils.ParseInt(line[16]),
		}

		result = append(result, data)
	}
	c.Summary = result

	return c
}
