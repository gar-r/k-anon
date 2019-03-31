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

func TestFloatRange_Min(t *testing.T) {
	r := NewFloatRange(0.1, 0.8)
	actual := r.Min()
	expected := 0.1
	if expected != actual {
		t.Errorf("expected %v, got %v", expected, actual)
	}
}

func TestFloatRange_Max(t *testing.T) {
	r := NewFloatRange(0.1, 0.8)
	actual := r.Max()
	expected := 0.8
	if expected != actual {
		t.Errorf("expected %v, got %v", expected, actual)
	}
}

func TestFloatRange_Equals(t *testing.T) {

	r := NewFloatRange(0.1, 0.8)

	t.Run("different type", func(t *testing.T) {
		p := NewSet()
		if r.Equals(p) {
			t.Errorf("expected false")
		}
	})

	t.Run("float range equals", func(t *testing.T) {
		p := NewFloatRange(0.1, 0.8)
		if !r.Equals(p) {
			t.Errorf("expected true")
		}
	})

	t.Run("float range not equals", func(t *testing.T) {
		p := NewFloatRange(0.7, 0.7)
		if r.Equals(p) {
			t.Errorf("expected false")
		}
	})

}

func TestFloatRange_Contains(t *testing.T) {

	r := NewFloatRange(0.1, 0.8)

	t.Run("different type", func(t *testing.T) {
		item := "test"
		if r.Contains(item) {
			t.Errorf("expected false")
		}
	})

	t.Run("float type", func(t *testing.T) {
		item := 0.5
		if !r.Contains(item) {
			t.Errorf("expected true")
		}
	})

	t.Run("value = min", func(t *testing.T) {
		item := 0.1
		if !r.Contains(item) {
			t.Errorf("expected true")
		}
	})

	t.Run("value = max", func(t *testing.T) {
		item := 0.8
		if !r.Contains(item) {
			t.Errorf("expected true")
		}
	})

}

func TestFloatRange_ContainsPartition(t *testing.T) {

	r := NewFloatRange(0.1, 0.8)

	t.Run("incompatible type", func(t *testing.T) {
		p := NewItem("test")
		if r.ContainsPartition(p) {
			t.Errorf("expected false")
		}
	})

	t.Run("range type equals", func(t *testing.T) {
		p := NewFloatRange(0.1, 0.8)
		if !r.ContainsPartition(p) {
			t.Errorf("expected true")
		}
	})

	t.Run("range type contains", func(t *testing.T) {
		p := NewFloatRange(0.3, 0.6)
		if !r.ContainsPartition(p) {
			t.Errorf("expected true")
		}
	})

	t.Run("range type not contains", func(t *testing.T) {
		p := NewFloatRange(0.3, 0.9)
		if r.ContainsPartition(p) {
			t.Errorf("expected false")
		}
	})

	t.Run("set type contains", func(t *testing.T) {
		p := NewSet(0.3, 0.4, 0.5)
		if !r.ContainsPartition(p) {
			t.Errorf("expected true")
		}
	})

	t.Run("set type not contains", func(t *testing.T) {
		r := NewFloatRange(0.1, 0.8)
		p := NewSet(0.3, 0.4, 0.9)
		if r.ContainsPartition(p) {
			t.Errorf("expected false")
		}
	})

}

func TestFloatRange_String(t *testing.T) {

	t.Run("singular range", func(t *testing.T) {
		r := NewFloatRange(0.5, 0.5)
		actual := r.String()
		expected := "(0.500000)"
		if expected != actual {
			t.Errorf("expected %v, got %v", expected, actual)
		}
	})

	t.Run("normal range", func(t *testing.T) {
		r := NewFloatRange(0.1, 0.8)
		actual := r.String()
		expected := "(0.100000..0.800000)"
		if expected != actual {
			t.Errorf("expected %v, got %v", expected, actual)
		}
	})

}
