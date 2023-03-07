package main

import (
	"modulo/auxiliar"
	"fmt"
	"github.com/badoux/checkmail"
)

func main()  {
	fmt.Println("Escrevendo do arquivo main")
	auxiliar.Escrever()
	erro := checkmail.ValidateFormat("devbook%gmail.com")
	fmt.Println(erro)
	
}