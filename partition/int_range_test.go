package partition

import (
	"bitbucket.org/dargzero/k-anon/testutil"
	"testing"
)

func TestNewIntRange(t *testing.T) {

	t.Run("normal int range", func(t *testing.T) {
		r := NewIntRange(5, 10)
		testutil.AssertEquals(5, r.lower, t)
		testutil.AssertEquals(10, r.upper, t)
	})

	t.Run("inverted int range", func(t *testing.T) {
		r := NewIntRange(5, 0)
		testutil.AssertEquals(5, r.lower, t)
		testutil.AssertEquals(0, r.upper, t)
	})

}

func TestIntRange_Contains(t *testing.T) {

	r := NewIntRange(5, 10)

	t.Run("item less than lower", func(t *testing.T) {
		if r.Contains(4) {
			t.Errorf("%v should not contain %v", r, 4)
		}
	})

	t.Run("item equals lower", func(t *testing.T) {
		if !r.Contains(5) {
			t.Errorf("%v should contain %v", r, 5)
		}
	})

	t.Run("item between bounds", func(t *testing.T) {
		if !r.Contains(6) {
			t.Errorf("%v should contain %v", r, 6)
		}
	})

	t.Run("item equals upper", func(t *testing.T) {
		if r.Contains(10) {
			t.Errorf("%v should not contain %v", r, 10)
		}
	})

	t.Run("item greater than upper", func(t *testing.T) {
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

	t.Run("lower subset", func(t *testing.T) {
		r2 := NewIntRange(5, 8)
		if !r.ContainsPartition(r2) {
			t.Errorf("%v should contain %v", r, r2)
		}
	})

	t.Run("upper subset", func(t *testing.T) {
		r2 := NewIntRange(8, 10)
		if r.ContainsPartition(r2) {
			t.Errorf("%v should not contain %v", r, r2)
		}
	})

	t.Run("item set contained", func(t *testing.T) {
		itemSet := NewItemSet(6, 7, 8)
		if !r.ContainsPartition(itemSet) {
			t.Errorf("%v should contain %v", r, itemSet)
		}
	})

	t.Run("item set not contained", func(t *testing.T) {
		itemSet := NewItemSet(6, 7, 8, 11)
		if r.ContainsPartition(itemSet) {
			t.Errorf("%v should not contain %v", r, itemSet)
		}
	})

}

func TestIntRange_Equals(t *testing.T) {

	r := NewIntRange(5, 10)

	t.Run("non int range input", func(t *testing.T) {
		other := NewItemSet()
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
	expected := "[5..10)"
	if expected != actual {
		t.Errorf("expected %v, got %v", expected, actual)
	}

}
