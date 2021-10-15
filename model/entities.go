package model

type Result struct {
	Tipo    string
	Tamanio int
	Valor   string
}

func NewResult(tipo string, tamanio int, valor string) Result {
	return Result{tipo, tamanio, valor}
}
