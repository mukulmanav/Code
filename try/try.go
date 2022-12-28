package Try

import (
	"fmt"
)

func Try() {
	s := "hello hello"
	for k, v := range s {
		fmt.Println(k, string(v))
	}
}
