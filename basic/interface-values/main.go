package main

import "fmt"

type Describer interface {
	Describe() string
}

type Person struct {
	Name string
	Age  int
}

func (p Person) Describe() string {
	return fmt.Sprintf("Person(Name: %s, Age: %d)", p.Name, p.Age)
}

type Product struct {
	Name  string
	Price float64
}

func (p Product) Describe() string {
	return fmt.Sprintf("Product(Name: %s, Price: %.2f)", p.Name, p.Price)
}

func main() {
	person := Person{Name: "Alice", Age: 30}
	product := Product{Name: "Laptop", Price: 999.99}
	fmt.Println(person.Describe())
	fmt.Println(product.Describe())
}