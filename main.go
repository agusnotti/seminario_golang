package main

import (
	"fmt"
	"strconv"

	"tudai2021.com/model"
)

func main() {
	e := procesarCadena("TX03ABC")

	fmt.Println(e)
}

func procesarCadena(input string) model.Result {
	tipo := input[0:2]
	valor := input[4:]
	tamanio, _ := strconv.Atoi(input[2:4])

	resultado := model.NewResult("", 0, "")

	if validarTamanio(tamanio, valor) {
		if validarTipo(tipo, valor) {
			resultado = model.NewResult(tipo, tamanio, valor)
		}
	}

	return resultado
}

func validarTamanio(tamanio int, valor string) bool {
	return tamanio == len(valor)
}

func validarTipo(tipo string, valor string) bool {

	if (tipo == "NN" && esNumero(valor)) || (tipo == "TX" && !esNumero(valor)) {
		return true
	}

	return false
}

func esNumero(valor string) bool {
	val, _ := strconv.Atoi(valor)
	return val != 0
}
