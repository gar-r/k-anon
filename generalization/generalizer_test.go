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

func Test_InvalidValueForHierarchy(t *testing.T) {
	generalizer := NewHierarchyGeneralizer(getExampleHierarchy())
	actual := generalizer.Generalize("X", 1)
	if nil != actual {
		t.Errorf("Expected nil, got %v", actual)
	}
}

func Test_HierarchyGeneralizer_Level1(t *testing.T) {
	generalizer := NewHierarchyGeneralizer(getExampleHierarchy())
	actual := generalizer.Generalize("C", 1)
	expected := NewPartition("C+", "C", "C-")
	if !expected.Equals(actual) {
		t.Errorf("Expected partition %v, got %v", expected, actual)
	}
}

func Test_HierarchyGeneralizer_Level2(t *testing.T) {
	generalizer := NewHierarchyGeneralizer(getExampleHierarchy())
	actual := generalizer.Generalize("C", 2)
	expected := NewPartition("A+", "A", "A-", "B", "B+", "B-", "C+", "C", "C-")
	if !expected.Equals(actual) {
		t.Errorf("Expected partition %v, got %v", expected, actual)
	}
}

func Test_StringGeneralizerPartitionLength(t *testing.T) {
	g := &StringGeneralizer{}
	p := g.Generalize("test", 2)
	if len(p.items) != 1 {
		t.Errorf("Expected partition size to be exactly 1, got %d instead", len(p.items))
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
				if !exp.Equals(p) {
					t.Errorf("Expected %v, got %v", exp, p)
				}
			}
		})
	}
}

func Test_SuppressorPartitionLength(t *testing.T) {
	g := &Suppressor{}
	p := g.Generalize("test", 1)
	if len(p.items) != 1 {
		t.Errorf("Expected partition size to be exactly 1, got %d instead", len(p.items))
	}
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
				if p != nil {
					t.Errorf("Expected nil, got %v", p)
				}
			} else {
				exp := NewPartition(test.expected)
				if !exp.Equals(p) {
					t.Errorf("Expected %v, got %v", exp, p)
				}
			}
		})
	}
}
