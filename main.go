package main

func main() {
	s1 := new(ServicioUno)
	s2 := new(ServicioDos)
	s3 := NewServicioTres(*s1, *s2)

	result := s3.Exec()

	print(result)

}
