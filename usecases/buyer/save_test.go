package buyer

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"pruebafinal/models"
	"testing"
)

func TestSave(t *testing.T) {
	excpected := models.Buyer{ID: 1, Code: 1234, Name: "Bruno", LastName: "Santi"}
	buyer := models.Buyer{Code: 1234, Name: "Bruno", LastName: "Santi"}
	i, err := Save(buyer)
	if i != excpected {
		t.Errorf("El metodo Save fallo, esperaba '%v', obtuvo '%v'", excpected, i)

	}
	if err != nil {
		t.Errorf("El metodo Save fallo, esperaba '%v', obtuvo '%v'", nil, err)

	}
}
func TestSaveId(t *testing.T) {
	buyer := models.Buyer{ID: 6, Code: 1234, Name: "Bruno", LastName: "Santi"}
	_,err := Save(buyer)
	expected := errors.New("New buyer must not include id")

	assert.Equal(t, expected,err)
}
func TestSaveFail(t *testing.T) {
	excpected := models.Buyer{ID: 4, Code: 1234, Name: "Bruno", LastName: "Santi"}
	buyer := models.Buyer{Code: 1234, Name: "Bruno", LastName: "Santi"}
	i, err := Save(buyer)
	assert.NotEqual(t, excpected, i)
	if err != nil {
		t.Errorf("El metodo Save fallo, esperaba '%v', obtuvo '%v'", nil, err)

	}

}
