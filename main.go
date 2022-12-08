package main

import (
	"fmt"
	"vexample/vetest"
	"vexample/vfunc"
	"vexample/vilse"
	"vexample/vint"
	"vexample/vmap"
	"vexample/vstruct"
)

func main() {
	vetest.Veteest()
	vilse.Vilse()
	vilse.Sc()
	vfunc.F1()
	vfunc.F2()
	vmap.Workwithmap()
	fmt.Println("createmap")
	vmap.Seemap()
	fmt.Println("structure functions")
	vstruct.Main2()
	vstruct.Main3()
	vstruct.Main4()
	vint.Vin2()

}
