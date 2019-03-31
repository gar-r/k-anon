package partition

import (
	"gonum.org/v1/gonum/floats"
	"testing"
)

func TestFloatRange_Split(t *testing.T) {

	t.Run("split range", func(t *testing.T) {
		r := NewFloatRange(0.335, 0.775)
		r1, r2 := r.Split()
		e1 := NewFloatRange(0.335, 0.555)
		e2 := NewFloatRange(0.555, 0.775)
		if !e1.Equals(r1) {
			t.Errorf("expected %v, got %v", e1, r1)
		}
		if !e2.Equals(r2) {
			t.Errorf("expected %v, got %v", e2, r2)
		}
	})

	t.Run("split loop test", func(t *testing.T) {
		var r Range
		r = NewFloatRange(0.314159, 1.9834928)
		for r.CanSplit() {
			r, _ = r.Split()
		}
		if !floats.EqualWithinAbs(0.314159, r.Min(), delta) ||
			!floats.EqualWithinAbs(0.314159, r.Max(), delta) {
			t.Errorf("expected %v, got %v", 0.314159, r)
		}
	})

}

func TestFloat(t *testing.T) {

}
