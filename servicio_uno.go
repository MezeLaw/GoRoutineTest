package main

import "github.com/stretchr/testify/mock"

type ServicioUno interface {
	SyncCall() string
}

type ServicioOne struct {
	nombreServicio string
}

func (s *ServicioOne) SyncCall() string {
	s.nombreServicio = "SERVICIO 1"
	return s.nombreServicio
}

type MockedServicioUno struct {
	mock.Mock
}

func (m *MockedServicioUno) SyncCall() string {
	args := m.Called()
	return args.String(0)

}
