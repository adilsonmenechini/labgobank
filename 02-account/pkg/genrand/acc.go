package genrand

import "strconv"

type cardType string

type generateAcoount struct {
	CardNumber  string
	AccountType cardType
}

// generateCard gera um cartão de crédito.
func GenerateAcoount(ctype string) *generateAcoount {
	return &generateAcoount{
		CardNumber:  generateAccountNumber(ctype),
		AccountType: cardType(ctype),
	}
}

func generateAccountNumber(cardType string) string {
	var prefixes []string
	var length int

	switch cardType {
	case "bb":
		prefixes = []string{"123"}
		length = 7
	case "itau":
		prefixes = []string{"32"}
		length = 6
	case "caixa":
		prefixes = []string{"21"}
		length = 6
	case "santander":
		prefixes = []string{"53"}
		length = 9
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
