package partition

import (
	"math"
	"testing"

	"git.okki.hu/garric/k-anon/testutil"
)

func TestNewIntRange(t *testing.T) {

	t.Run("normal int range", func(t *testing.T) {
		r := NewIntRange(5, 10)
		testutil.AssertEquals(5, r.min, t)
		testutil.AssertEquals(10, r.max, t)
	})

	t.Run("single int range", func(t *testing.T) {
		r := NewIntRange(5, 5)
		testutil.AssertEquals(5, r.min, t)
		testutil.AssertEquals(5, r.max, t)
	})

	t.Run("inverted int range", func(t *testing.T) {
		r := NewIntRange(5, 0)
		testutil.AssertEquals(5, r.min, t)
		testutil.AssertEquals(5, r.max, t)
	})

}

func TestIntRange_Contains(t *testing.T) {

	r := NewIntRange(5, 10)

	t.Run("item less than min", func(t *testing.T) {
		if r.Contains(4) {
			t.Errorf("%v should not contain %v", r, 4)
		}
	})

	t.Run("item equals min", func(t *testing.T) {
		if !r.Contains(5) {
			t.Errorf("%v should contain %v", r, 5)
		}
	})

	t.Run("item between bounds", func(t *testing.T) {
		if !r.Contains(6) {
			t.Errorf("%v should contain %v", r, 6)
		}
	})

	t.Run("item equals max", func(t *testing.T) {
		if !r.Contains(10) {
			t.Errorf("%v should contain %v", r, 10)
		}
	})

	t.Run("item greater than max", func(t *testing.T) {
		if r.Contains(11) {
			t.Errorf("%v should not contain %v", r, 11)
		}
	})

	t.Run("non integer input", func(t *testing.T) {
		if r.Contains("test") {
			t.Errorf("%v should not contain %v", r, "test")
		}
	})

}

func TestIntRange_ContainsPartition(t *testing.T) {

	r := NewIntRange(5, 10)

	t.Run("subset", func(t *testing.T) {
		r2 := NewIntRange(7, 8)
		if !r.ContainsPartition(r2) {
			t.Errorf("%v should contain %v", r, r2)
		}
	})

	t.Run("non-subset", func(t *testing.T) {
		r2 := NewIntRange(8, 12)
		if r.ContainsPartition(r2) {
			t.Errorf("%v should not contain %v", r, r2)
		}
	})

	t.Run("min subset", func(t *testing.T) {
		r2 := NewIntRange(5, 8)
		if !r.ContainsPartition(r2) {
			t.Errorf("%v should contain %v", r, r2)
		}
	})

	t.Run("max subset", func(t *testing.T) {
		r2 := NewIntRange(8, 10)
		if !r.ContainsPartition(r2) {
			t.Errorf("%v should not contain %v", r, r2)
		}
	})

	t.Run("item set contained", func(t *testing.T) {
		itemSet := NewSet(6, 7, 8)
		if !r.ContainsPartition(itemSet) {
			t.Errorf("%v should contain %v", r, itemSet)
		}
	})

	t.Run("item set not contained", func(t *testing.T) {
		itemSet := NewSet(6, 7, 8, 11)
		if r.ContainsPartition(itemSet) {
			t.Errorf("%v should not contain %v", r, itemSet)
		}
	})

	t.Run("invalid type", func(t *testing.T) {
		p := NewItem("test")
		if r.ContainsPartition(p) {
			t.Errorf("%v should not contain %v", r, p)
		}
	})

}

func TestIntRange_Equals(t *testing.T) {

	r := NewIntRange(5, 10)

	t.Run("non int range input", func(t *testing.T) {
		other := NewSet()
		if r.Equals(other) {
			t.Errorf("%v should not equal %v", r, other)
		}
	})

	t.Run("non equal int range", func(t *testing.T) {
		other := NewIntRange(3, 6)
		if r.Equals(other) {
			t.Errorf("%v should not equal %v", r, other)
		}
	})

	t.Run("equal int range", func(t *testing.T) {
		other := NewIntRange(5, 10)
		if !r.Equals(other) {
			t.Errorf("%v should be equal to %v", r, other)
		}
	})

	t.Run("nil input", func(t *testing.T) {
		if r.Equals(nil) {
			t.Errorf("%v should not be equal to nil", r)
		}
	})

}

func TestIntRange_String(t *testing.T) {

	r := NewIntRange(5, 10)
	actual := r.String()
	expected := "[5..10]"
	if expected != actual {
		t.Errorf("expected %v, got %v", expected, actual)
	}

}

