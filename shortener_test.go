package shortener_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	shortener "github.com/slonegd-otus-go/02_1_shortener"
)

func Test_shortener_Shortener(t *testing.T) {
	chars := []rune{'0', '1'}
	shortener := shortener.New(2, chars)

	tests := []struct {
		name       string
		url        string
		wantLenUrl int
	}{
		{"1", "1", 2},
		{"big string", "big string", 2},
		{"another string", "another string", 2},
		{"last able string", "last able string", 2},
		{"vocabluary full", "something", 0},
	}

	urls := []string{}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := shortener.Shorten(tt.url)
			assert.Equal(t, tt.wantLenUrl, len(got))

			assert.NotContains(t, urls, got)
			urls = append(urls, got)
		})
	}
}

func Test_shortener_Resolve(t *testing.T) {
	chars := []rune{'0', '1'}
	shortener := shortener.New(2, chars)

	long := []string{"test", "another", "something", "Harry Potter"}
	short := make([]string, 0, len(long))
	for _, url := range long {
		short = append(short, shortener.Shorten(url))
	}

	for i, long := range long {
		t.Run(long, func(t *testing.T) {
			got := shortener.Resolve(short[i])
			assert.Equal(t, long, got)
		})
	}
}
