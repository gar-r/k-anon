package generalization

import "testing"

func Test_EmptyBuilder(t *testing.T) {
	builder := &IntegerHierarchyBuilder{}
	builder.Build()
}
