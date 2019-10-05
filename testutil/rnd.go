package testutil

import (
	"math/rand"
	"strings"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandText(wordCount, wordLen int) string {
	sb := &strings.Builder{}
	for i := 0; i < wordCount; i++ {
		sb.WriteString(RandString(wordCount))
	}
	return sb.String()
}

func RandString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

