package model

import "k-anon/generalization"

type Vector struct {
	Items []*Data
}

type Data struct {
	Value       interface{}
	generalizer generalization.Generalizer
}

func (d *Data) Generalize(level int) *generalization.Partition {
	return d.generalizer.Generalize(d.Value, level)
}
