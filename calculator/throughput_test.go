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
		ResponseTimes: []float64{200, 200, 200, 200, 200},
	}

	actualRPM := r.RPM()
	actualRPS := r.RPS()

	assert.Equal(t, float64(300), actualRPM)
	assert.Equal(t, float64(5), actualRPS)
}
