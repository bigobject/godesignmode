package main

import "fmt"

type IF interface {
	Function(input string) (output string)
}

type instance1 struct {
}

func (instance1) Function(input string) (output string) {
	return "instance1"
}

type instance2 struct {
}

func (instance2) Function(input string) (output string) {
	return "instance2"
}

type IFContext struct {
	myif IF
}

func newFIContext(input string) *IFContext {
	switch input {
	case "instance1":
		return &IFContext{myif: instance1{}}
	case "instance2":
		return &IFContext{myif: instance2{}}
	default:
		return nil
	}
}

func (c IFContext) result() string {
	return c.myif.Function("")
}

func main() {
	input := []string{"instance1", "instance2", "instance3"}
	for _, v := range input {
		if ctx := newFIContext(v); ctx != nil {
			fmt.Println(ctx.result())
		} else {
			fmt.Println("unknow v:", v)
		}
	}

}
