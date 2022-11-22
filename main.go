//OBJETIVO: Programa que faz a conversão de escalas termométricas.
//DESAFIO: Um professor de ensino medio solicitou aos seus alunos que desenvolvessem um programa que converte o valor de ponto de ebulição da água de Kelvin para graus Celsius

package main

import "fmt"

const ebulicionK float64 = 373.00

func main() {
	ebulicionC := int(ebulicionK - 273)
	fmt.Printf("A temperatura de ebulição em Kelvin é: %gº %T  e o ponto de ebulição em Celsius é %dº %T", ebulicionK, ebulicionK, ebulicionC, ebulicionC)
}
