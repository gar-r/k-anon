package generalization

import "testing"

func Test_ExampleHierarchyLevels(t *testing.T) {
	h := getExampleHierarchy()
	expected := 3
	actual := h.Levels()
	if expected != actual {
		t.Errorf("Expected %d levels, but got %d", expected, actual)
	}
}

func Test_ExampleHierarchyGetLevel(t *testing.T) {
	h := getExampleHierarchy()
	expected := []*Partition{NewPartition("A+", "A", "A-", "B+", "B", "B-", "C+", "C", "C-")}
	actual := h.GetLevel(2)
	if len(expected) != len(actual) {
		t.Errorf("Partition sizes do not match")
	}
	for i := range expected {
		if !expected[i].Equal(actual[i]) {
			t.Errorf("Element mismatch at index %d", i)
		}
	}
}

func getExampleHierarchy() *Hierarchy {
	return &Hierarchy{Partitions: [][]*Partition{
		{
			NewPartition("A+"),
			NewPartition("A"),
			NewPartition("A-"),
			NewPartition("B+"),
			NewPartition("B"),
			NewPartition("B-"),
			NewPartition("C+"),
			NewPartition("C"),
			NewPartition("C-"),
		},
		{
			NewPartition("A+", "A", "A-"),
			NewPartition("B+", "B", "B-"),
			NewPartition("C+", "C", "C-"),
		},
		{
			NewPartition("A+", "A", "A-", "B+", "B", "B-", "C+", "C", "C-"),
		},
	}}
}
