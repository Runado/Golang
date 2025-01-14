package main

import "fmt"

func main() {
	fmt.Println("Arrays e Slices")
	fmt.Println("Em Golang há diversas maneiras de criar listas")

	var array1 [5]string
	array1[0] = "Posição 1"
	fmt.Println(array1)

	array2 := [4]string{"Posição 1", "Posição 2", "Posição 3", "Posição 4"}
	fmt.Println(array2)

	array3 := [...]int{1, 2, 3, 4, 5} // este tipo de array vai fixar o numero de itens com o número de valores que foi passado, neste caso foi passado 5 itens dessa forma o array vai ter 5 posições.
	fmt.Println(array3)

	// SLICE é o conceito de lista mais utilizado no golang pois ele possuí tamanho dinâmico e tem limitação apenas de tipo de dado.
	slice := []int{1, 2, 3, 4, 5, 6, 7, 10, 12, 13}
	fmt.Println(slice)
	// o Slice ele possuí um ponteiro e seu funcionamento é semelhante ao de um Array.

	//Type Of serve para verificar o tipo de dados.
	slice = append(slice, 10) // essa é uma forma de adicionar dados em um Slice, o que essa linha faz é basicamente criar um novo slice adicionando os dados que foi inserido.
	// no caso acima o append será no Slice de nome "Slice" e será adicionado o valor 10, será criado um novo slice com o novo dado inserido na ultima posição.
	fmt.Println(slice)

	slice2 := array2[1:3] // é importante saber que ao definir um range o primeiro item neste caso o "1" é inclusivo, já segundo indice neste caso o "3" ele é exclusivo.
	fmt.Println(slice2)

	array2[1] = "Posição Alterada"
	fmt.Println(slice2)

}
