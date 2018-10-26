package generalization

import (
	"strings"
	"testing"
)

func Test_Levels(t *testing.T) {
	h := getExampleHierarchy()
	expected := 3
	actual := h.GetLevelCount()
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
		if !expected[i].Equals(actual[i]) {
			t.Errorf("Element mismatch at index %d", i)
		}
	}
}

func Test_Valid(t *testing.T) {
	h := getExampleHierarchy()
	if !h.IsValid() {
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
	if h.IsValid() {
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
	if h.IsValid() {
		t.Errorf("This hierarchy should be invalid")
	}
}

func Test_Find(t *testing.T) {
	tests := []struct {
		name     string
		item     interface{}
		level    int
		expected *Partition
	}{
		{"Exists GetLevel 0", "C", 0, NewPartition("C")},
		{"Exists GetLevel 1", "C", 1, NewPartition("C+", "C", "C-")},
		{"Exists GetLevel 2", "C", 2, NewPartition("A+", "A", "A-", "B+", "B", "B-", "C+", "C", "C-")},
		{"Missing GetLevel 0", "X", 0, nil},
		{"Missing GetLevel 1", "X", 1, nil},
		{"Missing GetLevel 2", "X", 2, nil},
	}
	h := getExampleHierarchy()
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := h.Find(test.item, test.level)
			if test.expected == nil {
				if actual != nil {
					t.Errorf("Missing item was found in partition %v", actual)
				}
			} else if !test.expected.Equals(actual) {
				t.Errorf("Item was not located in the correct partition: %v", actual)
			}
		})
	}
}

func Test_GetLevelUnderIndex(t *testing.T) {
	h := getExampleHierarchy()
	idx := -1
	actual := h.GetLevel(idx)
	if nil != actual {
		t.Errorf("Expected nil, but got %v", actual)
	}
}

func Test_GetLevelOverIndex(t *testing.T) {
	h := getExampleHierarchy()
	idx := h.GetLevelCount() // max index + 1
	actual := h.GetLevel(idx)
	if nil != actual {
		t.Errorf("Expected nil, but got %v", actual)
	}
}

func Test_FindUnderIndex(t *testing.T) {
	h := getExampleHierarchy()
	idx := -1
	actual := h.Find("C", idx)
	if nil != actual {
		t.Errorf("Expected nil, but got %v", actual)
	}
}

func Test_FindOverIndex(t *testing.T) {
	h := getExampleHierarchy()
	idx := h.GetLevelCount() // max index + 1
	actual := h.Find("C", idx)
	if nil != actual {
		t.Errorf("Expected nil, but got %v", actual)
	}
}

func Test_StringEmpty(t *testing.T) {
	h := &Hierarchy{}
	expected := ""
	actual := h.String()
	if expected != actual {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}

func Test_StringSinglePartition(t *testing.T) {
	p := NewPartition(1, 2)
	h := &Hierarchy{Partitions: [][]*Partition{{p}}}
	actual := h.String()
	chunk := p.String()
	if !strings.Contains(actual, chunk) {
		t.Errorf("Expected %s to contain partition %s", actual, chunk)
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
