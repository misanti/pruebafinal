package buyer

import (
	"fmt"
	"pruebafinal/models"
)

func Get(id int) (models.Buyer, error) {
	for _, b := range buyers {
		if b.ID == id {
			return *b, nil
		}
	}
	return models.Buyer{}, fmt.Errorf("Buyer with ID '%v' not found", id)
}
