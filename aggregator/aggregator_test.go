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
			Average: 10,
			P95:     10,
			RPM:     6000,
			RPS:     100,
		},
	}

	assert.Equal(t, expected, actual)
}
