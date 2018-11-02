package algorithm

import (
	"github.com/gyuho/goraph"
	"k-anon/model"
	"strconv"
)

func BuildCoreGraph(t *model.Table) goraph.Graph {
	g := goraph.NewGraph()
	for i := range t.Rows {
		node := goraph.NewNode(strconv.Itoa(i))
		g.AddNode(node)
	}
	return g
}
