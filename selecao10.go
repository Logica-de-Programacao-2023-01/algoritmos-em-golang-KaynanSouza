package main

import (
	"fmt"
)

func main() {
	var idade int
	fmt.Print("Digite sua idade:")
	fmt.Scan(&idade)
	if idade < 1 {
		fmt.Println("Idade incorreta, digite novamente")
	} else if idade >= 1 && idade < 9 {
		fmt.Println("mirim")
	} else if idade >= 10 && idade <= 13 {
		fmt.Println("infantil")
	} else if idade >= 14 && idade <= 17 {
		fmt.Println("juvenil")
	} else {
		fmt.Println("adulto")
	}
}
