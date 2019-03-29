package generalization

import (
	"bitbucket.org/dargzero/k-anon/testutil"
	"fmt"
	"testing"
)

func TestHierarchyGeneralizer_Generalize(t *testing.T) {

	t.Run("invalid hierarchy", func(t *testing.T) {
		invalid := &Hierarchy{
			Partitions: [][]*ItemSet{
				{
					NewItemSet("A"),
					NewItemSet("B"),
				},
				{
					NewItemSet("C"),
				},
			},
		}
		generalizer := NewHierarchyGeneralizer(invalid)
		testutil.AssertNil(generalizer, t)
	})

	t.Run("invalid value in hierarchy", func(t *testing.T) {
		generalizer := NewHierarchyGeneralizer(GetGradeHierarchy())
		p := NewItemSet("X")
		actual := generalizer.Generalize(p, 1)
		testutil.AssertNil(actual, t)
	})

	t.Run("generalize tests", func(t *testing.T) {
		generalizer := NewHierarchyGeneralizer(GetGradeHierarchy())
		tests := []struct {
			p, expected *ItemSet
			level       int
		}{
			{level: 0, p: NewItemSet("C"), expected: NewItemSet("C")},
			{level: 1, p: NewItemSet("C"), expected: NewItemSet("C+", "C", "C-")},
			{level: 2, p: NewItemSet("C"), expected: NewItemSet("A+", "A", "A-", "B", "B+", "B-", "C+", "C", "C-")},
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
		p := NewItemSet("C+", "C", "C-")
		actual := generalizer.Generalize(p, 2)
		expected := NewItemSet("A+", "A", "A-", "B", "B+", "B-", "C+", "C", "C-")
		assertPartitionEquals(expected, actual, t)
	})

	t.Run("re-generalize complex partition", func(t *testing.T) {
		generalizer := NewHierarchyGeneralizer(GetGradeHierarchy())
		p := NewItemSet("C+", "C", "C-")
		actual := generalizer.Generalize(p, 0)
		expected := NewItemSet("C+", "C", "C-")
		assertPartitionEquals(expected, actual, t)
	})
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
			p := NewItemSet(test.item)
			actual := g.Generalize(p, test.n)
			if test.expected == nil {
				testutil.AssertNil(actual, t)
			} else {
				exp := NewItemSet(test.expected)
				assertPartitionEquals(exp, actual, t)
			}
		})
	}
}
