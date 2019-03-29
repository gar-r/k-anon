package generalization

import (
	"fmt"
	"strings"
	"testing"
)

func TestItemSet_Equals(t *testing.T) {

	t.Run("item sets are equal", func(t *testing.T) {
		p1 := NewItemSet(1, 2, 3)
		p2 := NewItemSet(3, 2, 1)
		assertPartitionEquals(p1, p2, t)
	})

	t.Run("item sets are different", func(t *testing.T) {
		p1 := NewItemSet(1, 2, 3, 4)
		p2 := NewItemSet(1, 2, 3)
		assertPartitionNotEquals(p1, p2, t)
	})

}

func TestItemSet_Contains(t *testing.T) {
	tests := []struct {
		name     string
		p        *ItemSet
		item     interface{}
		contains bool
	}{
		{name: "[1, 2, 3], 3 => true", p: NewItemSet(1, 2, 3), item: 3, contains: true},
		{name: "[1, 2, 3], 5 => false", p: NewItemSet(1, 2, 3), item: 5, contains: false},
		{name: "['A+', 'B-'], A+ => true", p: NewItemSet("A+", "B-"), item: "A+", contains: true},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if test.p.Contains(test.item) != test.contains {
				t.Errorf("%v contains %v should be %v", test.p, test.item, test.contains)
			}
		})
	}
}

func TestItemSet_String(t *testing.T) {

	tests := []struct {
		partition *ItemSet
		values    []string
	}{
		{NewItemSet(), []string{}},
		{NewItemSet("a"), []string{"a"}},
		{NewItemSet("a", "b"), []string{"a", "b"}},
		{NewItemSet("a", "b", "c"), []string{"a", "b", "c"}},
		{NewItemSet("a", 2, "c"), []string{"a", "2", "c"}},
	}
	for i, test := range tests {
		t.Run(fmt.Sprintf("PartitionToString #%d", i), func(t *testing.T) {
			actual := test.partition.String()
			for _, v := range test.values {
				if !strings.Contains(actual, v) {
					t.Errorf("expected %s to contain %s", actual, v)
				}
			}
		})
	}
}

func assertPartitionEquals(p1, p2 Partition, t *testing.T) {
	if !p1.Equals(p2) {
		t.Errorf("partitions should be equal: %v, %v", p1, p2)
	}
}

func assertPartitionNotEquals(p1, p2 Partition, t *testing.T) {
	if p1.Equals(p2) {
		t.Errorf("partitions should be different: %v, %v", p1, p2)
	}
}
