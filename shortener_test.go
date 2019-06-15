package shortener_test

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"

	shortener "github.com/slonegd-otus-go/02_1_shortener"
)

func Test_shortener_Shortener(t *testing.T) {
	chars := []rune{'0', '1'}
	shortener := shortener.New(2, chars)

	tests := []struct {
		name        string
		url         string
		wantLenUrl  int
		mustContain string
	}{
		{"http://otus.ru/1", "http://otus.ru/1", 2, "http://otus.ru/"},
		{"without scheme", "otus.ru/1", -1, ""},
		{"bad URL", "://otus.ru/2", -1, ""},
		{"https://otus.ru/big_string", "https://otus.ru/big_string", 2, "https://otus.ru/"},
		{"http://google.com/another_string", "http://google.com/another_string", 2, "http://google.com/"},
		{"https://otus.ru/last_able_string", "https://otus.ru/last_able_string", 2, "https://otus.ru/"},
		{"vocabluary full", "http://otus.ru/something", -1, ""},
	}

	urls := []string{}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := shortener.Shorten(tt.url)
			url, _ := url.Parse(got)
			assert.Equal(t, tt.wantLenUrl, len(url.Path)-1)
			assert.Contains(t, got, tt.mustContain)

			assert.NotContains(t, urls, got)
			if len(got) > 0 {
				urls = append(urls, got)
			}
		})
	}
}

func Test_shortener_Resolve(t *testing.T) {
	chars := []rune{'0', '1'}
	shortener := shortener.New(2, chars)

	long := []string{"http://otus.ru/test", "http://otus.ru/another", "https://otus.ru/teacher-lk/homework/49392/3743/", "http://otus.ru/Harry_Potter"}
	domain := []string{"http://otus.ru/", "http://otus.ru/", "https://otus.ru/", "http://otus.ru/"}
	short := make([]string, 0, len(long))
	for _, url := range long {
		short = append(short, shortener.Shorten(url))
	}

	for i, long := range long {
		t.Run(long, func(t *testing.T) {
			got := shortener.Resolve(short[i])
			assert.Equal(t, long, got)
			assert.Contains(t, got, domain[i])
		})
	}

	assert.Equal(t, "", shortener.Resolve("not exist"))
}
