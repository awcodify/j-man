package aggregator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAggregate(t *testing.T) {
	c := Collector{
		Summary: []Result{
			Result{
				Elapsed: 10,
			},
			Result{
				Elapsed: 10,
			},
		},
	}

	actual := c.Aggregate()
	expected := AggregatedResult{
		ResponseTime: responseTime{
			P95: 10,
		},
	}

	assert.Equal(t, expected, actual)
}
