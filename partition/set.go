package partition

import (
	"fmt"
	"sort"
	"strings"
)

type Set struct {
	Items map[interface{}]bool
}

func NewSet(items ...interface{}) *Set {
	p := &Set{Items: make(map[interface{}]bool)}
	for _, item := range items {
		p.Items[item] = true
	}
	return p
}

func (p *Set) Contains(item interface{}) bool {
	return p.Items[item]
}

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
