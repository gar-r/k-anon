package generalization

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

// Equal compares the Partition to another one and returns true if the elements match.
func (p *Partition) Equal(other *Partition) bool {
	if len(p.items) != len(other.items) {
		return false
	}
	for i := range p.items {
		if !other.items[i] {
			return false
		}
	}
	return true
}
