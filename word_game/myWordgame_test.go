package main

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestWordGame_TryConnect(t *testing.T) {
	tests := []struct {
		name string
		word string
	}{
		{"friendship", "friendship"},
		{"ababab", "ababab"},
		{"testwordgame", "testwordgame"},
		{"test blank word", "test blank word"},
		{"pairs and nicole", "pairs and nicole"},
		{"hello world", "hello world"},
		{"ababa abb aab", "ababa abb aab"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wg := NewWordGame(tt.word)
			cnt := 0
			for len(wg.WordParts) > 1 && cnt <= 1000 {
				cnt++
				currentIndex, touchedIndex := rand.Intn(len(wg.WordParts)), rand.Intn(len(wg.WordParts))
				fmt.Println("currentIndex:", currentIndex, "touchedIndex:", touchedIndex)
				wg.TryConnect(currentIndex, 1, touchedIndex, -1)
			}
			require.Equal(t, wg.Word, wg.WordParts[0])
		})
	}
}

func BenchmarkWordGame_TryConnect(b *testing.B) {
	for i := 0; i < b.N; i++ {
		wg := NewWordGame(randomString(30))
		cnt := 0
		for len(wg.WordParts) > 1 && cnt <= 1000 {
			cnt++
			currentIndex, touchedIndex := rand.Intn(len(wg.WordParts)), rand.Intn(len(wg.WordParts))
			wg.TryConnect(currentIndex, 1, touchedIndex, -1)
		}
	}
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func randomString(n int) string {
	b := make([]byte, n)
	for i := range b {
		if i > 0 && i%5 == 0 {
			b[i] = ' '
			continue
		}
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
