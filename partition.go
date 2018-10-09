package main

import (
	"fmt"
	"sort"
)

type Partition struct {
	Items []Ordered
}

func FromInts(ints ...int) *Partition {
	p := &Partition{Items: []Ordered{}}
	for _, i := range ints {
		p.Items = append(p.Items, Int(i))
	}
	return p
}

func (p *Partition) Len() int {
	return len(p.Items)
}

func (p *Partition) Less(i, j int) bool {
	return p.Items[i].Less(p.Items[j])
}

func (p *Partition) Swap(i, j int) {
	p.Items[i], p.Items[j] = p.Items[j], p.Items[i]
}

func (p *Partition) Split() (p1, p2 *Partition) {
	sort.Sort(p)
	count := p.Len()

	medianIdx := count / 2
	p1 = &Partition{Items: p.Items[0:medianIdx]}
	p2 = &Partition{Items: p.Items[medianIdx:count]}
	return
}

func (p *Partition) String() string {
	return fmt.Sprint(p.Items)
}

func Equal(p1 *Partition, p2 *Partition) bool {
	a := p1.Items
	b := p2.Items
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
