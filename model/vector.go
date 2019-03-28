package model

import "bitbucket.org/dargzero/k-anon/generalization"

type Vector struct {
	Items []*Data
}

type Data struct {
	Value       *generalization.Partition
	generalizer generalization.Generalizer
}

func NewIdentifier(value interface{}, generalizer generalization.Generalizer) *Data {
	return &Data{
		Value:       generalization.NewPartition(value),
		generalizer: generalizer,
	}
}

func NewIdentifierPartition(p *generalization.Partition, generalizer generalization.Generalizer) *Data {
	return &Data{
		Value:       p,
		generalizer: generalizer,
	}
}

func NewNonIdentifier(value interface{}) *Data {
	return &Data{
		Value:       generalization.NewPartition(value),
		generalizer: nil,
	}
}

func (d *Data) IsIdentifier() bool {
	return d.generalizer != nil
}

func (d *Data) Generalize(level int) *generalization.Partition {
	return d.generalizer.Generalize(d.Value, level)
}

func (d *Data) Levels() int {
	return d.generalizer.Levels()
}
