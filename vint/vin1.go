package vint

import "fmt"

func vint() {
	var w Writer = Consolewriter{}
	w.Write([]byte("Hello"))
}

type Writer interface {
	Write([]byte ) (int,error)

}

type Consolewriter struct{}

func(cw Consolewriter) Write(data []byte (int,error)){
	n,err :=fmt.Println(string(data))
	return n,err
}