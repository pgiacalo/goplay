package main

import (
	"fmt"
	"reflect"
)

func main() {
	type Mother struct {
		Ssn      int      `json:"required maxSsn=9"` //tag
		Name     string   `json:"required maxName=100"`
		Children []string `json:"optional"`
	}

	t := reflect.TypeOf(Mother{})
	field, _ := t.FieldByName("Name")
	fmt.Println(field.Tag) // json:"required maxLength=100"

}
