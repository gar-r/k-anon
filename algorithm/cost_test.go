package algorithm

import (
	"fmt"
	"k-anon/generalization"
	"k-anon/model"
	"testing"
)

func TestCalculateCost(t *testing.T) {
	tests := []struct {
		items1, items2 []int
		expectedCost   float64
	}{
		{items1: []int{8}, items2: []int{8}, expectedCost: 0},
		{items1: []int{1}, items2: []int{1}, expectedCost: 0},
		{items1: []int{1}, items2: []int{2}, expectedCost: 0.5},
		{items1: []int{8}, items2: []int{9}, expectedCost: 0.25},
		{items1: []int{1}, items2: []int{4}, expectedCost: 0.75},
		{items1: []int{1}, items2: []int{5}, expectedCost: 1},
		{items1: []int{1, 1}, items2: []int{3, 6}, expectedCost: 1.75},
		{items1: []int{1, 1}, items2: []int{1, 1}, expectedCost: 0},
		{items1: []int{1, 1}, items2: []int{1, 1}, expectedCost: 0},
		{items1: []int{1, 1, 1, 1}, items2: []int{4, 5, 1, 2}, expectedCost: 2.25},
	}
	g := getExampleGeneralizer()
	for i, test := range tests {
		t.Run(fmt.Sprintf("Test #%d", i), func(t *testing.T) {
			v1 := createVector(test.items1, g)
			v2 := createVector(test.items2, g)
			actualCost := CalculateCost(v1, v2)
			if test.expectedCost != actualCost {
				t.Errorf("Expected cost: %v, got %v", test.expectedCost, actualCost)
			}
		})
	}
}

func createVector(items []int, g generalization.Generalizer) *model.Vector {
	v := &model.Vector{}
	for _, item := range items {
		v.Items = append(v.Items, model.NewData(item, g))
	}
	return v
}

// Level 4: (1, 2, 3, 4, 5, 6, 7, 8, 9)
// Level 3: (1, 2, 3, 4) (5, 6, 7, 8, 9)
// Level 2: (1, 2) (3, 4) (5, 6) (7, 8, 9)
// Level 1: (1) (2) (3) (4)	(5) (6) (7) (8, 9)
// Level 0: (1) (2) (3) (4)	(5) (6) (7) (8) (9)
func getExampleGeneralizer() generalization.Generalizer {
	builder := generalization.IntegerHierarchyBuilder{
		Items: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
	}
	return generalization.NewHierarchyGeneralizer(builder.NewIntegerHierarchy())
}
