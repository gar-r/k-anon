package generalization

import (
	"bitbucket.org/dargzero/k-anon/partition"
	"sort"
)

func NewIntGeneralizer(lower, upper, step int) *HierarchyGeneralizer {
	r := makeRange(lower, upper, step)
	return NewIntGeneralizerFromItems(r...)
}

func NewIntGeneralizerFromItems(items ...int) *HierarchyGeneralizer {
	return &HierarchyGeneralizer{
		hierarchy: buildHierarchy(items),
	}
}

func buildHierarchy(items []int) *Hierarchy {
	if len(items) < 1 {
		return &Hierarchy{}
	}
	integers := deduplicate(items)
	h := &Hierarchy{
		Partitions: [][]*partition.ItemSet{
			{partition.NewItemSet(stripType(integers)...)},
		},
	}
	for !refined(h) {
		refine(h)
	}
	return h
}

func refine(h *Hierarchy) {
	level := h.GetLevel(0)
	var newPartitions []*partition.ItemSet
	for _, p := range level {
		if len(p.Items) > 1 {
			values := intValues(p)
			p1, p2 := split(values)
			newPartitions = append(newPartitions, partition.NewItemSet(stripType(p1)...))
			newPartitions = append(newPartitions, partition.NewItemSet(stripType(p2)...))
		} else {
			newPartitions = append(newPartitions, p)
		}
	}
	h.Partitions = append([][]*partition.ItemSet{newPartitions}, h.Partitions...)
}

func refined(h *Hierarchy) bool {
	level := h.GetLevel(0)
	for _, p := range level {
		if len(p.Items) != 1 {
			return false
		}
	}
	return true
}

func intValues(p *partition.ItemSet) []int {
	var values []int
	for item := range p.Items {
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

func deduplicate(items []int) []int {
	itemMap := make(map[int]bool)
	for _, item := range items {
		itemMap[item] = true
	}
	result := make([]int, 0, len(itemMap))
	for item := range itemMap {
		result = append(result, item)
	}
	return result
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

func makeRange(from, to, step int) []int {
	var r []int
	if step < 1 {
		step = 1
	}
	for i := from; i < to; i += step {
		r = append(r, i)
	}
	return r
}
