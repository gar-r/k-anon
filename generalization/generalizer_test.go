package generalization

import (
	"bitbucket.org/dargzero/k-anon/testutil"
	"fmt"
	"testing"
)

func TestHierarchyGeneralizer_Generalize(t *testing.T) {

	t.Run("invalid hierarchy", func(t *testing.T) {
		invalid := &Hierarchy{
			Partitions: [][]*Partition{
				{
					NewPartition("A"),
					NewPartition("B"),
				},
				{
					NewPartition("C"),
				},
			},
		}
		generalizer := NewHierarchyGeneralizer(invalid)
		testutil.AssertNil(generalizer, t)
	})

	t.Run("invalid value in hierarchy", func(t *testing.T) {
		generalizer := NewHierarchyGeneralizer(GetGradeHierarchy())
		p := NewPartition("X")
		actual := generalizer.Generalize(p, 1)
		testutil.AssertNil(actual, t)
	})

	t.Run("generalize tests", func(t *testing.T) {
		generalizer := NewHierarchyGeneralizer(GetGradeHierarchy())
		tests := []struct {
			p, expected *Partition
			level       int
		}{
			{level: 0, p: NewPartition("C"), expected: NewPartition("C")},
			{level: 1, p: NewPartition("C"), expected: NewPartition("C+", "C", "C-")},
			{level: 2, p: NewPartition("C"), expected: NewPartition("A+", "A", "A-", "B", "B+", "B-", "C+", "C", "C-")},
		}
		for _, test := range tests {
			t.Run(fmt.Sprintf("level %d", test.level), func(t *testing.T) {
				actual := generalizer.Generalize(test.p, test.level)
				assertPartitionEquals(test.expected, actual, t)
			})
		}
	})

	t.Run("generalize complex partition", func(t *testing.T) {
		generalizer := NewHierarchyGeneralizer(GetGradeHierarchy())
		p := NewPartition("C+", "C", "C-")
		actual := generalizer.Generalize(p, 1)
		expected := NewPartition("A+", "A", "A-", "B", "B+", "B-", "C+", "C", "C-")
		assertPartitionEquals(expected, actual, t)
	})
}

func Test_SuppressorPartitionLength(t *testing.T) {
	g := &Suppressor{}
	p := NewPartition("test")
	actual := g.Generalize(p, 1)
	testutil.AssertEquals(1, len(actual.items), t)
}

func Test_Suppressor(t *testing.T) {
	tests := []struct {
		item     interface{}
		n        int
		expected interface{}
	}{
		{"test", 1, "*"},
		{"test", 0, "test"},
		{"test", -1, nil},
		{"test", 2, nil},
	}
	g := &Suppressor{}
	for _, test := range tests {
		t.Run(fmt.Sprintf("%v, %d => %v", test.item, test.n, test.expected), func(t *testing.T) {
			p := NewPartition(test.item)
			actual := g.Generalize(p, test.n)
			if test.expected == nil {
				testutil.AssertNil(actual, t)
			} else {
				exp := NewPartition(test.expected)
				assertPartitionEquals(exp, actual, t)
			}
		})
	}
}
