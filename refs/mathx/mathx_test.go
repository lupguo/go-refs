package mathx

import (
	"testing"
)

func TestFormat(t *testing.T) {
	numbers := []float64{1.231, 1.235}
	for _, number := range numbers {
		t.Logf("%0.2f", number)
	}
}
