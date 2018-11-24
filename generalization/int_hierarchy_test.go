package generalization

import (
	"testing"
)

func Test_IntegerHierarchyBuilderValidity(t *testing.T) {
	h := buildIntHierarchy(2, 4, 6, 8, 1, 3, 5, 7, 9)
	if !h.IsValid() {
		t.Errorf("Hierarchy built with builder should be valid")
	}
}

func Test_IntegerHierarchyBuilderEmptySet(t *testing.T) {
	actual := buildIntHierarchy() // empty
	expected := &Hierarchy{}
	assertHierarchyEquals(expected, actual, t)
}
func Test_IntegerHierarchyBuilderSingleItem(t *testing.T) {
	actual := buildIntHierarchy(1)
	expected := &Hierarchy{
		Partitions: [][]*Partition{
			{NewPartition(1)},
		},
	}
	assertHierarchyEquals(expected, actual, t)
}

func Test_IntegerHierarchyBuilderLevel1(t *testing.T) {
	actual := buildIntHierarchy(1, 2)
	expected := &Hierarchy{
		Partitions: [][]*Partition{
			{NewPartition(1), NewPartition(2)},
			{NewPartition(1, 2)},
		},
	}
	assertHierarchyEquals(expected, actual, t)
}

func Test_IntegerHierarchyBuilderPartitionCutAtMedian(t *testing.T) {
	actual := buildIntHierarchy(1, 2, 3)
	expected := &Hierarchy{
		Partitions: [][]*Partition{
			{NewPartition(1), NewPartition(2), NewPartition(3)},
			{NewPartition(1), NewPartition(2, 3)},
			{NewPartition(1, 2, 3)},
		},
	}
	assertHierarchyEquals(expected, actual, t)
}
