package generalization

import (
	"testing"
)

func Test_EmptyBuilder(t *testing.T) {
	builder := &IntegerHierarchyBuilder{
		Items: []int{2, 4, 6, 8, 1, 3, 5, 7, 9},
	}
	builder.Build() // TODO: tests
}
