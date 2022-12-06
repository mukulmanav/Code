package vmap

import "fmt"

func Workwithmap() {
	statepopulation := make(map[string]int, 1)
	statepopulation = map[string]int{
		"california": 39256,
		"Texas":      2785,
		"Florida":    2061,
		"Ohio":       1161,
	}
	fmt.Println(statepopulation)
	fmt.Println("Ohio-", statepopulation["Ohio"])
	statepopulation["Georgia"] = 54753
	fmt.Println(statepopulation)
	delete(statepopulation, "Georgia")
	//_, ok : = statepopulation["Ohio"]
	fmt.Println(statepopulation)

}
