package utils

import "github.com/google/uuid"

func GenerateUUID() string {
	return uuid.New().String()
}

func ValidateUUID(id string) bool {
	_, err := uuid.Parse(id)
	return err == nil
}
