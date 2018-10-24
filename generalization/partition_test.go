package generalization

import "testing"

func Test_PartitionEquals(t *testing.T) {
	p1 := NewPartition(1, 2, 3)
	p2 := NewPartition(1, 2, 3)
	if !p1.Equal(p2) {
		t.Errorf("Expected partitions to be equal")
	}
}

func Test_PartitionNotEquals(t *testing.T) {
	p1 := NewPartition(1, 2, 3)
	p2 := NewPartition(1, 2, 2)
	if p1.Equal(p2) {
		t.Errorf("Expected partitions to differ")
	}
}

func Test_PartitionContains(t *testing.T) {
	tests := []struct {
		name     string
		p        Partition
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
				t.Errorf("Expected partition ")
			}
		})
	}
}
