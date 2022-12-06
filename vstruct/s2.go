package vstruct

import "fmt"

func Main3() {
	adoc := struct{ name string }{name: "tom"}
	fmt.Printf(adoc.name)
	bdoc := adoc
	bdoc.name = "jom"
	fmt.Println(bdoc)
}
