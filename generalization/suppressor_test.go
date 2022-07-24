package generalization

import (
	"fmt"
	"testing"

	"git.okki.hu/garric/k-anon/partition"
	"git.okki.hu/garric/k-anon/testutil"
)

func TestSuppressor_Levels(t *testing.T) {
	g := &Suppressor{}
	actual := g.Levels()
	expected := 2
	if expected != actual {
		t.Errorf("expected %v, got %v", expected, actual)
	}
}

func TestSuppressor_Generalize(t *testing.T) {
	tests := []struct {
		item     interface{}
		n        int
		expected interface{}
	}{
		{"test", 1, "*"},
		{"test", 0, "test"},
		{"test", -1, nil},
		{"test", 2, nil},
	}
	g := &Suppressor{}
	for _, test := range tests {
		t.Run(fmt.Sprintf("%v, %d => %v", test.item, test.n, test.expected), func(t *testing.T) {
			p := partition.NewItem(test.item)
			actual := g.Generalize(p, test.n)
			if test.expected == nil {
				testutil.AssertNil(actual, t)
			} else {
				expected := partition.NewItem(test.expected)
				if !expected.Equals(actual) {
					t.Errorf("partitions are not equal: %v, %v", expected, actual)
				}
			}
		})
	}
}
