package main

import (
	"fmt"
	"sort"
)

type Partition struct {
	items []Ordered
}

func FromInts(ints ...int) *Partition {
	p := &Partition{items: []Ordered{}}
	for _, i := range ints {
		p.items = append(p.items, Int(i))
	}
	return p
}

func (p *Partition) Len() int {
	return len(p.items)
}

func (p *Partition) Less(i, j int) bool {
	return p.items[i].Less(p.items[j])
}

func (p *Partition) Swap(i, j int) {
	p.items[i], p.items[j] = p.items[j], p.items[i]
}

func (p *Partition) Split() (p1, p2 *Partition) {
	sort.Sort(p)
	count := p.Len()

	medianIdx := count / 2
	p1 = &Partition{items: p.items[0:medianIdx]}
	p2 = &Partition{items: p.items[medianIdx:count]}
	return
}

func (p *Partition) String() string {
	return fmt.Sprint(p.items)
}

func Equal(p1 *Partition, p2 *Partition) bool {
	a := p1.items
	b := p2.items
	if (a == nil) != (b == nil) {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
