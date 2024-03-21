package generalization

import (
	"strings"

	"github.com/gar-r/k-anon/partition"
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

// Generalize generalizes the partition n levels further and returns the resulting partition.
func (g *PrefixGeneralizer) Generalize(p partition.Partition, n int) partition.Partition {
	const separator = " "
	if n > g.Levels() {
		return nil
	}
	item, success := p.(*partition.Item)
	if !success {
		return g.InitItem("")
	}
	s := stringify(item)
	if n == g.MaxWords || s == "" {
		return g.InitItem("*")
	}
	words := g.getPaddedWords(s)
	idx := g.MaxWords - n
	joined := strings.Join(words[:idx], separator)
	return g.InitItem(strings.TrimRight(joined, separator))
}

// Levels returns the maximum levels of the generalizer, in this case MaxWords+1.
func (g *PrefixGeneralizer) Levels() int {
	return g.MaxWords + 1
}

// InitItem wraps the item in an Item partition.
func (g *PrefixGeneralizer) InitItem(item interface{}) partition.Partition {
	return partition.NewItem(item)
}

func (g *PrefixGeneralizer) getPaddedWords(s string) []string {
	words := strings.Fields(s)
	if len(words) > g.MaxWords {
		words = words[:g.MaxWords]
	}
	padded := make([]string, g.MaxWords)
	copy(padded, words)
	return padded
}

func stringify(p *partition.Item) string {
	s, success := p.GetItem().(string)
	if success {
		return s
	}
	return ""
}
