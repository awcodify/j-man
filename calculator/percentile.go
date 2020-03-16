package calculator

import (
	"math"
	"sort"
)

// Percentile is calculate given percentile for response times
func (collection *Result) Percentile(percentile float64) float64 {
	data := collection.ResponseTimes

	// We need to sort first to calculate precentile
	sort.Float64s(data)

	calculatedRank := calculatePercentileRank(percentile, data)

	// rank should be int because it represent the position of data (index)
	// So, we need to round it into integer
	rank := math.Round(calculatedRank)

	return (data)[int(rank)-1]
}

// Average will calculate average response time
func (collection *Result) Average() float64 {
	data := collection.ResponseTimes
	var total float64

	for _, value := range data {
		total += value
	}
	return (total / float64(len(data)))
}

func calculatePercentileRank(percentile float64, collection []float64) float64 {
	numberOfItem := float64(len(collection))
	return (percentile / 100 * numberOfItem)
}
