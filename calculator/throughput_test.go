package calculator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRPM(t *testing.T) {
	r := Result{
		ResponseTimes: []float64{200, 200, 200, 200, 200},
	}

	actualRPM := r.RPM()
	actualRPS := r.RPS()

	assert.Equal(t, float64(300), actualRPM)
	assert.Equal(t, float64(5), actualRPS)
}
