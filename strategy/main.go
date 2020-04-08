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

func (c IFContext) result() string {
	return c.myif.Function("")
}

func main() {
	input := []string{"instance1", "instance2", "instance3"}
	for _, v := range input {
		switch v {
		case "instance1":
			fmt.Println(IFContext{myif: instance1{}}.result())
		case "instance2":
			fmt.Println(IFContext{myif: instance2{}}.result())
		case "instance3":
			fmt.Println("unknow v:", v)
		}
	}

}
