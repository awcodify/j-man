package calculator

import "testing"

func TestCalculatePercentileRank(t *testing.T) {
	// This test scenarios using this calculator: https://goodcalculators.com/percentile-calculator/
	result := Result{
		ResponseTimes: []float64{1, 1, 1, 1, 1, 1, 2, 2, 2, 2},
	}

	got := result.Percentile(95)

	if got != 2.0 {
		t.Errorf("Got %f, want 2", got)
	}

	result = Result{
		ResponseTimes: []float64{1, 2, 3, 4, 5, 6, 7},
	}

	got = result.Percentile(25)

	if got != 2.0 {
		t.Errorf("Got %f, want 2", got)
	}
}

func TestAverage(t *testing.T) {
	result := Result{
		ResponseTimes: []float64{2, 2, 2},
	}

	got := result.Average()

	if got != 2 {
		t.Errorf("Got %f, want 2", got)
	}

	result = Result{
		ResponseTimes: []float64{1, 2, 3},
	}

	got = result.Average()

	if got != 2 {
		t.Errorf("Got %f, want 2", got)
	}
}
