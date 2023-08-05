package main

import "fmt"

func main() {
	//Declaracion de constantes
	// const pi float64 = 3.14
	// const pi2 = 3.1415
	// fmt.Println("pi:",pi)
	// fmt.Println("pi2:",pi2)

	//Declaracion de variables enteras
	// base := 12
	// var altura int = 14
	// var area int

	// fmt.Println(base,altura,area)

	// Zero Values
	// var a int
	// var b float64
	// var c string
	// var d bool

	// fmt.Println(a,b,c,d)

	//Area Cuadrado
	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado

	fmt.Println("El área del cuadrado es:", areaCuadrado)

	x := 10
	y := 50
	//Suma
	result := x + y
	fmt.Println("Suma:", result)

	//Resta
	result = y - x
	fmt.Println("Resta:", result)

	//Multiplicacion
	result = x * y
	fmt.Println("multiplicacion:", result)

	//division
	result = y / x
	fmt.Println("division:", result)

	//Modulo o residuo
	result = y % x
	fmt.Println("residuo", result)

	//Incrementar
	x++
	fmt.Println("Incrementar:", x)

	//Decrementar
	x--
	fmt.Println("Decrementar:", x)

	// Rectángulo
	baseRectangulo := 20
	alturaRectangulo := 10

	areaRectangulo := baseRectangulo * alturaRectangulo

	fmt.Println("El Area del Rectángulo es :", areaRectangulo)

	// Circulo : AreaCirculo = pi por radio al cudrado
	const PI float64 = 3.14 // Constant
	var radioCirculo float64 = 10

	areaCirculo := PI * radioCirculo * radioCirculo

	fmt.Println("El Area del Circulo es :", areaCirculo)

	// Trapecio
	var baseUno float64 = 6
	var baseDos float64 = 15
	var alturaTrapecio float64 = 25

	areaTrapecio := ((baseUno + baseDos) * alturaTrapecio) / 2

	fmt.Println("El Area del Trapecio es :", areaTrapecio)
}
