package generalization

import (
	"fmt"
	"k-anon/testutil"
	"testing"
)

func Test_InvalidHierarchy(t *testing.T) {
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
}

func Test_InvalidValueForHierarchy(t *testing.T) {
	generalizer := NewHierarchyGeneralizer(GetGradeHierarchy1())
	actual := generalizer.Generalize("X", 1)
	testutil.AssertNil(actual, t)
}

func Test_HierarchyGeneralizer_Level0(t *testing.T) {
	generalizer := NewHierarchyGeneralizer(GetGradeHierarchy1())
	actual := generalizer.Generalize("C", 0)
	expected := NewPartition("C")
	assertPartitionEquals(expected, actual, t)
}

func Test_HierarchyGeneralizer_Level1(t *testing.T) {
	generalizer := NewHierarchyGeneralizer(GetGradeHierarchy1())
	actual := generalizer.Generalize("C", 1)
	expected := NewPartition("C+", "C", "C-")
	assertPartitionEquals(expected, actual, t)
}

func Test_HierarchyGeneralizer_Level2(t *testing.T) {
	generalizer := NewHierarchyGeneralizer(GetGradeHierarchy1())
	actual := generalizer.Generalize("C", 2)
	expected := NewPartition("A+", "A", "A-", "B", "B+", "B-", "C+", "C", "C-")
	assertPartitionEquals(expected, actual, t)
}

func Test_SuppressorPartitionLength(t *testing.T) {
	g := &Suppressor{}
	p := g.Generalize("test", 1)
	testutil.AssertEquals(1, len(p.items), t)
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
			p := g.Generalize(test.item, test.n)
			if test.expected == nil {
				testutil.AssertNil(p, t)
			} else {
				exp := NewPartition(test.expected)
				assertPartitionEquals(exp, p, t)
			}
		})
	}
}
