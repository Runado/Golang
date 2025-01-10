package main

import ( // para importar é necessário abrir parenteses.
	"fmt"
	"modulo/NewPacote" // importando o pacote NewPacote criado para esta Aula
)

func main() {
	fmt.Println("Escrevendo Arquivo MAIN!")
	auxiliar.Escrever()  // Nesta linha estamos executando a função escrever a partir do arquivo auxiliar importado a partir do NewPacote
	auxiliar.escrever2() // Nesta linha estou executando a função escrever2 a partir do arquivo auxiliar importando a partir do NewPacote
	// porém a função escrever 2 não será importada pois a função foi definida dentro do escopo LOCAL, ou seja só irá funcionar caso esteja na mesma pasta do arquivo que está referenciando-a.

}
