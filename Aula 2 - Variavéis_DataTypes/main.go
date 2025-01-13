package main

import "fmt"

// O Go ele não permite que uma variável criada fique sem utilidade dentro do código, nestes casos o código nem mesmo compila.
func main() {
	var variavel1 string = "Variável 1"
	variavel2 := "Variável 2" // Inferencia de Dados
	fmt.Println(variavel1)
	fmt.Println(variavel2)

	var ( // Multiplas variaveis definidas
		variavel3 string = "hohoho"
		variavel4 string = "lalala"
	)
	fmt.Println(variavel3, variavel4)

	variavel5, variavel6 := "Variável 5", "Variável 6"
	fmt.Println(variavel5, variavel6)

	const constante1 string = "constante 1"
	fmt.Println(constante1)

	variavel5, variavel6 = variavel6, variavel5
	fmt.Println(variavel5, variavel6)

}
