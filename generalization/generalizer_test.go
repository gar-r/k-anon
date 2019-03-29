package generalization

import (
	"bitbucket.org/dargzero/k-anon/partition"
	"bitbucket.org/dargzero/k-anon/testutil"
	"fmt"
	"testing"
)

func TestHierarchyGeneralizer_Generalize(t *testing.T) {

	t.Run("invalid hierarchy", func(t *testing.T) {
		invalid := &Hierarchy{
			Partitions: [][]*partition.ItemSet{
				{
					partition.NewItemSet("A"),
					partition.NewItemSet("B"),
				},
				{
					partition.NewItemSet("C"),
				},
			},
		}
		generalizer := NewHierarchyGeneralizer(invalid)
		testutil.AssertNil(generalizer, t)
	})

	t.Run("invalid value in hierarchy", func(t *testing.T) {
		generalizer := NewHierarchyGeneralizer(GetGradeHierarchy())
		p := partition.NewItemSet("X")
		actual := generalizer.Generalize(p, 1)
		testutil.AssertNil(actual, t)
	})

	t.Run("generalize tests", func(t *testing.T) {
		generalizer := NewHierarchyGeneralizer(GetGradeHierarchy())
		tests := []struct {
			p, expected *partition.ItemSet
			level       int
		}{
			{level: 0, p: partition.NewItemSet("C"), expected: partition.NewItemSet("C")},
			{level: 1, p: partition.NewItemSet("C"), expected: partition.NewItemSet("C+", "C", "C-")},
			{level: 2, p: partition.NewItemSet("C"), expected: partition.NewItemSet("A+", "A", "A-", "B", "B+", "B-", "C+", "C", "C-")},
		}
		for _, test := range tests {
			t.Run(fmt.Sprintf("level %d", test.level), func(t *testing.T) {
				actual := generalizer.Generalize(test.p, test.level)
				if !test.expected.Equals(actual) {
					t.Errorf("partitions are not equal: %v, %v", test.expected, actual)
				}
			})
		}
	})

	t.Run("generalize complex partition", func(t *testing.T) {
		generalizer := NewHierarchyGeneralizer(GetGradeHierarchy())
		p := partition.NewItemSet("C+", "C", "C-")
		actual := generalizer.Generalize(p, 2)
		expected := partition.NewItemSet("A+", "A", "A-", "B", "B+", "B-", "C+", "C", "C-")
		if !expected.Equals(actual) {
			t.Errorf("partitions are not equal: %v, %v", expected, actual)
		}
	})

	t.Run("re-generalize complex partition", func(t *testing.T) {
		generalizer := NewHierarchyGeneralizer(GetGradeHierarchy())
		p := partition.NewItemSet("C+", "C", "C-")
		actual := generalizer.Generalize(p, 0)
		expected := partition.NewItemSet("C+", "C", "C-")
		if !expected.Equals(actual) {
			t.Errorf("partitions are not equal: %v, %v", expected, actual)
		}
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
			p := partition.NewItemSet(test.item)
			actual := g.Generalize(p, test.n)
			if test.expected == nil {
				testutil.AssertNil(actual, t)
			} else {
				expected := partition.NewItemSet(test.expected)
				if !expected.Equals(actual) {
					t.Errorf("partitions are not equal: %v, %v", expected, actual)
				}
			}
		})
	}
}
