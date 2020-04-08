package main

import "fmt"

type Logger struct {
	name string
}

// An Option configures a Logger.
type Option interface {
	apply(*Logger)
}

// optionFunc wraps a func so it satisfies the Option interface.
type optionFunc func(*Logger)

func (f optionFunc) apply(log *Logger) {
	f(log)
}

func adda(l *Logger) {
	l.name += "a"
}

func addb(l *Logger) {
	l.name += "b"
}

func constructl(funcs ...optionFunc) *Logger {
	l := Logger{}
	for _, v := range funcs {
		v.apply(&l)
	}
	return &l
}

func main() {
	l := Logger{}

	fmt.Println(l)

	adda(&l)
	fmt.Println(l)

	ll := constructl(adda, adda, addb)
	fmt.Println(*ll)
}
