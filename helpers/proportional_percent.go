package helpers

import (
	"golang.org/x/exp/rand"
)

func GetProportionalPercent(percentsGas []int) int {
	pesos := make([]int, len(percentsGas))
	somaPesos := 0
	for i, percent := range percentsGas {
		pesos[i] = 100 - percent + 1
		somaPesos += pesos[i]
	}
	numeroAleatorio := rand.Intn(somaPesos)
	somaParcial := 0
	for i, peso := range pesos {
		somaParcial += peso
		if numeroAleatorio < somaParcial {
			return percentsGas[i]
		}
	}
	return percentsGas[len(percentsGas)-1]
}
