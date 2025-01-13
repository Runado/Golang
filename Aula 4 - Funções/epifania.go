package main

import "fmt"

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

func calculosMatematicos(n1, n2 int8) (int8, int8) { // esta função esta recebendo dois parametros N1 e N2 e retornando dois valores do tipo int8
	soma := n1 + n2
	subtracao := n1 - n2
	return soma, subtracao
} // caso não declarar o tipo da variavel, ela irá pegar o mesmo tipo da ultima que foi definida neste caso o N1 sera int8 tbm

func main() {
	soma := somar(10, 20)
	fmt.Println(soma)
	var f = func(txt string) string {
		fmt.Println(txt)
		return txt
	}
	resultado := f("Texto da função 1")
	fmt.Println(resultado)

	//resultadoSoma, _ := calculosMatematicos(10, 15) # neste exemplo estou passando os dois valores pra função porém estou guardando apenas 1 retorno da mesma, apenas o resultado da soma será retornado assim não irá haver um conflito na hora de compilar um código com uma variável que não está sendo utilizada.
	// útil quando tiver retornos naquela função que não serão utilizados em uma primeira instância.
	resultadoSoma, resultadoSubtracao := calculosMatematicos(50, 15)
	// as duas variáveis criadas estarão passando para a função os valores 10 e 15 e irão receber o retorno de forma sequencial.
	fmt.Println(resultadoSoma, resultadoSubtracao)
	// nesta linha estou printando o resultado de ambas as operações.

}

// !!! AS FUNÇÕES NO GOLANG PODEM TER MAIS DE UM RETORNO !!!
