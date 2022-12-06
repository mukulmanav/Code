package vstruct

import "fmt"

type Animal struct {
	name   string
	origin string
}

type Bird struct {
	Animal
	Speed  float32
	Canfly bool
}

func Main4() {
	fmt.Println("")
	b := Bird{}
	b.name = "Emu"
	b.origin = "Aus"
	b.Canfly = false
	fmt.Println(b)

}
