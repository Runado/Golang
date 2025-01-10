package auxiliar // Criando um pacote novo

import "fmt" // Importando I/O

func escrever2() { // No Golang a primeira letra de uma função é extremamente importante pois se for maiúscula significa que essa função poderá ser importada por outros arquivos, caso a função iniciar com uma letra minúscula ela será apenas local no arquivo em que ela foi criada e não poderá ser importada por outros arquivos.
	//Função com letra maiúscula = Global \ Função com letra mínuscula = Local
	fmt.Println("Escrevendo um Arquivo apartir do pacote auxiliar")
	fmt.Println("Para usar outro pacote, deve ser utilizada a partir de outra pasta")
}
