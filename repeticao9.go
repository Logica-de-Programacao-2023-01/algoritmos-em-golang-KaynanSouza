package main

import "fmt"

//Faça um algoritmo que leia vários números inteiros e mostre a média aritmética entre eles.
//A leitura deve ser interrompida quando for lido o valor zero.

func main() {
	var num int = 1
	var soma, media float64
	cont := 0
	i := 1
	for num != 0 {
		fmt.Print("Digite um numero:")
		fmt.Scan(&num)
		soma += float64(num * i)
		i++
		cont++
	}
	media = soma / float64(cont)
	fmt.Print("A media aritmetica é:", media)
}
