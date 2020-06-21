package main

import "fmt"

func main() {
	g := newEven()
	fmt.Println(g())
	f := func() int {
		return square(g())
	}
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	//gen := mapInt(newEven(), square)
	//fmt.Println(gen())
	//fmt.Println(gen())
	//fmt.Println(gen())
	//gen = nil // release for garbage collection
}

type intGen func() int

func newEven() intGen {
	n := 0
	return func() int {
		n += 2
		return n
	}
}

func mapInt(g intGen, f func(int) int) intGen {
	return func() int {
		return f(g())
	}
}

func square(i int) int {
	return i * i
}
