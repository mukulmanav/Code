//program to do insert . update and delete with methods

package Meth

import (
	"fmt"

	"gorm.io/driver/mysql"

	_ "github.com/go-sql-driver/mysql"

	"gorm.io/gorm"
)

type Person struct {
	gorm.Model
	Name string
	Age  int
}

func Meth() {
	fmt.Println("hi")
	dsn := "root:1234@tcp(127.0.0.1:3306)/ve?charset=utf8mb4&parseTime=True&loc=Local"

	var p Person
	p.Insert(dsn)
	p.Update(dsn)
	p.Del(dsn)

}

// insert entries
func (p Person) Insert(dsn string) {

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("no worries")
	}

	fmt.Println("successful")

	fmt.Print(db.Select(&Person{Name: "Mukul", Age: 21}))

}

// update
func (p Person) Update(dsn string) {

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("no worries")
	}

	p = Person{Name: "Mukul", Age: 21}
	db.Model(&p).Where(&Person{Age: 21}).Update("Age", 22)
	fmt.Printf("%v(%v)", p.Name, p.Age)
}

// delete(soft)
func (p Person) Del(dsn string) {

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("no worries")
	}
	//delete(soft)
	p = Person{}
	db.Where(&Person{Age: 22}).First(&p)
	fmt.Println("to be deleted", p.Name)
	db.Delete(&p)

	fmt.Printf("%v(%v)", p.Name, p.Age)
}
