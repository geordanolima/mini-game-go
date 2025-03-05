package helpers

import (
	"mini-game-go/domain"

	"golang.org/x/exp/rand"
)

func GetProportionalPercent() int {
	pesos := make([]int, len(domain.PercentsGas))
	somaPesos := 0
	for i, percent := range domain.PercentsGas {
		pesos[i] = 100 - percent + 1
		somaPesos += pesos[i]
	}
	numeroAleatorio := rand.Intn(somaPesos)
	somaParcial := 0
	for i, peso := range pesos {
		somaParcial += peso
		if numeroAleatorio < somaParcial {
			return domain.PercentsGas[i]
		}
	}
	return domain.PercentsGas[len(domain.PercentsGas)-1]
}
