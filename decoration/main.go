package main

import "fmt"

type decoratedObject struct {
}

func (decoratedObject) operation() {
	fmt.Println("decoratedObject:operation")
}

type decorateOperation interface {
	operation()
}

type decorator1 struct {
	origin decorateOperation
}

func (this decorator1) operation() {
	fmt.Println("decorator1:operation")
	this.origin.operation()
}

type decorator2 struct {
	origin decorateOperation
}

func (this decorator2) operation() {
	fmt.Println("decorator2:operation")
	this.origin.operation()
}

func main() {
	obj := decoratedObject{}
	decorated1 := decorator1{origin: obj}
	decorated2 := decorator2{origin: decorated1}
	decorated3 := decorator2{origin: decorated2}

	decorated3.operation()
}
