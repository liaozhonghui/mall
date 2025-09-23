package test

import (
	"math"
	"testing"
)

func TestXxx(t *testing.T) {
	got := math.Abs(-1)
	if got != 1 {
		t.Errorf("Abs want 1， got %v", got)
	}
}
