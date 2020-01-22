package main

import "fmt"
import "os"
import "net/http"
import "time"
import "bufio"
import "io"
import "strings"

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
	nome:= "Carol Rolo"
	versao:= 1.1
	fmt.Println("Ola Sra.", nome)
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
	sites:= leSiteDoArquivo()

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
	resp, err := http.Get(site)

	if err != nil{
		fmt.Println("Ocorreu um erro")
	}

	fmt.Println(site)

	if resp.StatusCode == 200{
	fmt.Println("Site", site, "foi carregado com sucesso!")
	} else {
		fmt.Println("Site", site, "esta com problemas. Status Code:", resp.StatusCode)
	}
}

func leSiteDoArquivo() []string {
	var sites []string
    arquivo, err := os.Open("sites.txt")

    if err != nil {
        fmt.Println("Ocorreu um erro:", err)
    }

    leitor := bufio.NewReader(arquivo)
    for {
        linha, err := leitor.ReadString('\n')
        linha = strings.TrimSpace(linha)

        sites = append(sites, linha)

        if err == io.EOF {
            break
        }

    }

    arquivo.Close()
    return sites
}





