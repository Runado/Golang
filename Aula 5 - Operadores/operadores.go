package main

import "fmt"

func main() {
	// ARITMÉTICOS
	soma := 1 + 2
	subtracao := 1 - 2
	divisao := 10 / 4
	multiplicacao := 10 * 10
	restodadivisao := 10 % 2
	fmt.Println(soma, subtracao, divisao, multiplicacao, restodadivisao)

	//*var numero1 int16 = 10
	//*var numero2 int32 = 25
	//*soma := numero1 + numero2
	//*fmt.Println(soma)
	//// esse script comentado acima não irá funcionar, pois para realizar qualquer operação aritmética ou não os tipos de dados devem ser iguais, somar um int16 com um int32 não é possível no Golang.

	//// !!!GOLANG É FORTEMENTE TIPADO!!!

	// ATRIBUIÇÃO
	var variavel string = "string"
	variavel2 := "string"
	fmt.Println(variavel, variavel2)
	// FIM DOS OPERADORES DE ATRIBUIÇÃO

	// OPERADORES RELACIONAIS
	fmt.Println(1 > 2)  //Maior
	fmt.Println(1 >= 2) // Maior ou Igual
	fmt.Println(1 == 2) // Igual
	fmt.Println(1 <= 2) // Menor ou Igual
	fmt.Println(1 < 2)  // Menor
	fmt.Println(1 != 2) // Diferente
	//Fim dos Operadores Relacionais

	// OPERADORES LÓGICOS
	fmt.Println("----------------------")
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso) // Operador AND
	fmt.Println(verdadeiro || falso) // Operador OU
	fmt.Println(!verdadeiro)         // Negação/Inverter o valor
	fmt.Println(!falso)              // Negação/Inverter o valor
	// OPERADORES UNÁRIOS
	numero := 10
	numero++
	numero += 15 // numero = numero + 15
	fmt.Println(numero)

	numero--
	numero -= 20 // numero = numero - 20

	numero *= 3  // numero = numero * 3
	numero /= 10 // numero = numero / 10
	numero %= 3  // numero = numero % 3

	fmt.Println(numero)
	// FIM DOS OEPRADORES UNÁRIOS
	var texto string
	if numero > 5 {
		texto = "Maior que 5"
	} else {
		texto = "Menor que 5"
	}
	fmt.Println(texto)
}
