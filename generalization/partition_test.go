package generalization

import (
	"fmt"
	"strings"
	"testing"
)

func Test_PartitionEquals(t *testing.T) {
	p1 := NewPartition(1, 2, 3)
	p2 := NewPartition(3, 2, 1)
	assertPartitionEquals(p1, p2, t)
}

func Test_PartitionNotEquals(t *testing.T) {
	p1 := NewPartition(1, 2, 3, 4)
	p2 := NewPartition(1, 2, 3)
	assertPartitionNotEquals(p1, p2, t)
}

func Test_PartitionContains(t *testing.T) {
	tests := []struct {
		name     string
		p        *Partition
		item     interface{}
		contains bool
	}{
		{name: "[1, 2, 3], 3 => true", p: NewPartition(1, 2, 3), item: 3, contains: true},
		{name: "[1, 2, 3], 5 => false", p: NewPartition(1, 2, 3), item: 5, contains: false},
		{name: "['A+', 'B-'], A+ => true", p: NewPartition("A+", "B-"), item: "A+", contains: true},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if test.p.Contains(test.item) != test.contains {
				t.Errorf("%v contains %v should be %v", test.p, test.item, test.contains)
			}
		})
	}
}

func Test_PartitionCombine(t *testing.T) {
	p1 := NewPartition(1, 2, 3, 4)
	p2 := NewPartition(1, 2, 7, 9)
	p3 := NewPartition(8, 5)
	expected := NewPartition(1, 2, 3, 4, 5, 7, 8, 9)
	actual := Combine(p1, p2, p3)
	assertPartitionEquals(expected, actual, t)
}

func Test_PartitionToString(t *testing.T) {

	tests := []struct {
		partition *Partition
		values    []string
	}{
		{NewPartition(), []string{}},
		{NewPartition("a"), []string{"a"}},
		{NewPartition("a", "b"), []string{"a", "b"}},
		{NewPartition("a", "b", "c"), []string{"a", "b", "c"}},
		{NewPartition("a", 2, "c"), []string{"a", "2", "c"}},
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

func assertPartitionEquals(p1, p2 *Partition, t *testing.T) {
	if !p1.Equals(p2) {
		t.Errorf("partitions should be equal: %v, %v", p1, p2)
	}
}

func assertPartitionNotEquals(p1, p2 *Partition, t *testing.T) {
	if p1.Equals(p2) {
		t.Errorf("partitions should be different: %v, %v", p1, p2)
	}
}
