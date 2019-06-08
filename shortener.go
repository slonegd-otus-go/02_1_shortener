package shortener

type shortener struct {
	data          map[string]string
	charsQuantity int
	chars         []rune
}

// New ctor
func New(charsQuantity int, chars []rune) *shortener {
	return &shortener{
		make(map[string]string),
		charsQuantity,
		chars,
	}
}

func (s *shortener) Shorten(url string) string {
	return ""
}

func (s shortener) Resolve(url string) string {
	return ""
}
