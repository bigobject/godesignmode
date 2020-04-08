package main

import "fmt"

type persuiter struct {
	name string
}

func (p persuiter) persuit(to target) {
	fmt.Println(p.name, "persuit", to.Name())
}

type proxy struct {
	ps persuiter
}

func (this proxy) persuit(to target) {
	this.ps.persuit(to)
}

type target struct {
	name string
}

func (t target) Name() string {
	return t.name
}

func main() {
	t := target{name: "lili"}

	p := proxy{ps: persuiter{name: "wanqiang"}}

	p.persuit(t)
}
