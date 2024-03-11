package main

import "fmt"

func main() {
	var name string
	fmt.Println("Введите имя:")
	fmt.Scan(&name)
	fmt.Println("Hello", name)
}
