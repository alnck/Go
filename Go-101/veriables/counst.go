package main

const i = "5"

const (
	a = iota
	b
	c
)

func main() {
	println(i, a, b, c)
}
