package main

import (
	"fmt"
	hello "unit-testing/hello"
)

func main() {

	p := hello.NewPerson("Javier") // Instancia de la estructura Persona

	s, err := p.Salute()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(s)
}
