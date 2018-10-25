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
	b.sort()
	l := []*Partition{NewPartition(b.Items)}
	return &Hierarchy{}
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

func (b *IntegerHierarchyBuilder) sort() {
	sort.Ints(b.Items)
}

func split(slice []int) (p1, p2 []int) {

}

func median(slice []int) int {
	if len(slice)%2 == 1 {

	}
}
