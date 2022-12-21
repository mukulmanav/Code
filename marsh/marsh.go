package Marsh

import (
	"encoding/json"
	"fmt"
	"log"
)

func Marsh() {
	var x = 5
	bytes, err := json.Marshal(x)
	if err != nil {
		fmt.Println("Can't serislize", x)
	}
	fmt.Printf("%v => %v, '%v'\n", x, bytes, string(bytes))
	// Deserialize int
	var r int
	err = json.Unmarshal(bytes, &r)
	if err != nil {
		fmt.Println("Can't deserislize", bytes)
	}
	fmt.Printf("%v => %v\n", bytes, r)
}

type Person struct {
	Name string
}

func Marsh2() {

	P := Person{Name: "mukul"}
	pbytes, err := json.Marshal(P)
	log.Print(err)
	log.Print(string(pbytes))
}
