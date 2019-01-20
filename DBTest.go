package main

// Uses GORM Object Relational Mapper

// Youtube lesson on GORM
//https://www.youtube.com/watch?v=VAGodyl84OY
// or
// https://www.youtube.com/watch?v=nVD9acHituc

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm" //object relational mapper
	_ "github.com/lib/pq"    //database driver
)

func init() {
	fmt.Println("DBTest init() called")
}

func main() {
	fmt.Println("DBTest.main() called")
	//postgres: local user=phil, docker user=postgres
	db, err := gorm.Open("postgres", "user=phil password=postgres sslmode=disable")
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect to database.")
	}
	defer db.Close()

	err = db.DB().Ping()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Ping complete, err=%v\n", err)

	//AutoMigrate will create the database tables (only if they do NOT already exist)
	db.AutoMigrate(&Owner{}, &Book{}, &Author{})

	//db.SingularTable(true)
	//db.DropTableIfExists(&Owner{}, &Book{}, &Author{})
	//db.CreateTable(&Owner{}, &Book{}, &Author{})

	owner := Owner{
		FirstName: "Joe",
		LastName:  "Blow",
	}
	db.Create(&owner)

	owner.LastName = "Smith"
	db.Debug().Save(&owner)
	//
	//db.Debug().Delete(&owner)

}

type Owner struct {
	gorm.Model
	FirstName string
	LastName  string
	Books     []Book
}

type Book struct {
	gorm.Model
	Title       string
	PublishDate time.Time
	OwnerID     uint     `sql:"index"`
	Authors     []Author `gorm:"many2many:books_authors"`
}

type Author struct {
	gorm.Model
	FirstName string
	LastName  string
}
