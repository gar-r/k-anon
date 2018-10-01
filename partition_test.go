package main

import (
	"testing"
)

func TestPartitionSplit(t *testing.T) {
	tests := []struct {
		p, p1, p2 *Partition
	}{
		{p: FromInts(1, 2, 3, 4), p1: FromInts(1, 2), p2: FromInts(3, 4)},
	}
	for _, test := range tests {
		t.Run(test.p.String(), func(t *testing.T) {
			p1, p2 := test.p.Split()
			if !Equal(p1, test.p1) || !Equal(p2, test.p2) {
				t.Errorf("split partitions incorrect: %q, %q", p1, p2)
			}
		})
	}
}

func TestPartitionEquals(t *testing.T) {
	p1 := FromInts(1, 2)
	p2 := FromInts(1, 2)
	if !Equal(p1, p2) {
		t.Errorf("equals error: %q <> %q", p1, p2)
	}
}

func TestStringValue(t *testing.T) {
	p := FromInts(1, 2, 3, 4, 5)
	expected := "[1 2 3 4 5]"
	actual := p.String()
	if expected != actual {
		t.Errorf("expected %q, got %q", expected, actual)
	}
}
