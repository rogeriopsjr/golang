package main

import "fmt"

type usuario struct {
	nome string
	idade uint
	altura float32
}

type usuario2 struct {
	nome string
	idade uint
	endereco endereco
}

type endereco struct {
	rua string
	numero uint
}

func main() {

	var user1 usuario
	user1 = usuario{"Rogério", 40, 1.77}
	
	fmt.Println(user1)

	fmt.Println(user1.nome)
	fmt.Println(user1.idade)

	fmt.Println("================================================")

	var user2 usuario
	user2.nome = "Juliana"
	user2.idade = 38
	
	fmt.Println(user2.nome)
	fmt.Println(user2.idade)

	fmt.Println("================================================")

	user3 := usuario{"Gabryella", 14, 1.60}
	fmt.Println(user3.nome)
	fmt.Println(user3.idade)

	fmt.Println("================================================")

// Seu eu não tiver todos os dados, posso fazer assim:
	user4 := usuario{nome:"peludinho"}
	fmt.Println(user4)
	fmt.Println(user4.nome)

	fmt.Println("================================================")
// Posso ter struct dentro de struct
	enderUser1 := endereco{"Francisco de Assis", 50}
	fmt.Println(enderUser1)



}