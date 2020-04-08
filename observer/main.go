package main

import (
	"fmt"
	"reflect"
)

type observerif interface {
	onNotify()
}

type subjectif interface {
	add(observer observerif)
	rm(observer observerif)
	notify()
}
type subject struct {
	observers []observerif
}

func (s *subject) add(o observerif) {
	s.observers = append(s.observers, o)
}

func (s *subject) rm(o observerif) {
	t := make([]observerif, 0)
	for _, v := range s.observers {
		if !reflect.DeepEqual(v, o) {
			t = append(t, v)
		}
	}
	s.observers = t
}
func (s subject) notify() {
	fmt.Println("len observers is:", len(s.observers))
	for _, v := range s.observers {
		v.onNotify()
	}
}

type observer struct {
	name string
}

func (o observer) onNotify() {
	fmt.Println(o.name, "got notify")
}

func main() {
	s := subject{}
	s.add(observer{name: "cc"})
	s.add(observer{name: "gl"})

	s.notify()
	s.rm(observer{name: "gl"})
	s.notify()
}
