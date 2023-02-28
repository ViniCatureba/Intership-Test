package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Faturamento struct {
	Dia   int     `json:"dia"`
	Valor float64 `json:"valor"`
}

func main() {
	dados, err := ioutil.ReadFile("dados.json")
	if err != nil {
		panic(err)
	}

	var faturamentos []Faturamento

	err = json.Unmarshal(dados, &faturamentos)
	if err != nil {
		panic(err)
	}

	var sliceDeDias []int
	var menorDia, maiorDia int
	var somaFaturamento, menorFaturamento, maiorFaturamento float64

	for _, faturamentosSingulares := range faturamentos {

		if faturamentosSingulares.Valor != 0.00 {
			if menorFaturamento == 0 {
				menorFaturamento = faturamentosSingulares.Valor
			}
			if faturamentosSingulares.Valor < menorFaturamento {
				menorDia = faturamentosSingulares.Dia
				menorFaturamento = faturamentosSingulares.Valor
			}
		}
		if faturamentosSingulares.Valor > maiorFaturamento {
			maiorDia = faturamentosSingulares.Dia
			maiorFaturamento = faturamentosSingulares.Valor
		}
		somaFaturamento += faturamentosSingulares.Valor
		sliceDeDias = append(sliceDeDias, faturamentosSingulares.Dia)
	}
	fmt.Printf("O MENOR faturamento foi no dia %v com um valor total de R$ %v\n\n\n", menorDia, menorFaturamento)

	fmt.Printf("O MAIOR faturamento foi no dia %v com um valor total de R$ %v\n\n\n", maiorDia, maiorFaturamento)

	lenSliceDeDias := float64(len(sliceDeDias))

	mediaMensalFaturamento := somaFaturamento / lenSliceDeDias

	var diasAcimaDaMediaMensal []int

	for _, faturamentosSingulares := range faturamentos {
		if faturamentosSingulares.Valor > mediaMensalFaturamento {
			diasAcimaDaMediaMensal = append(diasAcimaDaMediaMensal, faturamentosSingulares.Dia)
		}
	}
	fmt.Printf("Segue os dias que ultrapassaram o faturamento mensal de R$ %v: %v", mediaMensalFaturamento, diasAcimaDaMediaMensal)

}
