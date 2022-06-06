package utils

import (
	"regexp"
	"strconv"
)

//Função para formatar o CPF removendo
//tudo o que não for número
func removeNonDigits(cpf string) string {
	re, _ := regexp.Compile("[^0-9]+")
	cpf = re.ReplaceAllString(cpf, "")
	return cpf
}

// Função que converte um tipo rune para tipo inteiro
func toInt(r rune) int {
	return int(r - '0')
}

//calcular digito do documento
func calculateDigit(doc string, position int) string {

	var sum int
	for _, r := range doc {

		sum += toInt(r) * position
		position--

		if position < 2 {
			position = 9
		}
	}

	sum %= 11
	if sum < 2 {
		return "0"
	}

	return strconv.Itoa(11 - sum)
}

// Função para validar o CPF
func ValidateCPF(cpf string) bool {
	fmt_cpf := removeNonDigits(cpf)

	const (
		size     = 9
		position = 10
	)

	d := fmt_cpf[:size]
	digit := calculateDigit(d, position)
	d = d + digit
	digit = calculateDigit(d, position+1)

	return fmt_cpf == d+digit
}

// isEmailValid checks if the email provided is valid by regex.
func isEmailValid(e string) bool {
    emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
    return emailRegex.MatchString(e)
}