package main

import (
	"fmt"
	"math"
)

func areaCirculo(radio float64) float64{ 
	return math.Pi*radio*radio
}
func areaRectangulo(base,altura float64) float64 {
	return base*altura
}

func areaTrapezoide (B,b,h float64) float64{
	return h*(B+b)/2
}

func main() {
	fmt.Printf("Circulo %.2f \n",areaCirculo(2))
	fmt.Printf("Rectangulo %.2f \n",areaRectangulo(5,10))
	fmt.Printf("Trapezoide %.2f \n",areaTrapezoide(10,5,3))
}