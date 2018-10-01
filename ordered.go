package main

type Ordered interface {
	Less(other Ordered) bool
}

type Int int

func (i Int) Less(other Ordered) bool {
	j, ok := other.(Int)
	if !ok {
		return false
	}
	return i < j
}
