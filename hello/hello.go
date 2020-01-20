package main

import "fmt"
import "os"

// This is how you comment on Go 

func main(){

	exibeIntroducao()
	exibeMenu()
	comando:= leComado()

// Using Switch instead of If 
	switch comando {
	case 1:
		fmt.Println("1. Iniciar monitoramento")
	case 2:
		fmt.Println("2. Exibir Logs")
	case 0:
		fmt.Println("0. Sair do Programa")
		os.Exit(0)
	default:
		fmt.Println("Nao conheço esse comando")
		os.Exit(-1)
	}
}


func exibeIntroducao() {
	nome:= "Douglas"
	versao:= 1.1
	fmt.Println("Ola Sr.", nome)
	fmt.Println("Esse programa esta na versao", versao)
}

func exibeMenu() {
	fmt.Println("1. Iniciar monitoramento")
	fmt.Println("2. Exibir Logs")
	fmt.Println("0. Sair do Programa")
}

func leComado() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println("O comando escolhido foi", comandoLido)
	return (comandoLido)
}