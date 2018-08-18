package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	// números inteiros
	fmt.Println(1, 2, 1000)
	fmt.Println("Literal inteiro é", reflect.TypeOf(3200))

	//sem sinal somente valores positivos uint8 uint16 uint32 uint64
	var b byte = 255

	fmt.Println("O byte é", reflect.TypeOf(b))

	//com sinal ... int8 int16 int32 int64

	i1 := math.MaxInt64

	fmt.Println("O valor máximo do int é ", i1)
	fmt.Println("o tipo do i1 é ", reflect.TypeOf(i1))

	var i2 rune = 'a' //representa um mapeamento da tabela Unicode (int32)
	fmt.Println("O rune é ", reflect.TypeOf(i2))
	fmt.Println(i2) //valor do 'a' na tabela unicode

	//número reais float32 float64

	var x float32 = 49.99
	fmt.Println("o tipo do x é", reflect.TypeOf(x))
	fmt.Println("o tipo literal de 49.99 é", reflect.TypeOf(49.99)) //literal

	//boolean

	bo := true
	fmt.Println("O tipo de bo é", reflect.TypeOf(bo))
	fmt.Println(!bo)

	//String

	s1 := "Olá meu nome é GO!"

	fmt.Println(s1)
	fmt.Println("O tamanho da string é ", len(s1))

	//String com multiplas linhas

	s2 := `Olá 
	meu nome
	é
	GO!`

	fmt.Println(s2)

	fmt.Println("O tamanho da string em multiplas linhas é ", len(s2))

	// ??? char
	char := 'a'

	fmt.Println("o tipo char é ?? ñ há char ", reflect.TypeOf(char))

}
