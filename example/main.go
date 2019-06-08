package main

import shortener "github.com/slonegd-otus-go/02_1_shortener"

func main() {
	chars := []rune{'0', '1'}
	shortener := shortener.New(2, chars)
	println(shortener)
}
