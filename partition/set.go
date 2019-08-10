package partition

import (
	"fmt"
	"sort"
	"strings"
)

// Set represents a bundle of items in which each item is unique.
type Set struct {
	Items map[interface{}]bool
}

// NewSet creates a new instance of a Set from the given values.
func NewSet(items ...interface{}) *Set {
	p := &Set{Items: make(map[interface{}]bool)}
	for _, item := range items {
		p.Items[item] = true
	}
	return p
}

// Contains returns true when the set contains a given value.
func (p *Set) Contains(item interface{}) bool {
	return p.Items[item]
}

// ContainsPartition returns true when the set contains the other partition.
// Note, that the other partition must be a set as well, otherwise this method always returns false.
func (p *Set) ContainsPartition(other Partition) bool {
	p2, success := other.(*Set)
	if !success {
		return false
	}
	for item := range p2.Items {
		if !p.Contains(item) {
			return false
		}
	}
	return true
}

// Equals returns true when the set equals to the other partition.
// Note, that the other partition must be a set as well, otherwise this method always returns false.
func (p *Set) Equals(other Partition) bool {
	p2, success := other.(*Set)
	if !success || len(p2.Items) != len(p.Items) {
		return false
	}
	for i := range p.Items {
		if !p2.Contains(i) {
			return false
		}
	}
	return true
}

// String returns the string representation of the set, listing each value in lexicographical order.
func (p *Set) String() string {
	b := &strings.Builder{}
	var items []string
	for item := range p.Items {
		items = append(items, fmt.Sprintf("%v", item))
	}
	sort.Strings(items)
	for _, s := range items {
		b.WriteString(s)
		b.WriteString(", ")
	}
	s := strings.Trim(strings.TrimSpace(b.String()), ",")
	return fmt.Sprintf("[%s]", s)
}
