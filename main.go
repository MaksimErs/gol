package main

import (
	"fmt"

	"github.com/MaksimErs/gol/gret"
)

func main() {
	var name string
	fmt.Println("Введите имя:")
	fmt.Scan(&name)
	fmt.Println("Hello", name)
	gret.Hell()
}
