package main

import "fmt"

// Função é um typo, então posso atribuir a uma variável:
func soma(n1, n2 int) int {
	return n1 + n2
}

func calculos(n1, n2 int64) (int64, int64) {
	soma := n1 + n2
	subitracao := n1 - n2
	return soma, subitracao 
}

func main() {
	// Podem ter Parâmetro e Retorno:
	fmt.Println(soma(10, 20))

	somar_dois_numeros := soma(15, 30)
	fmt.Println(somar_dois_numeros)

	fmt.Println(calculos(10, 5))
	soma, subitração := calculos(50, 30)
	// Funçôes em go, podem ter mais de um retorno:
	fmt.Println("Soma:", soma)
	fmt.Println("Subitração:", subitração)

	// forma de usar só um retorno:
	soma2, _ := calculos(50, 25)
	fmt.Println(soma2)
	

}