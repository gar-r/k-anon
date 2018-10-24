package generalization

// Partition is a single partition or node in the generalization hierarchy, which contains a set of items.
type Partition struct {
	items map[interface{}]bool
}

// NewPartition creates a new Partition from the given slice of items
func NewPartition(items ...interface{}) Partition {
	p := Partition{items: make(map[interface{}]bool)}
	for _, item := range items {
		p.items[item] = true
	}
	return p
}

// Contains returns true if the given item is part of the Partition
func (p Partition) Contains(item interface{}) bool {
	return p.items[item]
}

// Equal compares the Partition to another one and returns true if the elements match
func (p Partition) Equal(other Partition) bool {
	for i := range p.items {
		if p.items[i] != other.items[i] {
			return false
		}
	}
	return true
}
