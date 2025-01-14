package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    uint8
}
type estudante struct {
	pessoa    // Isso é o mais próximo do conceito de herança no Golang, é possível importar as variaveis de outra struct dentro de uma struct, basta colocar ela dentro da struct e não atribuir nenhum tipo a ela.
	curso     string
	faculdade string
}

func main() {
	fmt.Println("Herança")
	p1 := pessoa{"Jose", "Lucas", 25, 180} // nessa linha estou passando os valores as variáveis da struct pessoa
	fmt.Println(p1)
	e1 := estudante{p1, "Segurança da Informação", "FATEC Ourinhos"} // por meio do conceito de "Herança" do Golang eu consigo printar o P1 por meio da struct Estudante
	fmt.Println(e1.altura)
}
