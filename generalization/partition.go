package generalization

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

// Partition is a single partition or node in a generalization hierarchy
type Partition interface {
	Contains(item interface{}) bool
	ContainsPartition(other Partition) bool
	Equals(other Partition) bool
	String() string
}

type ItemSet struct {
	items map[interface{}]bool
}

// NewItemSet creates a new ItemSet from the given slice of items
func NewItemSet(items ...interface{}) *ItemSet {
	p := &ItemSet{items: make(map[interface{}]bool)}
	for _, item := range items {
		p.items[item] = true
	}
	return p
}

// Contains returns true if the given item is part of the ItemSet
func (p *ItemSet) Contains(item interface{}) bool {
	return p.items[item]
}

// ContainsPartition returns true if the given partition is contained by this partition
func (p *ItemSet) ContainsPartition(other Partition) bool {
	p2, success := other.(*ItemSet)
	if !success {
		return false
	}
	for item := range p2.items {
		if !p.Contains(item) {
			return false
		}
	}
	return true
}

// Equals compares the ItemSet to another one and returns true if the elements match.
func (p *ItemSet) Equals(other Partition) bool {
	if other == nil {
		return false
	}
	p2, success := other.(*ItemSet)
	if !success || len(p2.items) != len(p.items) {
		return false
	}
	for i := range p.items {
		if !p2.Contains(i) {
			return false
		}
	}
	return true
}

// String returns the string representation of the partition
func (p *ItemSet) String() string {
	if len(p.items) < 1 {
		return ""
	}
	if len(p.items) > 1 && p.isIntSeries() {
		return p.intRangeString()
	}
	return p.itemsListString()
}

func (p *ItemSet) itemsListString() string {
	b := &strings.Builder{}
	for item := range p.items {
		b.WriteString(fmt.Sprintf("%v", item))
		b.WriteString(", ")
	}
	s := strings.Trim(strings.TrimSpace(b.String()), ",")
	return fmt.Sprintf("[%s]", s)
}

// Treats the partition data as int and prints the string representation of the range
// If there is an error during number conversion, it will return an error
func (p *ItemSet) intRangeString() string {
	min := math.MaxInt64
	max := math.MinInt64
	for item := range p.items {
		num, _ := item.(int)
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
	}
	return fmt.Sprintf("[%d..%d]", min, max)
}

func (p *ItemSet) isIntSeries() bool {
	var items []int
	for item := range p.items {
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
