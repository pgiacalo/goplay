package main


// Uses GORM Object Relational Mapper

// Youtube lesson on GORM
// https://www.youtube.com/watch?v=nVD9acHituc

import (
	"fmt"
	"github.com/jinzhu/gorm"	//object relational mapper
	_ "github.com/lib/pq"		//database driver
	"time"
)

func main(){
	fmt.Println("DBTest.main() called")
	//postgres: local user=phil, docker user=postgres
	db, err := gorm.Open("postgres", "user=phil password=postgres sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.DB().Ping()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Ping complete, err=%v\n", err)

	db.SingularTable(true)
	db.DropTableIfExists(&Owner{}, &Book{}, &Author{})
	db.CreateTable(&Owner{}, &Book{}, &Author{})

	//owner := Owner {
	//	FirstName: "Joe",
	//	LastName: "Blow",
	//}
	//db.Create(&owner)
	//
	//owner.LastName = "Smith"
	//db.Debug().Save(&owner)
	//
	//db.Debug().Delete(&owner)

}

type Owner struct {
	gorm.Model
	FirstName 	string
	LastName 	string
	Books 		[]Book
}

type Book struct {
	gorm.Model
	Title 			string
	PublishDate 	time.Time
	OwnerID 		uint `sql:"index"`
	Authors			[]Author `gorm:"many2many:books_authors"`
}

type Author struct {
	gorm.Model
	FirstName string
	LastName string
}
