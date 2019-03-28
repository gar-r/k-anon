package generalization

import (
	"bitbucket.org/dargzero/k-anon/testutil"
	"fmt"
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
	p := NewPartition("X")
	actual := generalizer.Generalize(p, 1)
	testutil.AssertNil(actual, t)
}

func Test_HierarchyGeneralizer_Level0(t *testing.T) {
	generalizer := NewHierarchyGeneralizer(GetGradeHierarchy1())
	p := NewPartition("C")
	actual := generalizer.Generalize(p, 0)
	assertPartitionEquals(p, actual, t)
}

func Test_HierarchyGeneralizer_Level1(t *testing.T) {
	generalizer := NewHierarchyGeneralizer(GetGradeHierarchy1())
	p := NewPartition("C")
	actual := generalizer.Generalize(p, 1)
	expected := NewPartition("C+", "C", "C-")
	assertPartitionEquals(expected, actual, t)
}

func Test_HierarchyGeneralizer_Level2(t *testing.T) {
	generalizer := NewHierarchyGeneralizer(GetGradeHierarchy1())
	p := NewPartition("C")
	actual := generalizer.Generalize(p, 2)
	expected := NewPartition("A+", "A", "A-", "B", "B+", "B-", "C+", "C", "C-")
	assertPartitionEquals(expected, actual, t)
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
