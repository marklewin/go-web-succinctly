package math

import "testing"

func TestAverage(t *testing.T) {
	var v float64
	v = Average([]float64{1, 2, 3, 4, 5})
	if v != 3.0 {
		t.Error("Expected 3.0, got ", v)
	}
}
