package main

import "fmt"

func main() {
	// Existem duas formas de declarar variáveis:
	var variavel1 string = "variavel 1"; // dessa forma 
	variavel2 := "variavel 2"; // declaração implicita

	fmt.Println(variavel1)
	fmt.Println(variavel2)

	// posso declarar assim:
	var (
		nome string = "rogério"
		idade int = 40
	)
	fmt.Println(nome)
	fmt.Println(idade)

	// posso também declarar assim:
	nome1, idade1 := "juliana", 38
	fmt.Println(nome1)
	fmt.Println(idade1)

}