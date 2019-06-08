package shortener

type shorten struct {
	data          map[string]string
	charsQuantity int
	chars         []rune
}

// New ctor
func New(charsQuantity int, chars []rune) *shorten {
	return &shorten{
		make(map[string]string),
		charsQuantity,
		chars,
	}
}

func (s *shorten) Shorten(url string) string {
	return ""
}

func (s shorten) Resolve(url string) string {
	return ""
}
