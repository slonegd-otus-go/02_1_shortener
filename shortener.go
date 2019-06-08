package shortener

import (
	"math"
	"math/rand"
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

func (s *shortener) Shorten(url string) string {
	if len(s.data) == s.maxLen {
		return ""
	}

	short := s.random()
	_, ok := s.data[short]

	for ok {
		s.increment(&short)
		_, ok = s.data[short]
	}

	s.data[short] = url

	return short
}

func (s shortener) Resolve(url string) string {
	return ""
}

func (s *shortener) random() string {
	str := ""
	for i := 0; i < s.runesQuantity; i++ {
		str += string(s.runes[rand.Intn(len(s.runes))])
	}
	return str
}

func (s *shortener) increment(str *string) {
	indexes := s.toIndexes(*str)
	for i := range indexes {
		indexes[i]++
		if indexes[i] < len(s.runes) {
			break
		}
		indexes[i] = 0
	}
	*str = s.fromIndexes(indexes)
}

func (s *shortener) toIndexes(str string) []int {
	res := make([]int, 0, len(str))
	for _, r := range str {
		res = append(res, s.findIndex(r))
	}
	return res
}

func (s *shortener) findIndex(r rune) int {
	for i, sr := range s.runes {
		if r == sr {
			return i
		}
	}
	return 0
}

func (s *shortener) fromIndexes(indexes []int) string {
	str := ""
	for _, i := range indexes {
		str += string(s.runes[i])
	}
	return str
}
