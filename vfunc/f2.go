package vfunc

import "fmt"

func F2() {
	d, err := divide(5.0, 3.0)
	if err != nil {
		fmt.Println(d)
	}
}

func divide(a, b float64) (float64, error) {
	if b == 0.0 {
		return 0.0, fmt.Errorf("cannot divide by zero")
	}
	return a / b, nil
}
