package main

type ServicioTres interface {
	Exec() string
}

type ServicioThree struct {
	S1 ServicioUno
	S2 ServicioDos
}

func NewServicioTres(s1 ServicioUno, s2 ServicioDos) ServicioTres {
	return &ServicioThree{
		S1: s1,
		S2: s2,
	}
}

func (s3 *ServicioThree) Exec() string {

	stringS1 := s3.S1.SyncCall()
	go func() {
		err := s3.S2.AsyncCall()
		if err != nil {
			println("ENTRE AL ERR DEL GO ROUTINE")
		}
	}()
	return stringS1
}
