package aggregator

import (
	"time"

	"github.com/awcodify/j-man/calculator"
)

// AggregatedResult TODO
type AggregatedResult struct {
	ResponseTime responseTime
}

// responseTime is time consumed by JMeter from start to the end
type responseTime struct {
	Average float64
	P95     float64
	RPM     float64
	RPS     float64
}

// Aggregate will aggregate response times
func (c Collector) Aggregate() AggregatedResult {
	responseTimes := make([]float64, 0, len(c.Summary))
	lastIndex := len(c.Summary) - 1
	for _, line := range c.Summary {
		responseTimes = append(responseTimes, float64(line.Elapsed))
	}

	calculator := calculator.Result{
		ResponseTimes: responseTimes,
		StartTime:     time.Unix(c.Summary[0].Timestamp, 0),
		EndTime:       time.Unix(c.Summary[lastIndex].Timestamp, 0),
	}

	return AggregatedResult{
		ResponseTime: responseTime{
			Average: calculator.Average(),
			P95:     calculator.Percentile(95),
			RPM:     calculator.RPM(),
			RPS:     calculator.RPS(),
		},
	}
}