func TestIntRange_Split(t *testing.T) {

	t.Run("split even range", func(t *testing.T) {
		r := NewIntRange(2, 8)
		r1, r2 := r.Split()
		e1 := NewIntRange(2, 4)
		e2 := NewIntRange(5, 8)
		if !e1.Equals(r1) {
			t.Errorf("expected %v, got %v", e1, r1)
		}
		if !e2.Equals(r2) {
			t.Errorf("expected %v, got %v", e2, r2)
		}
	})

	t.Run("split odd range", func(t *testing.T) {
		r := NewIntRange(2, 7)
		r1, r2 := r.Split()
		e1 := NewIntRange(2, 4)
		e2 := NewIntRange(5, 7)
		if !e1.Equals(r1) {
			t.Errorf("expected %v, got %v", e1, r1)
		}
		if !e2.Equals(r2) {
			t.Errorf("expected %v, got %v", e2, r2)
		}
	})

	t.Run("split range with singe element", func(t *testing.T) {
		r := NewIntRange(5, 5)
		r1, r2 := r.Split()
		e1 := NewIntRange(5, 5)
		e2 := NewIntRange(5, 5)
		if !e1.Equals(r1) {
			t.Errorf("expected %v, got %v", e1, r1)
		}
		if !e2.Equals(r2) {
			t.Errorf("expected %v, got %v", e2, r2)
		}
	})

	t.Run("split inverse range", func(t *testing.T) {
		r := NewIntRange(5, 0)
		r1, r2 := r.Split()
		e1 := NewIntRange(5, 5)
		e2 := NewIntRange(5, 5)
		if !e1.Equals(r1) {
			t.Errorf("expected %v, got %v", e1, r1)
		}
		if !e2.Equals(r2) {
			t.Errorf("expected %v, got %v", e2, r2)
		}
	})

	t.Run("centered range", func(t *testing.T) {
		r := NewIntRange(-5, 5)
		r1, r2 := r.Split()
		e1 := NewIntRange(-5, -1)
		e2 := NewIntRange(0, 5)
		if !e1.Equals(r1) {
			t.Errorf("expected %v, got %v", e1, r1)
		}
		if !e2.Equals(r2) {
			t.Errorf("expected %v, got %v", e2, r2)
		}
	})

	t.Run("large range", func(t *testing.T) {
		r := NewIntRange(math.MinInt32/2, math.MaxInt32/2)
		r1, r2 := r.Split()
		e1 := NewIntRange(-1073741824, -1)
		e2 := NewIntRange(0, 1073741823)
		if !e1.Equals(r1) {
			t.Errorf("expected %v, got %v", e1, r1)
		}
		if !e2.Equals(r2) {
			t.Errorf("expected %v, got %v", e2, r2)
		}
	})

}

func TestIntRange_CanSplit(t *testing.T) {

	t.Run("splittable range", func(t *testing.T) {
		r := NewIntRange(5, 7)
		if !r.CanSplit() {
			t.Errorf("range should be splittable: %v", r)
		}
	})

	t.Run("non-splittable range", func(t *testing.T) {
		r := NewIntRange(5, 5)
		if r.CanSplit() {
			t.Errorf("range should not be splittable: %v", r)
		}
	})

}

func TestIntRange_Min(t *testing.T) {
	r := NewIntRange(0, 10)
	if r.Min() != 0 {
		t.Errorf("expected 0, got %v", r.Min())
	}
}

func TestIntRange_Max(t *testing.T) {
	r := NewIntRange(0, 10)
	if r.Max() != 10 {
		t.Errorf("expected 0, got %v", r.Max())
	}
}

func TestIntRange_InitItem(t *testing.T) {

	r := NewIntRange(0, 10)

	t.Run("init int item", func(t *testing.T) {
		actual := r.InitItem(5)
		expected := NewIntRange(5, 5)
		if !expected.Equals(actual) {
			t.Errorf("expected %v, got %v", expected, actual)
		}
	})

	t.Run("init uint item", func(t *testing.T) {
		actual := r.InitItem(uint(5))
		expected := NewIntRange(5, 5)
		if !expected.Equals(actual) {
			t.Errorf("expected %v, got %v", expected, actual)
		}
	})

	t.Run("init float item", func(t *testing.T) {
		actual := r.InitItem(float64(5))
		expected := NewIntRange(5, 5)
		if !expected.Equals(actual) {
			t.Errorf("expected %v, got %v", expected, actual)
		}
	})

	t.Run("unknown item", func(t *testing.T) {
		actual := r.InitItem("test")
		expected := NewIntRange(0, 0)
		if !expected.Equals(actual) {
			t.Errorf("expected %v, got %v", expected, actual)
		}
	})

}

func TestIntRange_MaxSplit(t *testing.T) {
	tests := []struct {
		min, max int
		maxSplit int
	}{
		{0, 0, 0},
		{0, 1, 1},
		{5, 10, 3},
		{0, 10, 4},
	}
	for _, test := range tests {
		t.Run("range max split", func(t *testing.T) {
			r := NewIntRange(test.min, test.max)
			actual := r.MaxSplit()
			if test.maxSplit != actual {
				t.Errorf("expected %v, got %v", test.maxSplit, actual)
			}
		})
	}
}
