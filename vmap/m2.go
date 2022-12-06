package vmap

import "fmt"

func Seemap() {
	s := Createmap()
	fmt.Println(s)
}
func Createmap() (statepopulation map[string]int) {
	statepopulation = map[string]int{
		"california": 39256,
		"Texas":      2785,
		"Florida":    2061,
		"Ohio":       1161,
	}
	return statepopulation

}
