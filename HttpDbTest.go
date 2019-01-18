package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq" //database driver
)

type User struct {
	gorm.Model
	Name  string
	Email string
}

func allUsers(w http.ResponseWriter, r *http.Request) {
	db := openDatabase()
	defer db.Close()

	var users []User
	db.Find(&users)
	fmt.Println("{}", users)

	json.NewEncoder(w).Encode(users)
}

func findUser(w http.ResponseWriter, r *http.Request) {
	db := openDatabase()
	defer db.Close()

	vars := mux.Vars(r)
	id := vars["id"]
	fmt.Println("findUser() id =", id)

	u, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		fmt.Println("Given user id is not an integer.")
		panic(err.Error())
	}

	var user User
	user.Model.ID = uint(u)

	db.Find(&user)
	fmt.Println("{}", user)

	json.NewEncoder(w).Encode(user)
}

func newUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("New User Endpoint Hit")

	db := openDatabase()
	defer db.Close()

	vars := mux.Vars(r)
	fmt.Println(vars)

	name := vars["name"]
	email := vars["email"]

	fmt.Println(name)
	fmt.Println(email)

	db.Create(&User{Name: name, Email: email})
	fmt.Fprintf(w, "New User Successfully Created")
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	db := openDatabase()
	defer db.Close()

	vars := mux.Vars(r)
	name := vars["name"]
	fmt.Println("deleteUser() name =", name)

	var user User
	user.Name = name

	db.Where("name = ?", name).Find(&user)
	db.Delete(&user)

	fmt.Fprintf(w, "Successfully Deleted User")
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	db := openDatabase()
	defer db.Close()

	vars := mux.Vars(r)
	name := vars["name"]
	email := vars["email"]

	fmt.Println("updateUser() name =", name)

	var user User
	db.Where("name = ?", name).Find(&user)

	user.Email = email

	db.Save(&user)
	fmt.Fprintf(w, "Successfully Updated User")
}

// http url examples:
//
// 1) add a new user: 	http://localhost:8081/user/Tom/tom@gmail.com
// 2) find a user:		http://localhost:8081/user/1
func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/user/{name}/{email}", newUser).Methods("POST")
	myRouter.HandleFunc("/user/{name}/{email}", updateUser).Methods("PUT")
	myRouter.HandleFunc("/users", allUsers).Methods("GET")
	myRouter.HandleFunc("/user/{id}", findUser).Methods("GET")
	myRouter.HandleFunc("/user/{name}", deleteUser).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8081", myRouter))
}

func openDatabase() (db *gorm.DB) {
	db, err := gorm.Open("postgres", "user=phil password=postgres sslmode=disable")
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect to database")
	}
	return db
}

func initialMigration() {
	db := openDatabase()
	defer db.Close()

	// gorm AutoMigrate() will create the tables only if they don't already exist
	db.AutoMigrate(&User{})
}

func main() {
	fmt.Println("Go ORM Tutorial")

	initialMigration()
	// Handle Subsequent requests
	handleRequests()
}
