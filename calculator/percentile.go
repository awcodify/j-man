package calculator

import (
	"math"
	"sort"
)

// Result TODO
type Result struct {
	ResponseTimes []float64
}

// CalculatePercentileRank TODO
func CalculatePercentileRank(percentile float64, collection []float64) float64 {
	numberOfItem := float64(len(collection))
	return (percentile / 100 * numberOfItem)
}

// Percentile TODO
func (collection *Result) Percentile(percentile float64) float64 {
	data := collection.ResponseTimes

	// We need to sort first to calculate precentile
	sort.Float64s(data)

	calculatedRank := CalculatePercentileRank(percentile, data)

	// rank should be int because it represent the position of data (index)
	// So, we need to round it into integer
	rank := math.Round(calculatedRank)

	return (data)[int(rank)-1]
}

// Average TODO:
func (collection *Result) Average() float64 {
	data := collection.ResponseTimes
	var total float64

	for _, value := range data {
		total += value
	}
	return (total / float64(len(data)))
}
