package calculator

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRPM(t *testing.T) {
	now := time.Now()
	end := now.Add(time.Minute * time.Duration(2))
	r := Result{
		ResponseTimes: []float64{1, 1, 1, 1, 1},
		StartTime:     now,
		EndTime:       end,
	}

	actualRPM := r.RPM()
	actualRPS := r.RPS()

	assert.Equal(t, 2.50, actualRPM)
	assert.Equal(t, 0.04, actualRPS)

}
