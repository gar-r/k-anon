package partition

import (
	"fmt"
	"strings"
	"testing"
)

func TestItemSet_Equals(t *testing.T) {

	p1 := NewSet(1, 2, 3)

	t.Run("item sets are equal", func(t *testing.T) {
		p2 := NewSet(3, 2, 1)
		if !p1.Equals(p2) {
			t.Errorf("partitions are not equal: %v, %v", p1, p2)
		}
	})

	t.Run("item sets are different", func(t *testing.T) {
		p2 := NewSet(1, 2, 3, 4)
		if p1.Equals(p2) {
			t.Errorf("partitions should not be equal: %v, %v", p1, p2)
		}
	})

	t.Run("nil input", func(t *testing.T) {
		if p1.Equals(nil) {
			t.Errorf("partition should not be equal to nil")
		}
	})

}

func TestItemSet_Contains(t *testing.T) {
	tests := []struct {
		name     string
		p        *Set
		item     interface{}
		contains bool
	}{
		{name: "[1, 2, 3], 3 => true", p: NewSet(1, 2, 3), item: 3, contains: true},
		{name: "[1, 2, 3], 5 => false", p: NewSet(1, 2, 3), item: 5, contains: false},
		{name: "['A+', 'B-'], A+ => true", p: NewSet("A+", "B-"), item: "A+", contains: true},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if test.p.Contains(test.item) != test.contains {
				t.Errorf("%v contains %v should be %v", test.p, test.item, test.contains)
			}
		})
	}
}

func TestItemSet_ContainsPartition(t *testing.T) {

	p1 := NewSet(1, 2, 3, 4, 5, 6)

	t.Run("non compatible type", func(t *testing.T) {
		p2 := NewIntRange(0, 5)
		if p1.ContainsPartition(p2) {
			t.Errorf("%v should not contain %v", p1, p2)
		}
	})

	t.Run("does not contain all items", func(t *testing.T) {
		p2 := NewSet(2, 3, 8)
		if p1.ContainsPartition(p2) {
			t.Errorf("%v should not contain %v", p1, p2)
		}
	})

	t.Run("contains all items", func(t *testing.T) {
		p2 := NewSet(2, 3, 5, 6)
		if !p1.ContainsPartition(p2) {
			t.Errorf("%v should contain %v", p1, p2)
		}
	})

}

func TestItemSet_String(t *testing.T) {
	tests := []struct {
		partition *Set
		values    []string
	}{
		{NewSet(), []string{}},
		{NewSet("a"), []string{"a"}},
		{NewSet("a", "b"), []string{"a", "b"}},
		{NewSet("a", "b", "c"), []string{"a", "b", "c"}},
		{NewSet("a", 2, "c"), []string{"a", "2", "c"}},
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
