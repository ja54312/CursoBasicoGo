package main

import "fmt"

func main() {
	//Declaracion de variables
	helloMessage:= "hello"
	worldMessage:="world"

	//Print ln Imprime y agrega un salto de linea
	fmt.Println(helloMessage, worldMessage)
	fmt.Println(helloMessage, worldMessage)

	//Printf Imprime y agrega una funcion extra
	nombre := "Platzi"
	cursos := 500
	fmt.Printf("%s tiene mas de %d cursos\n",nombre,cursos)
	fmt.Printf("%v tiene mas de %v cursos\n",nombre,cursos)

	// Sprintf 
	message := fmt.Sprintf("%v tiene mas de %v cursos",nombre,cursos)
	fmt.Println(message)

	// Tipo datos
	fmt.Printf("helloMessage: %T\n",helloMessage)
	fmt.Printf("cursos: %T",cursos)
}
