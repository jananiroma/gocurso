package main

import (
	"fmt"
	"github.com/gbrayhan/gocurso/maps"
	"github.com/gbrayhan/gocurso/numbers"
	"github.com/gbrayhan/gocurso/structs"
)


func main() {

	fmt.Println( maps.GetMap("Alejandra",21))
	structs.InterfaceTest()

	//structs.InterfaceTest()
	number, err := numbers.Sum("50", 50)
	if err != nil {
		//panic(err)
		fmt.Println(err)
	}
	fmt.Println(number)
	
}