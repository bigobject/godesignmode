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

type creactor1 struct {
}

func (creactor1) createinstance(want string) IF {
	switch want {
	case "instance1":
		return instance1{}
	case "instance2":
		return instance2{}
	default:
		return nil
	}
}

type creactor2 struct {
}

func (creactor2) createinstance(want string) IF {
	switch want {
	case "instance1":
		return instance2{}
	case "instance2":
		return instance1{}
	default:
		return nil
	}
}

type CIF interface {
	createinstance(want string) IF
}

func main() {
	input := []string{"instance1", "instance2", "instance3"}
	var creatotr CIF = creactor1{}
	fmt.Println("creatinstance with creactor1")
	for _, v := range input {
		if instance := creatotr.createinstance(v); instance != nil {
			fmt.Println(instance.Function(""))
		} else {
			fmt.Println("creatinstance failed, name:", v)
		}
	}

	creatotr = creactor2{}
	fmt.Println("creatinstance with creactor2")
	for _, v := range input {
		if instance := creatotr.createinstance(v); instance != nil {
			fmt.Println(instance.Function(""))
		} else {
			fmt.Println("creatinstance failed, name:", v)
		}
	}

}
