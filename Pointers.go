package main

import "fmt"

type Person struct {
	name    string
	age     int
	isAdult bool //bool values default to false
}

// This function CAN change the given Person object,
// since it is passed a pointer to the original Person.
// Actually, it is passed a copy of the pointer (containing the same address).
func validateAge(p *Person) {
	fmt.Printf("validateAge(): p = %+v, \t%p\n", p, p)
	p.isAdult = (p.age >= 21)
}

// This function CANNOT change the given Person object,
// since it is passed a COPY of the original Person
func checkAge(p Person) {
	fmt.Printf("checkAge(): p = %+v, \t\t%p\n", p, &p)
	p.isAdult = (p.age >= 21)
}

func changeInt(i *int, newValue int) {
	//dereference the pointer to set the value to newValue
	*i = newValue //this reads as "the value at address i equals newValue"
}

func main() {

	p := Person{name: "Phil", age: 65}
	fmt.Printf("1) Person: %+v, \t\t\t%p\n", p, &p)

	//function passed a copy of p
	checkAge(p)
	fmt.Printf("2) Person: %+v, \t\t\t%p\n", p, &p)

	//function passed a copy of pointer (containing the address of p)
	validateAge(&p)
	fmt.Printf("3) Person: %+v, \t\t\t%p\n", p, &p)

	fmt.Println()
	i := 5
	fmt.Println("value before:", i)
	changeInt(&i, 10)
	fmt.Println("value after: ", i)

	var a int
	var b, c = &a, &a
	fmt.Println()
	fmt.Printf("b and c point to the same address: %p, %p\n", b, c)   // 0x1040a124 0x1040a124
	fmt.Printf("b and c are stored in 2 locations: %p, %p\n", &b, &c) // 0x1040c108 0x1040c110

}
