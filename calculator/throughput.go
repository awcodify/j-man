package calculator

import (
	"math"
)

const (
	toSecond  = 1000
	toMinutes = toSecond * 60
)

func (r *Result) RPS() float64 {
	return r.rate(func(milisecond float64) float64 {
		return milisecond * toSecond
	})
}

func (r *Result) RPM() float64 {
	return r.rate(func(milisecond float64) float64 {
		return milisecond * toMinutes
	})
}

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func (r Result) rate(fn func(float64) float64) float64 {
	numberOfRequests := float64(len(r.ResponseTimes))
	return fn(numberOfRequests / (r.Average() * numberOfRequests))
}
