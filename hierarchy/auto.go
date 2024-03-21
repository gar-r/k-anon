package hierarchy

import (
	"errors"
	"math"

	"github.com/gar-r/k-anon/partition"
)

// AutoBuild generates a Hierarchy based on a given set of items and the nChildren
// child count parameter. On the first level of the hierarchy each item will be
// in the same partition and on each subsequent level existing partitions are
// split into nChildren smaller partitions.
func AutoBuild(nChildren int, items ...interface{}) (Hierarchy, error) {
	if len(items) < nChildren {
		return nil, errors.New("items size too small")
	}
	levels := calcLevels(nChildren, len(items))
	children := autoBuild(nChildren, 1, levels, items...)
	return Build(partition.NewSet(items...),
		children...)
}

func autoBuild(split, level, levels int, items ...interface{}) []Hierarchy {
	result := make([]Hierarchy, 0)
	parts := chop(split, items...)
	for _, p := range parts {
		var n Hierarchy
		if level == levels {
			n = N(partition.NewSet(p...))
		} else {
			n = N(partition.NewSet(p...), autoBuild(split, level+1, levels, p...)...)
		}
		result = append(result, n)
	}
	return result
}

func calcLevels(k int, n int) int {
	return int(math.Round(math.Log(float64(n*(k-1)+1)) / math.Log(float64(k))))
}

func chop(split int, items ...interface{}) [][]interface{} {
	result := make([][]interface{}, 0)
	chunkSize := chunkSize(split, items...)
	for i := 0; i < len(items); i += chunkSize {
		if i+chunkSize >= len(items) {
			result = append(result, items[i:])
		} else {
			result = append(result, items[i:i+chunkSize])
		}
	}
	return result
}

func chunkSize(split int, items ...interface{}) int {
	if split >= len(items) {
		return 1
	}
	return int(math.Ceil(float64(len(items)) / float64(split)))
}
