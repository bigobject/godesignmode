package main

import "fmt"

type cloneable interface {
	Clone() cloneable
}

type foo struct {
	i int
}

func (f foo) Clone() cloneable {
	c := f
	return c
}

func main() {

	a := foo{i: 2}
	fmt.Println(a)

	b := a.Clone()
	fmt.Println(b)
}
