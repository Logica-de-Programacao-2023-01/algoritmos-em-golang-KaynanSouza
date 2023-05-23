package main

import (
	"fmt"
	"sort"
)

//Faça um algoritmo que leia três números reais e mostre-os em ordem crescente.

func main() {
	var x int
	slice := []int{}
	for i := 0; i < 3; i++ {
		fmt.Print("Digite o primeiro valor:")
		fmt.Scan(&x)
		slice = append(slice, x)
	}
	sort.Ints(slice)
	fmt.Print(slice)
}
