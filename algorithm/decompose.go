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
	originalLen int
}

func NewDecomposer(g *simple.UndirectedGraph, k int) *Decomposer {
	size := 0
	if g.Nodes() != nil && g.Nodes().Len() > 0 {
		size = g.Nodes().Len()
	}
	return &Decomposer{g: g, k: k, originalLen: size}
}

func (d *Decomposer) Decompose() {
	threshold := d.getThreshold()
	for {
		c := d.pickComponent(threshold)
		if c == nil {
			break
		}
		d.partitionComponent(c)
	}
}

func (d *Decomposer) pickComponent(threshold int) []graph.Node {
	components := topo.ConnectedComponents(d.g)
	for _, c := range components {
		if d.calculateSize(c) > threshold {
			return c
		}
	}
	return nil
}

func (d *Decomposer) partitionComponent(component []graph.Node) {
	u, v, t := d.getSplitParams(component)
	s := d.calculateSize(component)
	if t >= d.k && s-t >= d.k {
		d.performSplitTypeA(u, v)
	} else if s-t == d.k-1 {
		d.performSplitTypeB(u, v)
	} else if t == d.k-1 {
		d.performSplitTypeC(u, v)
	} else {
		d.performSplitTypeD(u, v, component)
	}
}

func (d *Decomposer) performSplitTypeA(u graph.Node, v graph.Node) {
	d.g.RemoveEdge(u.ID(), v.ID())
}

func (d *Decomposer) performSplitTypeB(u graph.Node, v graph.Node) {
	d.cutSubTrees(v, func(subRoot graph.Node) bool {
		return u.ID() != subRoot.ID()
	})
}

func (d *Decomposer) performSplitTypeC(u graph.Node, v graph.Node) {
	d.performSplitTypeB(v, u)
}

func (d *Decomposer) performSplitTypeD(u graph.Node, v graph.Node, component []graph.Node) {
	p1, p2 := d.getPartitions(u, component)
	if d.calculateSize(p2) == d.k-1 {
		d.cutSubTrees(u, func(subRoot graph.Node) bool {
			return containsNode(p1, subRoot)
		})
	} else {
		d.cutSubTrees(u, func(subRoot graph.Node) bool {
			return containsNode(p2, subRoot)
		})
	}
}

func (d *Decomposer) getPartitions(u graph.Node, component []graph.Node) ([]graph.Node, []graph.Node) {
	var comp1, comp2 []graph.Node
	subTrees := d.getSubTrees(component, u)
	for _, subTree := range subTrees {
		if d.calculateSize(comp1) < d.k-1 {
			comp1 = append(comp1, subTree...)
		} else {
			comp2 = append(comp2, subTree...)
		}
	}
	return comp1, comp2
}

func (d *Decomposer) cutSubTrees(u graph.Node, condition func(subRoot graph.Node) bool) {
	sv := d.g.NewNode() // insert Steiner's vertex for remaining unconnected components
	edges := d.g.From(u.ID())
	for edges.Next() {
		n := edges.Node()
		if condition(n) {
			d.g.RemoveEdge(u.ID(), n.ID())
			d.g.NewEdge(sv, n)
		}
	}
}

func (d *Decomposer) getSplitParams(component []graph.Node) (graph.Node, graph.Node, int) {
	u := pickRandomVertex(component)
	s := d.calculateSize(component)
	for {
		largest := d.getLargestComponent(d.getSubTrees(component, u))
		t := d.calculateSize(largest)
		v := d.getNextRootCandidate(largest, u)
		if s-t >= d.k-1 {
			return u, v, d.calculateSize(largest)
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

func (d *Decomposer) getSubTrees(component []graph.Node, root graph.Node) [][]graph.Node {
	gCopy := simple.NewUndirectedGraph()
	graph.Copy(gCopy, d.g)
	nodes := gCopy.Nodes()
	for nodes.Next() {
		n := nodes.Node()
		if !containsNode(component, n) {
			gCopy.RemoveNode(n.ID())
		}
	}
	gCopy.RemoveNode(root.ID())
	return topo.ConnectedComponents(gCopy)
}

// Calculates the component size, skipping Steiner's vertices.
// As per definition Steiner's vertices don't contribute to the component size.
func (d *Decomposer) calculateSize(component []graph.Node) int {
	count := 0
	for _, n := range component {
		if int64(d.originalLen) > n.ID() {
			count++
		}
	}
	return count
}
