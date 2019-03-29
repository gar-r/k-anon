package partition

import (
	"fmt"
	"strings"
)

type ItemSet struct {
	Items map[interface{}]bool
}

func NewItemSet(items ...interface{}) *ItemSet {
	p := &ItemSet{Items: make(map[interface{}]bool)}
	for _, item := range items {
		p.Items[item] = true
	}
	return p
}

func (p *ItemSet) Contains(item interface{}) bool {
	return p.Items[item]
}

func (p *ItemSet) ContainsPartition(other Partition) bool {
	p2, success := other.(*ItemSet)
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

func (p *ItemSet) Equals(other Partition) bool {
	if other == nil {
		return false
	}
	p2, success := other.(*ItemSet)
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

func (p *ItemSet) String() string {
	b := &strings.Builder{}
	for item := range p.Items {
		b.WriteString(fmt.Sprintf("%v", item))
		b.WriteString(", ")
	}
	s := strings.Trim(strings.TrimSpace(b.String()), ",")
	return fmt.Sprintf("[%s]", s)
}
