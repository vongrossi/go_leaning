package main

import (
	"fmt"
)

func tipo(i interface{}) string {
	switch i.(type) {

	case int:
		return "inteiro"
	case float32, float64:
		return "float"
	case string:
		return "string"
	case bool:
		return "boolean"
	case func():
		return "função"

	default:
		return "tipo não identificado"
	}

}

func main() {

	fmt.Println(tipo(3))
	fmt.Println(tipo(2 < 5))
	fmt.Println(tipo("azul"))
	fmt.Println(tipo(2.9))

}
