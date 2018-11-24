package model

import (
	"k-anon/generalization"
	"testing"
)

func TestData_Generalize(t *testing.T) {
	d := &Data{
		Value:       "dummy",
		generalizer: &stubGeneralizer{"stub"},
	}
	actual := d.Generalize(1)
	expected := "stub"
	if !actual.Contains(expected) {
		t.Errorf("expected generalized partition to contain stub value")
	}
}

func TestVector_Generalize(t *testing.T) {
	v := &Vector{
		Items: []*Data{
			{Value: 1, generalizer: &generalization.Suppressor{}},
			{Value: 2, generalizer: &generalization.Suppressor{}},
			{Value: 3, generalizer: &generalization.Suppressor{}},
		},
	}
	expected := generalization.NewPartition("*")
	for _, item := range v.Items {
		if !expected.Equals(item.Generalize(1)) {
			t.Errorf("expected suppressed partition")
		}
	}
}

type stubGeneralizer struct {
	stubValue interface{}
}

func (g *stubGeneralizer) Generalize(item interface{}, n int) *generalization.Partition {
	return generalization.NewPartition(g.stubValue)
}

func (g *stubGeneralizer) Levels() int {
	return 1
}
