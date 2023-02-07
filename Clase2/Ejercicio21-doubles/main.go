package main

import (
	"fmt"
	"testdoubles/simulator"
)

func main() {
	fmt.Println("Hello World")
	sm := simulator.NewCatchSimulator(10)
	fmt.Println(sm.GetLinearDistance([2]float64{1, 3})) // 3,16

}
