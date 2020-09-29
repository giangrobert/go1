package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {

	s := Person{Name: "Sean", Age: 50}
	//s := Person{}
	//s.Name = "Sean"
	//s.Age = 42
	fmt.Println(s.Name)

	sp := &s
	fmt.Println(sp.Age)

	sp.Age = 51
	fmt.Println(sp.Age)
}
