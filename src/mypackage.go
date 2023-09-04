package mypackage

import "fmt"

// CarPublic Car con acceso publico
type CarPublic strict {
	Brand string
	Year int
}

type carPrivate struct {
	brand string
	year int
}

func PrintMessage(text string) {
	fmt.Println(text)

}