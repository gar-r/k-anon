package generalization

import (
	"testing"
)

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
