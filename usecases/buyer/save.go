package buyer

import (
	"errors"
	"pruebafinal/models"
)

func Save(buyer models.Buyer) (models.Buyer, error) {
	if buyer.ID != 0 {
		return models.Buyer{}, errors.New("New buyer must not include id")
	}
	buyer.ID = nextID
	nextID++
	buyers = append(buyers, &buyer)
	return buyer, nil
}
