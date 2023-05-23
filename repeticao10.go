package main

import (
	"fmt"
)

// Faça um algoritmo que leia vários números inteiros e mostre o maior deles.
// A leitura deve ser interrompida quando for lido o valor zero.
func main() {
	var num int = 1
	maior := 1
	menor := 100
	for num != 0 {
		fmt.Print("Digite um valor:")
		fmt.Scan(&num)
		if num != 0 {
			if num > maior {
				maior = num
			} else {
				menor = num
			}
		}
	}
	fmt.Println("O maior numero sera:", maior)
	fmt.Print("O menor numero sera:", menor)
}
