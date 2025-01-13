package main

import ( // para importar é necessário abrir parenteses.
	"fmt"
	auxiliar "modulo/NewPacote" // importando o pacote NewPacote criado para esta Aula

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Escrevendo Arquivo MAIN!")
	auxiliar.Escrever() // Nesta linha estamos executando a função escrever a partir do arquivo auxiliar importado a partir do NewPacote
	//auxiliar.escrever2() // Nesta linha estou executando a função escrever2 a partir do arquivo auxiliar importando a partir do NewPacote
	// porém a função escrever 2 não será importada pois a função foi definida dentro do escopo LOCAL, ou seja só irá funcionar caso esteja na mesma pasta do arquivo que está referenciando-a.
	erro := checkmail.ValidateFormat("devbook@gmail.com")
	fmt.Println(erro)
}

//Para buscar pacotes Externos é utilizado o comando GO GET [URL], assim o pacote será adicionado no arquivo go.mod e em seguida poderá ser utilizado dentro do arquivo main.go
