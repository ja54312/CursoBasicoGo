package main

import (
	pk "CursoBasicoGo/src/mypackage.go"
	"fmt"
)

func main() {
	var myCar pk.CarPublic
	myCar.Brand = "Ferrari"
	myCar.Year = 2020
	fmt.Println(myCar)
	
	pk.PrintMessage("Hola Platzi")

}