package buyer

import (
	"github.com/stretchr/testify/assert"
	"pruebafinal/models"
	"testing"
)

func TestGetById(t *testing.T) {
	expected := models.Buyer{1,1234,"Bruno", "Santi"}
	buyers = append(buyers, &expected)

	got, _ := Get(expected.ID)
	if got != expected {
		t.Error("No existe")
	} else{
		println("Existe el buyer con id" , got.ID)
	}
}

func TestGetByIdFail(t *testing.T) {
	expected := models.Buyer{4,4321,"Bruno", "Santi"}
	got, _ := Get(expected.ID)
	assert.NotEqual(t, expected, got)

}