package shortener

import (
	"math"
	"math/rand"
	"net/url"
	"strings"
)

type shortener struct {
	data          map[string]string
	runesQuantity int
	runes         []rune
	maxLen        int
}

// New ctor
func New(runesQuantity int, runes []rune) *shortener {
	return &shortener{
		data:          make(map[string]string),
		runesQuantity: runesQuantity,
		runes:         runes,
		maxLen:        int(math.Pow(float64(len(runes)), float64(runesQuantity))),
	}
}

func (s *shortener) Shorten(in string) string {
	if len(s.data) == s.maxLen {
		return ""
	}

	url, err := url.Parse(in)
	if err != nil {
		println(err.Error())
		return ""
	}
	if len(url.Host) == 0 {
		println("Must be host name")
		return ""
	}

	var builder strings.Builder
	if len(url.Scheme) != 0 {
		builder.WriteString(url.Scheme)
		builder.WriteString("://")
	}
	builder.WriteString(url.Host)
	builder.WriteRune('/')
	domain := builder.String()

	short := s.random()
	_, ok := s.data[domain+short]

	for ok {
		short = s.increment(short)
		_, ok = s.data[domain+short]
	}

	s.data[domain+short] = in

	return domain + short
}

func (s shortener) Resolve(url string) string {
	return s.data[url]
}

func (s *shortener) random() string {
	var builder strings.Builder
	for i := 0; i < s.runesQuantity; i++ {
		builder.WriteRune(s.runes[rand.Intn(len(s.runes))])
	}
	return builder.String()
}

func (s *shortener) increment(in string) string {
	indexes := s.toIndexes(in)
	for i := range indexes {
		indexes[i]++
		if indexes[i] < len(s.runes) {
			break
		}
		indexes[i] = 0
	}
	return s.fromIndexes(indexes)
}

func (s *shortener) toIndexes(str string) []int {
	res := make([]int, 0, len(str))
	for _, r := range str {
		res = append(res, s.findIndex(r))
	}
	return res
}

func (s *shortener) findIndex(r rune) int {
	i := 0
	for i = range s.runes {
		if r == s.runes[i] {
			break
		}
	}
	return i
}

func (s *shortener) fromIndexes(indexes []int) string {
	var builder strings.Builder
	for _, i := range indexes {
		builder.WriteRune(s.runes[i])
	}
	return builder.String()
}
