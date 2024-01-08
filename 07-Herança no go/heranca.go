package main

import "fmt"

type pessoa struct {
	nome string
	idade uint
	
}

type estudante struct {
	pessoa   // Essa é uma herança no Go!!!
	curso string
	periodo uint
}

func main() {
	fmt.Println("Herança. Só que não!!!")

	pessoa1 := pessoa{"Rogério", 40}
	fmt.Println(pessoa1)

	fmt.Println("======================================================")

	estudante := estudante{pessoa1, "Golang", 3}
	fmt.Println(estudante)
	fmt.Println(estudante.nome)
	fmt.Println(estudante.curso)

}