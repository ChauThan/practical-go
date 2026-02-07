package main

import "fmt"

type Describer interface {
	Describe() string
}

type Person struct {
	Name string
	Age  int
}

func (p *Person) Describe() string {

	if(p == nil){
		return "Person is nil"
	}
	return fmt.Sprintf("Person(Name: %s, Age: %d)", p.Name, p.Age)
}

func main() {
	var d Describer

	// Assigning a nil *Person to the Describer interface
	var p *Person = nil
	d = p

	fmt.Println(d.Describe())
}

