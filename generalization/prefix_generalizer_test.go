package generalization

import (
	"bitbucket.org/dargzero/k-anon/testutil"
	"testing"
)

func TestPrefixGeneralizer_Generalize(t *testing.T) {

	g := &PrefixGeneralizer{MaxWords: 5}

	t.Run("level 0", func(t *testing.T) {
		actual := g.Generalize("this is a test string", 0)
		expected := NewPartition("this is a test string")
		assertPartitionEquals(expected, actual, t)
	})

	t.Run("generalize one level", func(t *testing.T) {
		actual := g.Generalize("this is a test string", 1)
		expected := NewPartition("this is a test")
		assertPartitionEquals(expected, actual, t)
	})

	t.Run("generalize multiple levels", func(t *testing.T) {
		actual := g.Generalize("this is a test string", 3)
		expected := NewPartition("this is")
		assertPartitionEquals(expected, actual, t)
	})

	t.Run("exceeds max words", func(t *testing.T) {
		actual := g.Generalize("this is a test string", 20)
		testutil.AssertNil(actual, t)
	})

	t.Run("generalize last word", func(t *testing.T) {
		actual := g.Generalize("this is a test string", 5)
		expected := NewPartition("")
		assertPartitionEquals(expected, actual, t)
	})

	t.Run("empty string input", func(t *testing.T) {
		actual := g.Generalize("", 3)
		expected := NewPartition("")
		assertPartitionEquals(expected, actual, t)
	})

	t.Run("non string input", func(t *testing.T) {
		actual := g.Generalize(10, 3)
		expected := NewPartition("")
		assertPartitionEquals(expected, actual, t)
	})

	t.Run("fg", func(t *testing.T) {
		actual := g.Generalize("cats are wild", 1)
		expected := NewPartition("cats are wild")
		assertPartitionEquals(expected, actual, t)
	})

}

func TestPrefixGeneralizer_Levels(t *testing.T) {
	g := &PrefixGeneralizer{MaxWords: 10}
	testutil.AssertEquals(10, g.Levels(), t)
}
