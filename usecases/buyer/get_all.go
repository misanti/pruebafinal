package buyer

import "pruebafinal/models"

var (
	buyers []*models.Buyer
	nextID = 1
)

func GetAll() []*models.Buyer {
	return buyers
}
