package main

import ("fmt"
		"reflect"
)

type human interface {
	getNome() string
	getIdade() int
	getProfissao() string
	
}

type pessoa struct {
	nome string
	idade int
	profissao string
}

func (p pessoa) getNome() string {
	return p.nome
}

func (p pessoa) getIdade() int {
	return p.idade
}

func (p pessoa) getProfissao() string {
	return p.profissao
}

func retorno(h human) {
	fmt.Printf("Nome: %s\n", h.getNome())
	fmt.Printf("Idade: %d anos\n",h.getIdade())
	fmt.Printf("Profissão: %s\n", h.getProfissao())
} 

func main () {


p := pessoa{nome: "Alisson", idade: 27, profissao: "Desenvolvedor"}
c := pessoa{nome: "Carla", idade: 25, profissao: "Arquiteta"}


//fmt.Printf("Olá, me chamo %s, tenho %d anos e trabalho como %s", p.getNome(), p.getIdade(), p.getProfissao())
retorno(p)
retorno(c)


//Utilizar uma função que vá fazer os métodos com seus retornos é muito mais eficaz do que printar cada objeto 
//individualmente conforme mostra abaixo.
//Da maneira correta, desenvolvo uma funcao que execute os 3 métodos com os 3 dados, me trazendo as informaçoes a partir 
//da execução dos métodos na interface. Sendo assim, eu posso criar vários objetos e usar a mesma funcao pra chamar todos eles.
//Ou seja, se eu fizesse p e c e quisesse printar cada um, seja chamando o método individualmente ou só pelo atributo,
//Precisaria fazer 1 print pra cada atributo conforme mostra abaixo, e para cada objeto. 2 objetos resultariam em 6 linhas de prints.
//Da outra maneira, eu uso a funcao com 3 prints como resultado, e preciso apenas passar o objeto pra ser executado pelos métodos
//da interface. E mesmo que eu passe 10 objetos, usarei as mesmas 3 linhas da funcao retorno, diferente do outro jeito, que seriam 30
//Linhas de prints

fmt.Printf("Nome: %s\n", p.nome)
fmt.Printf("Idade: %d\n", p.idade)
fmt.Printf("Profissão: %s\n", p.profissao)

}