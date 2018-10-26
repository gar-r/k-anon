package generalization

import (
	"fmt"
	"strings"
)

// Partition is a single partition or node in the generalization hierarchy, which contains a set of items.
type Partition struct {
	items map[interface{}]bool
}

// NewPartition creates a new Partition from the given slice of items.
func NewPartition(items ...interface{}) *Partition {
	p := &Partition{items: make(map[interface{}]bool)}
	for _, item := range items {
		p.items[item] = true
	}
	return p
}

// Contains returns true if the given item is part of the Partition
func (p *Partition) Contains(item interface{}) bool {
	return p.items[item]
}

// Combine merges a number of partitions and creates a new partition containing all elements from the input partitions.
func Combine(partitions ...*Partition) *Partition {
	p := NewPartition()
	for _, partition := range partitions {
		for item := range partition.items {
			p.items[item] = true
		}
	}
	return p
}

// Equals compares the Partition to another one and returns true if the elements match.
func (p *Partition) Equals(other *Partition) bool {
	if other == nil || len(p.items) != len(other.items) {
		return false
	}
	for i := range p.items {
		if !other.Contains(i) {
			return false
		}
	}
	return true
}

// String returns the string representation of the partition
func (p *Partition) String() string {
	b := &strings.Builder{}
	for item := range p.items {
		b.WriteString(fmt.Sprintf("%v", item))
		b.WriteString(", ")
	}
	s := strings.Trim(strings.TrimSpace(b.String()), ",")
	return fmt.Sprintf("[%s]", s)
}
