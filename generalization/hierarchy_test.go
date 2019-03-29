package generalization

import (
	"bitbucket.org/dargzero/k-anon/partition"
	"bitbucket.org/dargzero/k-anon/testutil"
	"strings"
	"testing"
)

func Test_Levels(t *testing.T) {
	h := GetGradeHierarchy()
	expected := 3
	actual := h.GetLevelCount()
	testutil.AssertEquals(expected, actual, t)
}

func Test_GetLevel(t *testing.T) {
	h := GetGradeHierarchy()
	expected := []*partition.ItemSet{partition.NewItemSet("A+", "A", "A-", "B+", "B", "B-", "C+", "C", "C-")}
	actual := h.GetLevel(2)
	testutil.AssertEquals(len(expected), len(actual), t)
	for i := range expected {
		if !expected[i].Equals(actual[i]) {
			t.Errorf("element mismatch at index %d", i)
		}
	}
}

func Test_Valid(t *testing.T) {
	h := GetGradeHierarchy()
	assertValid(h, t)
}

func Test_InvalidMultipleValuesOnLevel(t *testing.T) {
	h := &Hierarchy{Partitions: [][]*partition.ItemSet{
		{
			partition.NewItemSet(1, 2, 3),
			partition.NewItemSet(5, 6, 3), // <= error: 3 is present in both partitions in the same level
		},
	}}
	assertInvalid(h, t)
}

func Test_InvalidItemsDoNotAddUp(t *testing.T) {
	h := &Hierarchy{Partitions: [][]*partition.ItemSet{
		{
			partition.NewItemSet(1),
			partition.NewItemSet(2),
			partition.NewItemSet(3),
			partition.NewItemSet(4),
		},
		{
			partition.NewItemSet(1, 2),
			partition.NewItemSet(3, 5), // <= error: 5 is not part of the hierarchy
		},
		{
			partition.NewItemSet(1, 2, 3, 4),
		},
	}}
	assertInvalid(h, t)
}

func Test_GetLevelUnderIndex(t *testing.T) {
	h := GetGradeHierarchy()
	idx := -1
	actual := h.GetLevel(idx)
	testutil.AssertNil(actual, t)
}

func Test_GetLevelOverIndex(t *testing.T) {
	h := GetGradeHierarchy()
	idx := h.GetLevelCount() // max index + 1
	actual := h.GetLevel(idx)
	testutil.AssertNil(actual, t)
}

func Test_StringEmpty(t *testing.T) {
	h := &Hierarchy{}
	expected := ""
	actual := h.String()
	testutil.AssertEquals(expected, actual, t)
}

func Test_StringSinglePartition(t *testing.T) {
	p := partition.NewItemSet("a", "b")
	h := &Hierarchy{Partitions: [][]*partition.ItemSet{{p}}}
	actual := h.String()
	expected1 := "[a, b]"
	expected2 := "[b, a]"
	if !strings.Contains(actual, expected1) && !strings.Contains(actual, expected2) {
		t.Errorf("expected %s to contain partition %s", actual, p.String())
	}
}

func assertValid(h *Hierarchy, t *testing.T) {
	if !h.IsValid() {
		t.Errorf("hierarchy should be valid")
	}
}

func assertInvalid(h *Hierarchy, t *testing.T) {
	if h.IsValid() {
		t.Errorf("hierarchy should be invalid")
	}
}

func assertHierarchyEquals(h1 *Hierarchy, h2 *Hierarchy, t *testing.T) {
	if !Equals(h1, h2) {
		t.Errorf("expected:\n%s\nactual:\n%s\n", h1, h2)
	}
}
