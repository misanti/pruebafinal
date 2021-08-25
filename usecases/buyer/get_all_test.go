package buyer

import (
	"github.com/stretchr/testify/assert"
	"pruebafinal/models"
	"testing"
)

func TestGetAll(t *testing.T) {
	buyer1 := models.Buyer{1, 1234, "Bruno", "Santi"}
	buyer2 := models.Buyer{2, 5678, "Micaela", "Santi"}

	buyers = append(buyers, &buyer1, &buyer2)
	got := GetAll()
	assert.Equal(t, buyers, got)

}

func TestGetAllFail(t *testing.T) {
	buyer1 := models.Buyer{1, 1234, "Bruno", "Santi"}
	buyer2 := models.Buyer{1, 1234, "Bruno", "Santi"}

	buyers = append(buyers, &buyer1, &buyer2)
	buyerslist := append(buyers, &buyer1)

	got := GetAll()
	assert.NotEqual(t, buyerslist, got)

}
