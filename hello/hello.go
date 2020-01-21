package main

import "fmt"
import "os"
import "net/http"
import "time"


const monitoramento = 3	 
const delay = 5

func main(){

	exibeIntroducao()

	for{
		exibeMenu()
		comando:= leComado()
		switch comando {
		case 1:
			iniciarMonitoramento()
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
	fmt.Println("")
	return (comandoLido)
}

func iniciarMonitoramento()  {
	fmt.Println("Monitorando")
	sites:= []string {"http://random-status-code.herokuapp.com", "http://alura.com.br", "http://caelum.com.br"}

	for i:= 0; i < monitoramento; i++{
		for i, site:= range sites{
				fmt.Println("Testando site: ", i, ":", site)
				testaSite(site)
		}
		time.Sleep(delay * time.Second)
		fmt.Println("")
	}

	fmt.Println("")	
}


func testaSite(site string) {
	resp, _ := http.Get(site)
	fmt.Println(site)
	if resp.StatusCode == 200{
	fmt.Println("Site", site, "foi carregado com sucesso!")
	} else {
		fmt.Println("Site", site, "esta com problemas. Status Code:", resp.StatusCode)
	}
}


