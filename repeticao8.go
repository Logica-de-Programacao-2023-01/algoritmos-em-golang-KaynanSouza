package main

import "fmt"

/*Faça um algoritmo que leia um número inteiro positivo e mostre todos os seus divisores.*/
func main() {
	var num int
	fmt.Print("escreva um numero:")
	fmt.Scan(&num)
	for i := 1; i < num; i++ {
		if num%i == 0 {
			print(" ", i)
		}
	}
}
