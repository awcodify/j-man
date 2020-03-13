package calculator

import (
	"math"
	"time"
)

func (r *Result) RPS() float64 {
	return r.rate(func(d time.Duration) float64 {
		return d.Seconds()
	})
}

func (r *Result) RPM() float64 {
	return r.rate(func(d time.Duration) float64 {
		return d.Minutes()
	})
}

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}

func (r Result) rate(fn func(time.Duration) float64) float64 {
	length := len(r.ResponseTimes)
	diff := r.EndTime.Sub(r.StartTime)
	return toFixed(float64(length)/fn(diff), 2)
}
