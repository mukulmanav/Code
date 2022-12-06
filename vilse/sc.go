package vilse

import "fmt"

func Sc() {
	switch 2 {
	case 1, 5, 7:
		fmt.Println("one")
	case 2, 3, 4:
		fmt.Println("2")
	default:
		fmt.Println("none")
	}
}
