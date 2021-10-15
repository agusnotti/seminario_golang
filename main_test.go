package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParser(t *testing.T) {
	var cases = []struct {
		Input   string // input string in order to be parsed
		Success bool   // paser result
		Type    string // the input type
		Value   string // the input value
		Length  int    // value length
	}{
		{"TX02AB", true, "TX", "AB", 2},
		{"TX0222", false, "", "", 0},
		{"NN100987654321", true, "NN", "0987654321", 10},
		{"TX06ABCDE", false, "", "", 0},
		{"NN04000A", false, "", "", 0},
	}

	for _, testData := range cases {

		//acá agregar chequeos propios del test por ejemplo:
		test := procesarCadena(testData.Input)
		assert.Equal(t, test.Tipo, testData.Type, testData.Success)
		assert.Equal(t, test.Valor, testData.Value, testData.Success)
		assert.Equal(t, test.Tamanio, testData.Length, testData.Success)
	}
}

//video clase 2.1 min 18.58
//go test -coverprofile=out.test
// go tool cover -html out.test -o out.html
