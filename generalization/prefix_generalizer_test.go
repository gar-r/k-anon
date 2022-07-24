package generalization

import (
	"fmt"
	"testing"

	"git.okki.hu/garric/k-anon/partition"
	"git.okki.hu/garric/k-anon/testutil"
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

	t.Run("level exceeds max words", func(t *testing.T) {
		actual := g.Generalize(partition.NewItem("this is a test string"), 20)
		testutil.AssertNil(actual, t)
	})

	t.Run("input word count exceeds max words", func(t *testing.T) {

		s := "this is a test string which is longer than max words"

		t.Run("level 0", func(t *testing.T) {
			actual := g.Generalize(partition.NewItem(s), 0)
			expected := partition.NewItem("this is a test string")
			if !expected.Equals(actual) {
				t.Errorf("expected %v, got %v", expected, actual)
			}
		})

		t.Run("level 1", func(t *testing.T) {
			actual := g.Generalize(partition.NewItem(s), 1)
			expected := partition.NewItem("this is a test")
			if !expected.Equals(actual) {
				t.Errorf("expected %v, got %v", expected, actual)
			}
		})

		t.Run("level 2", func(t *testing.T) {
			actual := g.Generalize(partition.NewItem(s), 2)
			expected := partition.NewItem("this is a")
			if !expected.Equals(actual) {
				t.Errorf("expected %v, got %v", expected, actual)
			}
		})

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

func BenchmarkPrefixGeneralizerWordCount(b *testing.B) {
	const maxWords = 5000
	const wordLen = 5
	for i := 0; i <= maxWords; i += 100 {
		b.Run(fmt.Sprintf("%d/%d/%d", maxWords, i, wordLen), func(b *testing.B) {
			text := testutil.RandText(i, wordLen)
			benchmarkPrefixRangeGeneralizer(text, maxWords, b)
		})
	}
}

func BenchmarkPrefixGeneralizerWordLength(b *testing.B) {
	const maxWords = 50
	const maxWordLen = 250
	for i := 90; i <= maxWordLen; i += 10 {
		b.Run(fmt.Sprintf("%d", i), func(b *testing.B) {
			text := testutil.RandText(maxWords, i)
			benchmarkPrefixRangeGeneralizer(text, maxWords, b)
		})
	}
}

func BenchmarkPrefixGeneralizerMaxWords(b *testing.B) {
	const words = 10
	const wordLen = 10
	text := testutil.RandText(words, wordLen)
	for i := 10; i <= 500; i += 10 {
		b.Run(fmt.Sprintf("%d", i), func(b *testing.B) {
			benchmarkPrefixRangeGeneralizer(text, i, b)
		})
	}
}

func benchmarkPrefixRangeGeneralizer(text string, maxWords int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		g := &PrefixGeneralizer{maxWords}
		p := g.Generalize(partition.NewItem(text), g.Levels()/2)
		if p == nil {
			b.Error("partition was nil")
		}
	}
}
