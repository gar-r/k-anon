package hierarchy

import (
	"bitbucket.org/dargzero/k-anon/partition"
	"errors"
	"fmt"
)

// Hierarchy is a tree representing a generalization hierarchy.
// The hierarchy contains a partition, in which items are from the same domain.
// A hierarchy is only valid if:
//    - it is a full tree
//    - the partition of a given level contains all partitions of child levels
type Hierarchy interface {
	Levels() int
	Find(p partition.Partition) Hierarchy
	Partition() partition.Partition
	Children() []Hierarchy
	Parent() Hierarchy
}

func Build(partition partition.Partition, children ...Hierarchy) (Hierarchy, error) {
	node := newNode(partition, children)
	if err := validate(node); err != nil {
		return nil, err
	}
	createParentLinks(node)
	return node, nil
}

func N(partition partition.Partition, children ...Hierarchy) Hierarchy {
	return newNode(partition, children)
}

func newNode(partition partition.Partition, children []Hierarchy) *node {
	node := &node{data: partition}
	for _, child := range children {
		childNode := newNode(child.Partition(), child.Children())
		node.children = append(node.children, childNode)
	}
	return node
}

type node struct {
	data     partition.Partition
	children []*node
	parent   *node
}

func (n *node) Levels() int {
	return countLevels(n, 1)
}

func (n *node) Children() []Hierarchy {
	var result []Hierarchy
	for _, child := range n.children {
		result = append(result, child)
	}
	return result
}

func (n *node) Parent() Hierarchy {
	return n.parent
}

func (n *node) Partition() partition.Partition {
	return n.data
}

func (n *node) Find(p partition.Partition) Hierarchy {
	return findPartition(n, p)
}

func findPartition(node *node, p partition.Partition) Hierarchy {
	if node.data.Equals(p) {
		return node
	}
	for _, child := range node.children {
		result := findPartition(child, p)
		if result != nil {
			return result
		}
	}
	return nil
}

func validate(node *node) error {
	if err := validateDepth(node, 1, node.Levels()); err != nil {
		return err
	}
	if err := validatePartitions(node); err != nil {
		return err
	}
	return nil
}

func createParentLinks(node *node) {
	for _, child := range node.children {
		child.parent = node
		createParentLinks(child)
	}
}

func validateDepth(node *node, level, maxLevel int) error {
	if len(node.children) == 0 {
		if level != maxLevel {
			return errors.New("hierarchy must be a full tree")
		}
	}
	for _, child := range node.children {
		if err := validateDepth(child, level+1, maxLevel); err != nil {
			return err
		}
	}
	return nil
}

func validatePartitions(node *node) error {
	for _, child := range node.children {
		if !node.data.ContainsPartition(child.data) {
			return errors.New(fmt.Sprintf("%v does not contain child %v", node.data, child.data))
		}
		if err := validatePartitions(child); err != nil {
			return err
		}
	}
	return nil
}

func countLevels(node *node, level int) int {
	max := level
	for _, child := range node.children {
		l := countLevels(child, level+1)
		if max < l {
			max = l
		}
	}
	return max
}
