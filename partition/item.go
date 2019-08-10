package partition

import "fmt"

// Item represents a partition from a single items.
type Item struct {
	item interface{}
}

// NewItem creates a new instance of Item from the given item.
func NewItem(item interface{}) *Item {
	return &Item{item: item}
}

// Contains returns true when the partition contains the given item.
func (p *Item) Contains(item interface{}) bool {
	return p.item == item
}

// ContainsPartition is always false in case of Item partitions.
func (p *Item) ContainsPartition(other Partition) bool {
	return false
}

// Equals returns true when this partition is equal to the other partition.
func (p *Item) Equals(other Partition) bool {
	q, success := other.(*Item)
	if !success {
		return false
	}
	return p.item == q.item
}

// String returns the string representation of the partition.
func (p *Item) String() string {
	return fmt.Sprintf("%v", p.item)
}

// GetItem returns the item contained in the partition.
func (p *Item) GetItem() interface{} {
	return p.item
}
