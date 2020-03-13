package aggregator

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestAggregate(t *testing.T) {
	now := time.Now()
	start := now.Unix()
	end := now.Add(time.Minute * time.Duration(2)).Unix()

	c := Collector{
		Summary: []Result{
			Result{
				Timestamp: start,
				Elapsed:   10,
			},
			Result{
				Timestamp: end,
				Elapsed:   10,
			},
		},
	}

	actual := c.Aggregate()
	expected := AggregatedResult{
		ResponseTime: responseTime{
			Average: 10,
			P95:     10,
			RPM:     1,
			RPS:     0.02,
		},
	}

	assert.Equal(t, expected, actual)
}
