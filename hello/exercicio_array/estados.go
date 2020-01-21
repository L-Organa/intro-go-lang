package main

import "fmt"

func main(){
	estadosSudeste()
}

func estadosSudeste() [4]string  {
	var estado [4]string
	estado[0] = "SP"
	estado[1] = "RJ"
	estado[2] = "MG"
	estado[3] = "ES"
	fmt.Println(estado)
	return estado
}
	
