package generalization

import (
	"bitbucket.org/dargzero/k-anon/partition"
	"bitbucket.org/dargzero/k-anon/testutil"
	"testing"
)

func TestPrefixGeneralizer_Generalize(t *testing.T) {

	g := &PrefixGeneralizer{MaxWords: 5}

	t.Run("level 0", func(t *testing.T) {
		actual := g.Generalize(partition.NewItem("this is a test string"), 0)
		expected := partition.NewItem("this is a test string")
		if !expected.Equals(actual) {
			t.Errorf("partitions are not equal: %v, %v", expected, actual)
		}
	})

	t.Run("generalize one level", func(t *testing.T) {
		actual := g.Generalize(partition.NewItem("this is a test string"), 1)
		expected := partition.NewItem("this is a test")
		if !expected.Equals(actual) {
			t.Errorf("partitions are not equal: %v, %v", expected, actual)
		}
	})

	t.Run("generalize multiple levels", func(t *testing.T) {
		actual := g.Generalize(partition.NewItem("this is a test string"), 3)
		expected := partition.NewItem("this is")
		if !expected.Equals(actual) {
			t.Errorf("partitions are not equal: %v, %v", expected, actual)
		}
	})

	t.Run("exceeds max words", func(t *testing.T) {
		actual := g.Generalize(partition.NewItem("this is a test string"), 20)
		testutil.AssertNil(actual, t)
	})

	t.Run("generalize last word", func(t *testing.T) {
		actual := g.Generalize(partition.NewItem("this is a test string"), 5)
		expected := partition.NewItem("*")
		if !expected.Equals(actual) {
			t.Errorf("partitions are not equal: %v, %v", expected, actual)
		}
	})

	t.Run("empty string input", func(t *testing.T) {
		actual := g.Generalize(partition.NewItem(""), 3)
		expected := partition.NewItem("*")
		if !expected.Equals(actual) {
			t.Errorf("partitions are not equal: %v, %v", expected, actual)
		}
	})

	t.Run("non string input", func(t *testing.T) {
		actual := g.Generalize(partition.NewItem(10), 3)
		expected := partition.NewItem("*")
		if !expected.Equals(actual) {
			t.Errorf("partitions are not equal: %v, %v", expected, actual)
		}
	})

	t.Run("non item partition", func(t *testing.T) {
		actual := g.Generalize(partition.NewSet(), 3)
		expected := partition.NewItem("")
		if !expected.Equals(actual) {
			t.Errorf("partitions are not equal: %v, %v", expected, actual)
		}
	})
}

func TestPrefixGeneralizer_Levels(t *testing.T) {
	g := &PrefixGeneralizer{MaxWords: 10}
	testutil.AssertEquals(11, g.Levels(), t)
}
