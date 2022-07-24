package generalization

import (
	"math"
	"strconv"
	"testing"

	"git.okki.hu/garric/k-anon/partition"
	"git.okki.hu/garric/k-anon/testutil"
)

func TestNewIntRangeGeneralizer(t *testing.T) {

	t.Run("single item range", func(t *testing.T) {
		gen := NewIntRangeGeneralizer(10, 10)
		r, success := gen.r.(*partition.IntRange)
		if !success {
			t.Errorf("type assertion failed")
		}
		if r.Min() != 10 && r.Max() != 10 {
			t.Errorf("expected [10..10], got %v", r)
		}
	})

	t.Run("normal range", func(t *testing.T) {
		gen := NewIntRangeGeneralizer(0, 10)
		r, success := gen.r.(*partition.IntRange)
		if !success {
			t.Errorf("type assertion failed")
		}
		if r.Min() != 0 && r.Max() != 10 {
			t.Errorf("expected [0..10], got %v", r)
		}
	})

	t.Run("inverted range", func(t *testing.T) {
		gen := NewIntRangeGeneralizer(10, 5)
		r, success := gen.r.(*partition.IntRange)
		if !success {
			t.Errorf("type assertion failed")
		}
		if r.Min() != 10 && r.Max() != 10 {
			t.Errorf("expected [10..10], got %v", r)
		}
	})
}

func TestNewFloatRangeGeneralizer(t *testing.T) {

	t.Run("single range", func(t *testing.T) {
		gen := NewFloatRangeGeneralizer(0.34, 0.34)
		r, success := gen.r.(*partition.FloatRange)
		if !success {
			t.Errorf("type assertion failed")
		}
		if r.Min() != 0.34 && r.Max() != 0.34 {
			t.Errorf("expected (0.34..0.34), got %v", r)
		}
	})

	t.Run("normal range", func(t *testing.T) {
		gen := NewFloatRangeGeneralizer(0.34, 1.67)
		r, success := gen.r.(*partition.FloatRange)
		if !success {
			t.Errorf("type assertion failed")
		}
		if r.Min() != 0.34 && r.Max() != 1.67 {
			t.Errorf("expected (0.34..1.67), got %v", r)
		}
	})

	t.Run("inverted range", func(t *testing.T) {
		gen := NewFloatRangeGeneralizer(0.34, 0.12)
		r, success := gen.r.(*partition.FloatRange)
		if !success {
			t.Errorf("type assertion failed")
		}
		if r.Min() != 0.34 && r.Max() != 0.34 {
			t.Errorf("expected (0.34..0.34), got %v", r)
		}
	})

}

func TestRangeGeneralizer_Generalize(t *testing.T) {

	t.Run("tests with int range", func(t *testing.T) {

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

	})

	t.Run("tests with float range", func(t *testing.T) {
		g := NewFloatRangeGeneralizer(0.379, 1.932)

		t.Run("generalize level 0", func(t *testing.T) {
			p := g.InitItem(0.45)
			actual := g.Generalize(p, 0)
			expected := partition.NewFloatRange(0.45, 0.45)
			if !expected.Equals(actual) {
				t.Errorf("expected %v, got %v", expected, actual)
			}
		})

		t.Run("generalize to max level", func(t *testing.T) {
			p := g.InitItem(0.45)
			actual := g.Generalize(p, g.Levels()-1)
			expected := partition.NewFloatRange(0.379, 1.932)
			if !expected.Equals(actual) {
				t.Errorf("expected %v, got %v", expected, actual)
			}
		})

		t.Run("generalize above max level", func(t *testing.T) {
			p := g.InitItem(0.45)
			actual := g.Generalize(p, g.Levels())
			if actual != nil {
				t.Errorf("expected nil, got %v", actual)
			}
		})

	})

}

func TestRangeGeneralizer_Levels(t *testing.T) {

	t.Run("returns max split + 1", func(t *testing.T) {
		g := &RangeGeneralizer{r: TestRange{maxSplit: 10}}
		actual := g.Levels()
		expected := 10 + 1
		if expected != actual {
			t.Errorf("expected %v, got %v", expected, actual)
		}
	})

}

func TestRangeGeneralizer_InitItem(t *testing.T) {
	testRange := TestRange{}
	g := &RangeGeneralizer{r: testRange}
	actual := g.InitItem(nil)
	if testRange != actual {
		t.Errorf("expected %v, got %v", testRange, actual)
	}
}

func BenchmarkIntRangeGeneralizer(b *testing.B) {
	sizes := []int{10, 100, 1000, 10000, 100000, 1000000, 10000000, 100000000, 1000000000}
	for _, size := range sizes {
		b.Run(strconv.Itoa(size), func(b *testing.B) {
			benchmarkIntRangeGeneralizer(size, b)
		})
	}
}

func benchmarkIntRangeGeneralizer(max int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		g := NewIntRangeGeneralizer(0, max)
		g.Generalize(g.InitItem(max/2), g.Levels())
	}
}

func BenchmarkFloatRangeGeneralizer(b *testing.B) {
	sizes := []float64{10, 100, 1000, 10000, 100000, 1000000, 10000000, 100000000, 1000000000}
	for _, size := range sizes {
		b.Run(strconv.FormatFloat(size, 'g', 0, 64), func(b *testing.B) {
			benchmarkFloatRangeGeneralizer(size, b)
		})
	}
}

func benchmarkFloatRangeGeneralizer(max float64, b *testing.B) {
	for i := 0; i < b.N; i++ {
		g := NewFloatRangeGeneralizer(0, max)
		g.Generalize(g.InitItem(max/2), g.Levels())
	}
}

type TestRange struct {
	maxSplit int
}

func (TestRange) Contains(item interface{}) bool {
	return false
}

func (TestRange) ContainsPartition(other partition.Partition) bool {
	return false
}

func (TestRange) Equals(other partition.Partition) bool {
	return false
}

func (TestRange) String() string {
	return ""
}

func (TestRange) Min() float64 {
	return 0
}

func (TestRange) Max() float64 {
	return 0
}

func (TestRange) CanSplit() bool {
	return false
}

func (TestRange) Split() (partition.Range, partition.Range) {
	return nil, nil
}

func (t TestRange) MaxSplit() int {
	return t.maxSplit
}

func (t TestRange) InitItem(item interface{}) partition.Range {
	return t
}
