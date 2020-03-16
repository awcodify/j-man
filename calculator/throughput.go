package calculator

const (
	toSecond = 1000
	toMinute = toSecond * 60
)

func (r *Result) RPS() float64 {
	return r.rate() * toSecond
}

func (r *Result) RPM() float64 {
	return r.rate() * toMinute
}

func (r Result) rate() float64 {
	numberOfRequests := float64(len(r.ResponseTimes))
	return numberOfRequests / (r.Average() * numberOfRequests)
}
