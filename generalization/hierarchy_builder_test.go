package generalization

import (
	"testing"
)

func Test_IntegerHierarchyBuilderValidity(t *testing.T) {
	h := buildHierarchy(2, 4, 6, 8, 1, 3, 5, 7, 9)
	if !h.IsValid() {
		t.Errorf("Hierarchy built with builder should be valid")
	}
}

func Test_IntegerHierarchyBuilderEmptySet(t *testing.T) {
	actual := buildHierarchy() // empty
	expected := &Hierarchy{}
	assertEquals(t, expected, actual)
}
func Test_IntegerHierarchyBuilderSingleItem(t *testing.T) {
	actual := buildHierarchy(1)
	expected := &Hierarchy{
		Partitions: [][]*Partition{
			{NewPartition(1)},
		},
	}
	assertEquals(t, expected, actual)
}

func Test_IntegerHierarchyBuilderLevel1(t *testing.T) {
	actual := buildHierarchy(1, 2)
	expected := &Hierarchy{
		Partitions: [][]*Partition{
			{NewPartition(1), NewPartition(2)},
			{NewPartition(1, 2)},
		},
	}
	assertEquals(t, expected, actual)
}

func Test_IntegerHierarchyBuilderPartitionCutAtMedian(t *testing.T) {
	actual := buildHierarchy(1, 2, 3)
	expected := &Hierarchy{
		Partitions: [][]*Partition{
			{NewPartition(1), NewPartition(2), NewPartition(3)},
			{NewPartition(1), NewPartition(2, 3)},
			{NewPartition(1, 2, 3)},
		},
	}
	assertEquals(t, expected, actual)
}

func assertEquals(t *testing.T, expected *Hierarchy, actual *Hierarchy) {
	if !equals(expected, actual) {
		t.Errorf("Expected:\n%s\nActual:\n%s\n", expected, actual)
	}
}

func equals(h1, h2 *Hierarchy) bool {
	if h1.GetLevelCount() != h2.GetLevelCount() {
		return false
	}
	for i := 0; i < h1.GetLevelCount(); i++ {
		l1 := h1.GetLevel(i)
		l2 := h2.GetLevel(i)
	partition:
		for _, p1 := range l1 {
			for _, p2 := range l2 {
				if p1.Equals(p2) {
					continue partition
				}
			}
			return false
		}
	}
	return true
}

func buildHierarchy(items ...int) *Hierarchy {
	builder := &IntegerHierarchyBuilder{
		Items: items,
	}
	return builder.Build()
}
