package genrand

import (
	"math/rand"
	"time"
)

func generateRandomNumber(min, max int) int {
	return rand.Intn(max-min) + min
}
func generateExpirationDate(y int) (int, int) {
	// Gera um mês e um ano aleatórios para a data de expiração do cartão
	currentYear := time.Now().Year()
	year := currentYear + y
	month := generateRandomNumber(1, 12)
	return month, year
}

func generateCCV(c int) int {
	ccv := 0
	for i := 0; i < c; i++ {
		ccv += generateRandomNumber(0, 9)
	}
	return ccv
}
