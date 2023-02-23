package main

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestServicioThree_Exec(t *testing.T) {
	mockS1 := MockedServicioUno{}
	mockS2 := MockedServicioDos{}

	mockS1.On("SyncCall").Return("MOCKED SERVICIO PAPA")
	mockS2.On("AsyncCall").Return(errors.New("test mocked err"))

	s3 := NewServicioTres(&mockS1, &mockS2)

	result := s3.Exec()

	assert.NotNil(t, result)

}
