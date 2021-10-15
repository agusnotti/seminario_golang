# Seminario Golang 2021 - Trabajo Práctico

## Test Coverage: 87,5%

### Consigna: Crear una función que dada una cadena con un formato determinado genere una instancia de una estructura con los valores de los campos correspondientes al formato de la cadena. 
Por ejemplo:

````
TX03ABC
````

Deberá generar una estructura con los siguientes valores:

````
{TX 3 ABC}
````

Donde la estructura deberá tener una definicion del tipo:

```golang
type Result struct {
    Type    string 
    Value   string
    Length  int
}
```

La cadena de caracteres tiene el siguiente formato:

* los primeros dos caracteres son el tipo
* los segundos dos caracteres son el largo del valor
* los siguientes N caracteres son el valor, donde N es el valor del bullet anterior.

Entonces `NN040001` equivale a:

````
Type=NN
Length=04
Value=0001
````
----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------

Para llevar a cabo la resolución del ejercicio se realizaron cuatro funciones:

1. Dado una cadena ingresada como input, se encarga de procesar la misma (separandola en sus partes constituyentes: Tipo, Tamaño y Valor), retornando una estructura si esta cumple con las condiciones correctas.

```golang
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
```

2. Valida si el largo de la cadena ingresada como parámetro (valor) es igual al tamaño (tamanio) ingresado como parámetro.


```golang
func validarTamanio(tamanio int, valor string) bool {
	return tamanio == len(valor)
}
```

3. Dado el ***tipo*** de la cadena ingresada (expresado por los dos primeros caracteres del input) valida que el mismo se corresponda con el valor de la cadena (expresado a partir del cuarto caracter)


```golang
func validarTipo(tipo string, valor string) bool {

	if (tipo == "NN" && esNumero(valor)) || (tipo == "TX" && !esNumero(valor)) {
		return true
	}

	return false
}
```

4. Valida si el valor ingresado como parametro es un número o no. 

```golang
func esNumero(valor string) bool {
	val, _ := strconv.Atoi(valor)
	return val != 0
}
```