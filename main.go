package main

/*
import (
	"fmt"
	"vexample/vetest"
	"vexample/vfunc"
	"vexample/vilse"
	"vexample/vint"
	"vexample/vmap"
	"vexample/vstruct"
)
*/
import (
	"log"
	"vexample/vcrud"

	"github.com/gin-gonic/gin"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading env file")
	}
}

func main() {
	/*vetest.Veteest()
	vilse.Vilse()
	vilse.Sc()
	vfunc.F1()
	vfunc.F2()
	vmap.Workwithmap()
	fmt.Println("createmap")
	vmap.Seemap()
	fmt.Println("structure functions")
	vstruct.Main2()
	vstruct.Main3()
	vstruct.Main4()
	vint.Vin2()
	*/

	//GORM COMMANDS

	vcrud.Init()
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080
}
