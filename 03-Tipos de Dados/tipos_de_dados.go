package main

import (
	"errors"
	"fmt"
)

func main() {

// Tipos numéricos:  := -> quando por inferência é só int!!!
// int8, int16, int32, int64
// float32, float64
	var idade int16 = 40
	salario := 3000
	altura := 1.77
	fmt.Println(idade, altura, salario)

// Tipos numéricos sem sinal: => uint8, 16, 32 e 64
	var peso uint16 = 75
	fmt.Println(peso)

// elias de int32 == rune ; uint8 == byte
	var letra rune = 'A'
	fmt.Println(letra)

// Tipo booleano => bool => true or false
	var estaChovendo bool = true
	fmt.Println(estaChovendo)

// Tipo erro => error!!!
	var erro error = errors.New("sobre o erro")
	fmt.Println(erro)
}
