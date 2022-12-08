package vint

import "fmt"

func main() {
	var i interface{} = "true"
	switch i.(type) {
	case int:
		fmt.Println("integer")
	case string:
		fmt.Println("string")
	default:
		fmt.Println("idk")
	}
}
