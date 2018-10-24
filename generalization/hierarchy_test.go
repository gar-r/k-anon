package generalization

import "testing"

func Test_Levels(t *testing.T) {
	h := getExampleHierarchy()
	expected := 3
	actual := h.Levels()
	if expected != actual {
		t.Errorf("Expected %d levels, but got %d", expected, actual)
	}
}

func Test_GetLevel(t *testing.T) {
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

func Test_Valid(t *testing.T) {
	h := getExampleHierarchy()
	if !h.Valid() {
		t.Errorf("Example hierarchy should be valid")
	}
}

func Test_InvalidMultipleValuesOnLevel(t *testing.T) {
	h := &Hierarchy{Partitions: [][]*Partition{
		{
			NewPartition(1, 2, 3),
			NewPartition(5, 6, 3), // <= error: 3 is present in both partitions in the same level
		},
	}}
	if h.Valid() {
		t.Errorf("This hierarchy should be invalid")
	}
}

func Test_InvalidItemsDoNotAddUp(t *testing.T) {
	h := &Hierarchy{Partitions: [][]*Partition{
		{
			NewPartition(1),
			NewPartition(2),
			NewPartition(3),
			NewPartition(4),
		},
		{
			NewPartition(1, 2),
			NewPartition(3, 5), // <= error: 5 is not part of the hierarchy
		},
		{
			NewPartition(1, 2, 3, 4),
		},
	}}
	if h.Valid() {
		t.Errorf("This hierarchy should be invalid")
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
