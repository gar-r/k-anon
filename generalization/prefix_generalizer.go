package generalization

import (
	"bitbucket.org/dargzero/k-anon/partition"
	"strings"
)

// PrefixGeneralizer can be used to generalize plain text.
// The entered text will be interpreted as a series of words, MaxWords in total. All whitespace
// characters will be discarded, and converted to spaces.
// The text to be generalized can contain fewer words than MaxWords, but it will still be
// considered to contain MaxWords number of words in regards to the level of generalization.
// MaxWords should be chosen with considerations on performance. Choosing a very large value
// for MaxWords for 'convenience' will result in degraded performance.
//
// Example: using a prefix generalizer with MaxWords = 5, the string "cats are wild" will be
// interpreted as "cats are wild _ _", and a generalization of level 3 will yield "cats are".
type PrefixGeneralizer struct {
	MaxWords int
}

func (g *PrefixGeneralizer) Generalize(p partition.Partition, n int) partition.Partition {
	const separator = " "
	if n > g.Levels() {
		return nil
	}
	itemSet, success := p.(*partition.ItemSet)
	if !success {
		return partition.NewItemSet("")
	}
	s := stringify(itemSet)
	if n == g.MaxWords || s == "" {
		return partition.NewItemSet("*")
	}
	words := g.getPaddedWords(s)
	idx := g.MaxWords - n
	joined := strings.Join(words[:idx], separator)
	return partition.NewItemSet(strings.TrimRight(joined, separator))
}

func (g *PrefixGeneralizer) Levels() int {
	return g.MaxWords + 1
}

func (g *PrefixGeneralizer) getPaddedWords(s string) []string {
	words := strings.Fields(s)
	padded := make([]string, g.MaxWords)
	copy(padded, words)
	return padded
}

func stringify(p *partition.ItemSet) string {
	for item := range p.Items {
		s, success := item.(string)
		if success {
			return s
		}
	}
	return ""
}
