package calculator

import (
	"time"
)

// Result is parsed from JMeter csv result
type Result struct {
	ResponseTimes []float64
	StartTime     time.Time
	EndTime       time.Time
}
