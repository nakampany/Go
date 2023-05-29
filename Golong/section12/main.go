package main

import "fmt"

type Strings interface {
	ToString() string
}

type Person struct {
	Name string
	Age  int
}

type Car struct {
	Number string
	Model  string
}

func (p *Person) ToString() string {
	return fmt.Sprintf(p.Name, p.Age)
}
func (p *Car) ToString() string {
	return fmt.Sprintf(p.Number, p.Model)
}

func main() {
	v := []Strings{&Person{Name: "a", Age: 1}, &Car{Number: "1", Model: "1"}}

	for _, s := range v {
		fmt.Println(s)
	}
	// &{a 1}
	// &{1 1}
}
