package vint

import "fmt"

type Incrementer interface {
	Increment() int
}

func Vin2() {
	myInt := Intcounter(0)
	var inc Incrementer = &myInt
	for i := 0; i < 10; i++ {
		fmt.Println(inc.Increment())
	}
}

type Intcounter int

func (ic *Intcounter) Increment() int {
	*ic++
	return int(*ic)
}
