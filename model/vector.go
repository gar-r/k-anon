package model

import "k-anon/generalization"

type Vector struct {
	Items []*Data
}

type Data struct {
	Value       interface{}
	generalizer generalization.Generalizer
}

func NewData(value interface{}, generalizer generalization.Generalizer) *Data {
	return &Data{
		Value:       value,
		generalizer: generalizer,
	}
}

func (d *Data) Generalize(level int) *generalization.Partition {
	return d.generalizer.Generalize(d.Value, level)
}

func (d *Data) Levels() int {
	return d.generalizer.Levels()
}
