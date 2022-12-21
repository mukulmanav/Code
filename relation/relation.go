package relation

import (
	"fmt"

	"gorm.io/driver/mysql"

	_ "github.com/go-sql-driver/mysql"

	"gorm.io/gorm"
)

type Employ struct {
	gorm.Model
	Name         string
	Age          int
	emp_phone_no int
	emp_dept     string
}

type Dept struct {
	gorm.Model
	Dept_no   int
	Dept_name string
}

func Relation() {
	fmt.Print("hi")
	dsn := "root:1234@tcp(127.0.0.1:3306)/ve?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("no worries")
	}

	fmt.Println("successful", db)
	// SELECT * FROM `Employ` left join Dept on Employ.emp_dept = Dept.Dept_name
	//db.Model(&Employ{}).Select("name,Employ.Age,Employ.emp_phone_no,Employ.emp_dept  ").Joins("left join Dept on Employ.emp_dept = Dept.Dept_name").Scan(&Dept{})

	rows, err := db.Table("employs").Select("*").Rows()

	fmt.Print(rows)

	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("no worries")
	}
	//db.Table("Employ").Select("*").Joins("left join Dept on Employ.emp_dept = Dept.Dept_name").Scan(&Dept{})

	//db.Joins("JOIN emails ON emails.user_id = users.id AND emails.email = ?", "jinzhu@example.org").Joins("JOIN credit_cards ON credit_cards.user_id =users.id").Where("credit_cards.number = ?", "411111111111").Find(&user)
	//Marsh.Marsh()

}
