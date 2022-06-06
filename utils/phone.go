package utils

import (
	"fmt"
)

const (
	tmin = 12
	tMAX = 13
)

var dddValidation = []string{"11","12","13","14","15","16","17","18","19","21","22","24","27",
                             "28","31","32","33","34","35","37","38","41","42","43","44","45",
							 "46","47","48","49","51","53","54","55","61","62","63","64","65",
							 "66","67","68","69","71","73","74","75","77","79","81","82","83",
							 "84","85","86","87","88","89","91","92","93","94","95","96","97","98","99"}

func isValidBrazilianPhone () bool {
	// Constantes iniciais
	// Primeiro, vamos revmover tudo que não for número
	numero := removeNonDigits("+55 (12) 9 9670-1570")

	// O "numero" é uma string. Vamos conferir se é brasileiro:
	if numero[0:2] != "55"{
		fmt.Println("Telefone não é Brasileiro")
		return false
	}

	// Validar se tem a quantidade de numeros correto
	tamanho := len(numero)
	if tamanho != tmin && tamanho != tMAX {
		return false
	}

	// Verificar primeiro dígito para Mobile
	if len(numero) == tMAX {
		primeiro := numero[4:5]
		if primeiro != "6" && primeiro != "7" && primeiro != "8" && primeiro != "9" {
			return false
		}
	}

	// Verificar primeiro dígito para Fixo
	if len(numero) == tmin {
		primeiro := numero[4:5]
		if primeiro != "2" && primeiro != "3" && primeiro != "4" && primeiro != "5" {
			return false
		}
	}

	found := false
	numeroDDD := numero[2:4]
	// Agora, vamos verificar se o DDD é válido
	for _, ddd := range dddValidation {
		if ddd == numeroDDD {
			found = true
			break
		}
	}

	if found == false {
		return false
	}
	return true
}