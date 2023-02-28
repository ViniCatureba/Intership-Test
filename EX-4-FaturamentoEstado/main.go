package main

import "fmt"

type FaturamentoPorEstado struct {
	Estado string
	Valor  float64
}

func main() {
	faturamentoDistribuidoraPorEstado := []FaturamentoPorEstado{
		{
			Estado: "SP",
			Valor:  67836.43,
		},

		{
			Estado: "RJ",
			Valor:  36678.66,
		},

		{
			Estado: "MG",
			Valor:  29229.88,
		},
		{
			Estado: "ES",
			Valor:  27165.48,
		},
		{
			Estado: "Outros",
			Valor:  19849.53,
		},
	}

	var somaValor float64
	for _, EstadoEValor := range faturamentoDistribuidoraPorEstado {
		somaValor += EstadoEValor.Valor
	}

	var totalporc float64
	for _, EstadoEValor := range faturamentoDistribuidoraPorEstado {
		porcentualDeRepresentacao := (EstadoEValor.Valor / somaValor) * 100
		totalporc += porcentualDeRepresentacao
		fmt.Printf("O estado de %v, representa um percentual de %v%%\n", EstadoEValor.Estado, porcentualDeRepresentacao)
	}

}
