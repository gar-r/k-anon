package generalization

import "sort"

// IntegerHierarchyBuilder can automatically generate a hierarchy from a numeric domain.
// The builder will construct the Hierarchy from top to bottom, splitting partitions into two based on the median value.
// The resulting Hierarchy will always be a valid generalization hierarchy.
type IntegerHierarchyBuilder struct {
	Items []int
}

func (b *IntegerHierarchyBuilder) Build() *Hierarchy {
	b.deduplicate()
	h := &Hierarchy{
		Partitions: [][]*Partition{
			{NewPartition(stripType(b.Items)...)},
		},
	}
	for !refined(h) {
		refine(h)
	}
	return h
}

func refine(h *Hierarchy) {
	level := h.getLevel(0)
	var newPartitions []*Partition
	for _, p := range level {
		if len(p.items) > 1 {
			values := intValues(p)
			p1, p2 := split(values)
			newPartitions = append(newPartitions, NewPartition(stripType(p1)...))
			newPartitions = append(newPartitions, NewPartition(stripType(p2)...))
		} else {
			newPartitions = append(newPartitions, p)
		}
	}
	h.Partitions = append([][]*Partition{newPartitions}, h.Partitions...)
}

func refined(h *Hierarchy) bool {
	level := h.getLevel(0)
	for _, p := range level {
		if len(p.items) != 1 {
			return false
		}
	}
	return true
}

func intValues(p *Partition) []int {
	var values []int
	for item := range p.items {
		i := item.(int)
		values = append(values, i)
	}
	return values
}

func stripType(slice []int) []interface{} {
	var result []interface{}
	for _, item := range slice {
		result = append(result, item)
	}
	return result
}

func (b *IntegerHierarchyBuilder) deduplicate() {
	itemMap := make(map[int]bool)
	for _, item := range b.Items {
		itemMap[item] = true
	}
	b.Items = make([]int, 0, len(itemMap))
	for item := range itemMap {
		b.Items = append(b.Items, item)
	}
}

func split(slice []int) ([]int, []int) {
	m := median(slice)
	var p1 []int
	var p2 []int
	for _, i := range slice {
		if float64(i) < m {
			p1 = append(p1, i)
		} else {
			p2 = append(p2, i)
		}
	}
	return p1, p2
}

func median(slice []int) float64 {
	len := len(slice)
	sort.Ints(slice)
	if len%2 == 1 {
		return float64(slice[len/2])
	}
	return float64(slice[len/2-1]+slice[len/2]) / 2
}
