//Performing Crud cmmands using mux

package mserv

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

const DNS = "root:1234@tcp(127.0.0.1:3306)/ve?charset=utf8mb4&parseTime=True&loc=Local"

type User struct {
	gorm.Model
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func Initializerouter() {

	r := mux.NewRouter()

	r.HandleFunc("/users", Getusers).Methods("GET")
	r.HandleFunc("/users/{id}", Getuser).Methods("GET")
	r.HandleFunc("/users", Createuser).Methods("POST")
	r.HandleFunc("/users/{id}", Updateuser).Methods("PUT")
	r.HandleFunc("/users/{id}", Deleteuser).Methods("DELETE")
	http.Handle("/", r)

	log.Fatal(http.ListenAndServe(":9000", r))
	fmt.Println(DNS)
}

func Initialmigration() {
	DB, err = gorm.Open(mysql.Open(DNS), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("no worries")
	}
	DB.AutoMigrate(&User{})
}
func Getusers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/json")
	var user []User
	DB.Find(&user)
	json.NewEncoder(w).Encode(user)
}
func Getuser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/json")
	params := mux.Vars(r)
	var user User
	DB.First(&user, params["id"])
	json.NewEncoder(w).Encode(user)
}

func Createuser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/json")
	var user User
	json.NewDecoder(r.Body).Decode(&user)
	DB.Create(&user)
	json.NewEncoder(w).Encode(user)
}

func Updateuser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/json")
	params := mux.Vars(r)
	var user User
	DB.First(&user, params["id"])
	json.NewDecoder(r.Body).Decode(&user)
	DB.Save(&user)
	json.NewEncoder(w).Encode(user)
}

func Deleteuser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/json")
	params := mux.Vars(r)
	var user User
	DB.Delete(&user, params["id"])
	json.NewEncoder(w).Encode("Deleted successfully")
}
