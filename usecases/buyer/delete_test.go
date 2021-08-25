package buyer

import (
	"github.com/stretchr/testify/assert"
	"pruebafinal/models"
	"testing"
)

func TestDelete(t *testing.T) {
	buyer := models.Buyer{Code: 1234, Name: "Bruno", LastName: "Santi"}
	Save(buyer)
	Delete(1)
	expected := models.Buyer{ID:0, Code:0, Name:"", LastName:""}
	got, _ := Get(buyer.ID)
	assert.Equal(t, expected, got)
}


func TestDeleteFail(t *testing.T) {
	buyer := models.Buyer{Code: 1234, Name: "Bruno", LastName: "Santi"}
	Save(buyer)
	Delete(4)
	expected := "Buyer with ID '4' not found"
	var got = "Buyer with ID '4' not found"
	assert.Equal(t, expected, got)
}

