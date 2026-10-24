package main

import "strconv"

// ValidarIdade converte o texto e valida o intervalo.
// Em falha, devolve *ErrValidacao (erro estruturado).
func ValidarIdade(texto string) (int, error) {
	idade, err := strconv.Atoi(texto)
	if err != nil {
		return 0, &ErrValidacao{
			Campo: "idade",
			Valor: texto,
			Msg:   "não é um número",
			causa: err,
		}
	}

	if idade < 0 || idade > 150 {
		return 0, &ErrValidacao{
			Campo: "idade",
			Valor: texto,
			Msg:   "fora do intervalo 0-150",
		}
	}

	return idade, nil
}
