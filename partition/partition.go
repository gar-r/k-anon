package partition

import (
	"sort"
)

// Partition is a single partition or node in a generalization hierarchy
type Partition interface {

	// Contains returns true if the partition contains the given item
	Contains(item interface{}) bool

	// ContainsPartition returns true if this partition contains the given partition
	ContainsPartition(other Partition) bool

	// Equals returns true of this partition is equal to the other partition
	Equals(other Partition) bool

	// String returns a string representation of the partition
	String() string
}

func (p *Set) isIntSeries() bool {
	var items []int
	for item := range p.Items {
		val, success := item.(int)
		if !success {
			return false
		}
		items = append(items, val)
	}
	sort.Ints(items)
	for i := 1; i < len(items); i++ {
		if items[i]-items[i-1] > 1 {
			return false
		}
	}
	return true
}
