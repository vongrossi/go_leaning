package main

import (
	"fmt"
)

func notaParaConceito(n float64) string {

	var nota = int(n)

	//Por padrão o switch em go executa uma condição e caso a encontre já sai fora do switch
	//para que o switch continue a executar o proximo caso é necessario a palavra falltrough
	//em go não é necessario o uso de break para sair de um case no switch

	switch nota {
	case 10:
		fallthrough
	case 9:
		return "A"
	case 8, 7:
		return "B"
	case 6, 5:
		return "C"
	case 4, 3:
		return "D"
	case 2, 1:
		return "E"
	default:
		return "Valor invalido"
	}

}

func main() {

	fmt.Println(notaParaConceito(9.8))
	fmt.Println(notaParaConceito(6.4))
	fmt.Println(notaParaConceito(11))

}
