package main

import "fmt"

func main() {
	// ARITIMÉTICOS:
	// soma => +; subtração => -; divisão => /; multiplicação => *
	// resto => %
	// Só posso realizar cálculos com valores do mesmo tipo:
	var n1 int16 = 15
	var n2 int16 = 25
	fmt.Println( n1 + n2)

	// ATRIBUIÇÃO: =; :=
	var filha string = "Gaby"
	idade := 14
	fmt.Println(filha, idade)

	fmt.Println("===========================================================")


	// RELACIONAIS: retornam true or false
	fmt.Println(10 > 5) // maior que
	fmt.Println(15 >= 20) // maior ou igual a
	fmt.Println(15 == 15) // igual a
	fmt.Println(25 < 50) // menor que
	fmt.Println(32 <= 17) // menor ou igual a
	fmt.Println(10 != 15) //diferente de

	fmt.Println("===========================================================")

	// OPERADORES LÓGICOS: && -> E ; || -> OU ; ! -> NEGAÇÃO
	fmt.Println(10 > 5 && 15 < 25) // as duas comparações precisam ser true
	fmt.Println(10 < 50 || 5 == 5) // uma precisa ser true
	fmt.Println(!(10 > 5)) // inverte o resultado da expressão

	fmt.Println("===========================================================")

	// OPERADORES UNÁRIOS:
	//  ++; --; +=; -=

	// No goLang -> não existe o operador ternário!!!



}
