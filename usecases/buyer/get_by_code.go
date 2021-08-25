package buyer

import (
	"fmt"
	"pruebafinal/models"
)

func GetByCode(code int) (models.Buyer, error) {
	for _, b := range buyers {
		if b.Code == code {
			return *b, nil
		}
	}
	return models.Buyer{}, fmt.Errorf("Buyer with code '%v' not found", code)
}
