package main

import (
	"errors"
	"github.com/stretchr/testify/mock"
)

type ServicioDos interface {
	AsyncCall() error
}

type ServicioTwo struct {
	nombreServicio string
}

func (s *ServicioTwo) AsyncCall() error {

	return errors.New("Servicio 2 err")
}

type MockedServicioDos struct {
	mock.Mock
}

func (m *MockedServicioDos) AsyncCall() error {
	args := m.Called()
	return args.Error(0)

}
