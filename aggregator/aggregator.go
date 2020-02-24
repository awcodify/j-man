package aggregator

import (
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
}

// Aggregate will aggregate response times
func (c Collector) Aggregate() AggregatedResult {
	responseTimes := make([]float64, 0, len(c.Summary))
	for _, line := range c.Summary {
		responseTimes = append(responseTimes, float64(line.Elapsed))
	}

	calculator := calculator.Result{ResponseTimes: responseTimes}

	return AggregatedResult{
		ResponseTime: responseTime{
			Average: calculator.Average(),
			P95:     calculator.Percentile(95),
		},
	}
}
