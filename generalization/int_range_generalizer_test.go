package generalization

import (
	"bitbucket.org/dargzero/k-anon/partition"
	"bitbucket.org/dargzero/k-anon/testutil"
	"math"
	"testing"
)

func TestIntRangeGeneralizer_Generalize(t *testing.T) {

	g := NewIntRangeGeneralizer(5, 10)

	t.Run("generalize level 0", func(t *testing.T) {
		p := partition.NewIntRange(6, 7)
		actual := g.Generalize(p, 0)
		expected := partition.NewIntRange(6, 7)
		if !expected.Equals(actual) {
			t.Errorf("expected %v, got %v", expected, actual)
		}
	})

	t.Run("generalize level 1", func(t *testing.T) {
		p := partition.NewIntRange(6, 7)
		actual := g.Generalize(p, 1)
		expected := partition.NewIntRange(6, 7)
		if !expected.Equals(actual) {
			t.Errorf("expected %v, got %v", expected, actual)
		}
	})

	t.Run("generalize level 2", func(t *testing.T) {
		p := partition.NewIntRange(6, 7)
		actual := g.Generalize(p, 2)
		expected := partition.NewIntRange(5, 7)
		if !expected.Equals(actual) {
			t.Errorf("expected %v, got %v", expected, actual)
		}
	})

	t.Run("generalize level 3", func(t *testing.T) {
		p := partition.NewIntRange(6, 7)
		actual := g.Generalize(p, 3)
		expected := partition.NewIntRange(5, 10)
		if !expected.Equals(actual) {
			t.Errorf("expected %v, got %v", expected, actual)
		}
	})

	t.Run("cannot generalize further", func(t *testing.T) {
		p := partition.NewIntRange(6, 7)
		actual := g.Generalize(p, 4)
		testutil.AssertNil(actual, t)
	})

	t.Run("incompatible type", func(t *testing.T) {
		p := partition.NewSet()
		actual := g.Generalize(p, 0)
		testutil.AssertNil(actual, t)
	})

	t.Run("not found in path", func(t *testing.T) {
		p := partition.NewIntRange(3, 7)
		actual := g.Generalize(p, 0)
		testutil.AssertNil(actual, t)
	})

	t.Run("continue generalization", func(t *testing.T) {
		p := partition.NewIntRange(5, 7)
		actual := g.Generalize(p, 3)
		expected := partition.NewIntRange(5, 10)
		if !expected.Equals(actual) {
			t.Errorf("expected %v, got %v", expected, actual)
		}
	})

	t.Run("large range", func(t *testing.T) {
		gen := NewIntRangeGeneralizer(0, math.MaxInt64-1)
		p := partition.NewIntRange(5, 6)
		actual := gen.Generalize(p, 32)
		expected := partition.NewIntRange(0, 4294967294)
		if !expected.Equals(actual) {
			t.Errorf("expected %v, got %v", expected, actual)
		}
	})

	t.Run("tiny range", func(t *testing.T) {
		gen := NewIntRangeGeneralizer(0, 2)
		p := partition.NewIntRange(0, 0)
		actual := gen.Generalize(p, 1)
		expected := partition.NewIntRange(0, 0)
		if !expected.Equals(actual) {
			t.Errorf("expected %v, got %v", expected, actual)
		}
	})
}

func TestIntRangeGeneralizer_Levels(t *testing.T) {

	t.Run("singular range", func(t *testing.T) {
		gen := NewIntRangeGeneralizer(0, 0)
		actual := gen.Levels()
		expected := 1
		if expected != actual {
			t.Errorf("expected %v, got %v", expected, actual)
		}
	})

	t.Run("test range", func(t *testing.T) {
		gen := NewIntRangeGeneralizer(5, 10)
		actual := gen.Levels()
		expected := 4
		if expected != actual {
			t.Errorf("expected %v, got %v", expected, actual)
		}
	})

	t.Run("small range", func(t *testing.T) {
		gen := NewIntRangeGeneralizer(0, 10)
		actual := gen.Levels()
		expected := 5
		if expected != actual {
			t.Errorf("expected %v, got %v", expected, actual)
		}
	})

	t.Run("large range", func(t *testing.T) {
		gen := NewIntRangeGeneralizer(0, math.MaxInt64-1)
		actual := gen.Levels()
		expected := 64
		if expected != actual {
			t.Errorf("expected %v, got %v", expected, actual)
		}
	})

}
