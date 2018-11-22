package algorithm

import (
	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/simple"
	"gonum.org/v1/gonum/graph/topo"
	"math"
)

type Decomposer struct {
	g           *simple.UndirectedGraph
	k           int
	vertexCount int
}

func NewDecomposer(g *simple.UndirectedGraph, k int) *Decomposer {
	return &Decomposer{g: g, k: k, vertexCount: g.Nodes().Len()}
}

func (d *Decomposer) Decompose() {
	for {
		c := d.pickComponent()
		if c == nil {
			break
		}
		d.partitionComponent(c)
	}
}

func (d *Decomposer) pickComponent() []graph.Node {
	components := topo.ConnectedComponents(d.g)
	threshold := d.getThreshold()
	for _, c := range components {
		if len(c) > threshold {
			return c
		}
	}
	return nil
}

func (d *Decomposer) partitionComponent(component []graph.Node) {
	u, v, t := d.getSplitParams(component)
	s := len(component)
	if t >= d.k && s-t >= d.k {
		d.splitTypeA(u, v)
	} else if s-t == d.k-1 {
		d.splitTypeB(u, v)
	} else if t == d.k-1 {
		d.splitTypeC(u, v)
	} else {
		var comp1, comp2 []graph.Node
		subTrees := getSubTrees(d.g, u)
		for _, subTree := range subTrees {
			if len(comp1) < d.k-1 {
				comp1 = append(comp1, subTree...)
			} else {
				comp2 = append(comp2, subTree...)
			}
		}
		if len(comp1) == d.k-1 {
			d.cutEdgesToComponent(comp2, u)
		} else if len(comp2) == d.k-1 {
			d.cutEdgesToComponent(comp1, u)
		} else {

		}
	}
}

func (d *Decomposer) cutEdgesToComponent(component []graph.Node, node graph.Node) {
	for _, n := range component {
		d.g.RemoveEdge(node.ID(), n.ID())
	}
}

func (d *Decomposer) splitTypeA(u graph.Node, v graph.Node) {
	d.g.RemoveEdge(u.ID(), v.ID())
}

func (d *Decomposer) splitTypeB(u graph.Node, v graph.Node) {
	sv := d.g.NewNode()
	edges := d.g.From(v.ID())
	for edges.Next() {
		n := edges.Node()
		if u.ID() != n.ID() {
			d.g.RemoveEdge(v.ID(), n.ID())
		}
		d.g.NewEdge(sv, n)
	}
}

func (d *Decomposer) splitTypeC(u graph.Node, v graph.Node) {
	d.splitTypeB(v, u)
}

func (d *Decomposer) getSplitParams(component []graph.Node) (graph.Node, graph.Node, int) {
	u := pickRandomVertex(component)
	for {
		largest := d.getLargestComponent(getSubTrees(d.g, u))
		v := d.getNextRootCandidate(largest, u)
		if len(component)-len(largest) >= d.k-1 {
			return u, v, len(largest)
		}
		u = v
	}
}

func (d *Decomposer) getNextRootCandidate(component []graph.Node, root graph.Node) graph.Node {
	for _, v := range component {
		if d.g.HasEdgeBetween(root.ID(), v.ID()) {
			return d.g.Node(v.ID())
		}
	}
	panic("no edge between root candidate and largest sub-tree")
}

func (d *Decomposer) getThreshold() int {
	threshold := 2*d.k - 1
	if 3*d.k-5 > threshold {
		threshold = 3*d.k - 5
	}
	return threshold
}

func (d *Decomposer) getLargestComponent(components [][]graph.Node) []graph.Node {
	max := math.MinInt64
	var result []graph.Node
	for _, c := range components {
		size := d.calculateSize(c)
		if size > max {
			max = size
			result = c
		}
	}
	return result
}

// calculates the component size, skipping Steiner's vertices
func (d *Decomposer) calculateSize(component []graph.Node) int {
	count := 0
	for _, n := range component {
		if int64(d.vertexCount) > n.ID() {
			count++
		}
	}
	return count
}
