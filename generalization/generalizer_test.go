package generalization

import (
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
	if nil != generalizer {
		t.Errorf("Expected nil, but got %v", generalizer)
	}
}

func Test_HierarchyGeneralizer_Level1(t *testing.T) {
	generalizer := NewHierarchyGeneralizer(getExampleHierarchy())
	actual := generalizer.Generalize("C", 1)
	expected := NewPartition("C+", "C", "C-")
	if !expected.Equal(actual) {
		t.Errorf("Expected partition %v, got %v", expected, actual)
	}
}

func Test_HierarchyGeneralizer_Level2(t *testing.T) {
	generalizer := NewHierarchyGeneralizer(getExampleHierarchy())
	actual := generalizer.Generalize("C", 2)
	expected := NewPartition("A+", "A", "A-", "B", "B+", "B-", "C+", "C", "C-")
	if !expected.Equal(actual) {
		t.Errorf("Expected partition %v, got %v", expected, actual)
	}
}

func Test_StringGeneralizer(t *testing.T) {
	tests := []struct {
		input    string
		n        int
		expected interface{}
	}{
		{"Greenland", 0, "Greenland"},
		{"Greenland", 4, "Green"},
		{"Greenland", 8, "G"},
		{"Greenland", 9, "*"},
		{"Greenland", 10, nil},
		{"Greenland", -1, nil},
	}
	generalizer := &StringGeneralizer{}
	for _, test := range tests {
		t.Run(fmt.Sprintf("%s, %d", test.input, test.n), func(t *testing.T) {
			p := generalizer.Generalize(test.input, test.n)
			if test.expected == nil {
				if p != nil {
					t.Errorf("Expected nil, but got %v", p)
				}
			} else {
				exp := NewPartition(test.expected)
				if !exp.Equal(p) {
					t.Errorf("Expected %v, got %v", exp, p)
				}
			}
		})
	}
}
