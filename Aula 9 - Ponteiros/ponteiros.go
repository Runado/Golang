package main

import "fmt"

func main() {
	fmt.Println("Ponteiros ")

	var variavel1 int = 10
	var variavel2 int = variavel1

	fmt.Println(variavel1, variavel2)

	variavel1++
	fmt.Println(variavel1, variavel2)
	// Esse exemplo para entende-lo ele deve ser executado, é possível verificar que o golang ao atribuir o valor de uma variavel como sendo outra variavel, ele faz apenas uma cópia do valor que está atribuida a aquela variável ou seja ao incrementar +1, ele irá adicionar uma unidade apenas a variavel que está incrementando.

	//Ponteiro é uma referência de memória
	var variavel3 int
	var ponteiro *int

	variavel3 = 100
	ponteiro = &variavel3
	fmt.Println(ponteiro) // Neste caso o ponteiro irá guardar não o valor mas a localização da variavel3 no endereço de memória assim sempre que for printado, será printado com o valor que está atribuido naquele endereço de memória.

}
