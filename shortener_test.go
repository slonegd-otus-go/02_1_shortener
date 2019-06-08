package shortener_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	shortener "github.com/slonegd-otus-go/02_1_shortener"
)

func Test_shortener_Shorten(t *testing.T) {
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

			urls := append(urls, got)
			assert.NotContains(t, urls, got)
		})
	}
}
