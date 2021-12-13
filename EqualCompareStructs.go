package main

import (
  "fmt"
  "reflect"
)

type Developer struct {
  Name string
  Age  int
}

func main() {
  fmt.Println("Comparing 2 Structs in Go")

  elliot := Developer{
    Name: "Elliot",
    Age:  26,
  }

  elliot2 := Developer{
    Name: "Elliot",
    Age:  26,
  }

  fmt.Println(reflect.DeepEqual(elliot, elliot2))

}
