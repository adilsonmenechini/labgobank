package genrand

import (
	"strconv"
)

type generateCard struct {
	CardNumber      string
	CardType        cardType
	CNumber         int
	CCV             int
	ExpirationMonth int
	ExpirationYear  int
}

func GenerateCard(ctype string, cYear, cNum int) *generateCard {
	expirationMonth, expirationYear := generateExpirationDate(cYear)
	return &generateCard{
		CardNumber:      generateCreditCardNumber(ctype),
		CCV:             generateCCV(cNum),
		ExpirationMonth: expirationMonth,
		ExpirationYear:  expirationYear,
	}
}

func generateCreditCardNumber(cardType string) string {
	var prefixes []string
	var length int

	switch cardType {
	case "visa":
		prefixes = []string{"4"}
		length = 16
	case "mastercard":
		prefixes = []string{"51", "52", "53", "54", "55"}
		length = 16
	case "amex":
		prefixes = []string{"34", "37"}
		length = 15
	case "discover":
		prefixes = []string{"6011", "644", "645", "646", "647", "648", "649", "65"}
		length = 16
	default:
		return "Invalid card type"
	}

	prefix := prefixes[generateRandomNumber(0, len(prefixes))]
	number := prefix
	for i := len(prefix); i < length-1; i++ {
		number += strconv.Itoa(generateRandomNumber(0, 9))
	}

	sum := 0
	double := false
	for i := len(number) - 1; i >= 0; i-- {
		digit, _ := strconv.Atoi(string(number[i]))
		if double {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}
		sum += digit
		double = !double
	}
	checkDigit := (10 - (sum % 10)) % 10
	number += strconv.Itoa(checkDigit)

	return number
}
