package buyer

import (
	"github.com/stretchr/testify/assert"
	"pruebafinal/models"
	"testing"
)

func TestGetByCode(t *testing.T) {
	expected := models.Buyer{1, 1234, "Bruno", "Santi"}
	buyers = append(buyers, &expected)
	got, _ := GetByCode(expected.Code)
	assert.Equal(t, expected, got)

}

func TestGetByCodeFail(t *testing.T) {
	expected := models.Buyer{4, 4321, "Bruno", "Santi"}
	got, _ :=GetByCode(expected.Code)
	assert.NotEqual(t, expected, got)
}

