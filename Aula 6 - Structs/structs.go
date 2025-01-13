package main

import "fmt"

type usuario struct {
	nome     string
	idade    int8
	endereco endereco // Nessa linha estou colocando uma struct dentro de outra struct.

}
type endereco struct {
	logradouro string
	numero     uint8
}

func main() {
	fmt.Println("Arquivo Structs")
	// PRIMEIRA FORMA PARA DEFINIR OS VALORES DAS VARIAVEIS NA STRUCT
	var u usuario
	u.nome = "Jose"
	u.idade = 24
	fmt.Println(u)
	enderecoExemplo := endereco{"Rua dos bobos", 0} // definindo valores para a Struct Endereco
	// SEGUNDA FORMA PARA DEFINIR OS VALORES DAS VARIAVEIS NA STRUCT
	usuario2 := usuario{"Jose", 25, enderecoExemplo}
	fmt.Println(usuario2)
	// TERCEIRA FORMA PARA DEFINIR OS VALORES DAS VARIAVEIS NA STRUCT
	usuario3 := usuario{idade: 25}
	fmt.Println(usuario3)

	usuario4 := usuario{nome: "Jose"}
	fmt.Println(usuario4)
}
