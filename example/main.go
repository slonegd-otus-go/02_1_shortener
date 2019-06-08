package main

import shortener "github.com/slonegd-otus-go/02_1_shortener"

func main() {
	chars := []rune{'a', 'b', 'c'}
	shortener := shortener.New(6, chars)

	long := "otus.ru/long_url"
	short := shortener.Shorten(long)
	println(short, shortener.Resolve(short))

	long = "otus.ua/long_url"
	short = shortener.Shorten(long)
	println(short, shortener.Resolve(short))
}
