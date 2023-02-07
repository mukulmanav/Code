package main

import (
	"fmt"
	"vexample/Try"
	"vexample/lasagna"
	"vexample/vetest"
	"vexample/vfunc"
	"vexample/vilse"
	"vexample/vint"
	"vexample/vmap"
	"vexample/vstruct"

	//_ "github.com/go-sql-driver/mysql"

	"vexample/extra"
	Marsh "vexample/marsh"
	Meth "vexample/meth"
	"vexample/mserv"
	"vexample/relation"

	"gorm.io/gorm"
)

type Employ struct {
	gorm.Model
	Name         string
	Age          int
	emp_phone_no int
	emp_dept     string
}

type Dept struct {
	gorm.Model
	Dept_no   int
	Dept_name string
}

type Person struct {
	gorm.Model
	Name string
	Age  int
}

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

	//Gorm code
	fmt.Println("hi")
	dsn := "root:1234@tcp(127.0.0.1:3306)/ve?charset=utf8mb4&parseTime=True&loc=Local"

	//var p Person
	//p.Insert(dsn)
	//p.Update(dsn)
	//p.Del(dsn)

	//fmt.Println("Method")
	fmt.Println(string(dsn))
	Meth.Meth()
	fmt.Println("Relation")
	relation.Relation()
	fmt.Println("Marshal")
	Marsh.Marsh2()
	Try.Try()
	lasagna.Lasagna()
	fmt.Println("MUX")
	mserv.Initialmigration()
	mserv.Initializerouter()
	fmt.Println("IAM")
	extra.Extra()
}
