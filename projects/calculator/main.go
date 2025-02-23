package main

import "fmt"

func ReadUserInput(inputType string) (string, float32, error) {
	var numero float32
	var operador string
	switch inputType {
	case "number":
		fmt.Println("Digite um Numero:")
		_, err := fmt.Scanln(&numero)
		if err != nil {
			return operador, numero, err
		}
		return operador, numero, nil
	case "operator":
		fmt.Println("Digite uma operacao:")
		_, err := fmt.Scanln(&operador)
		if err != nil {
			fmt.Println("Erro:", err)
			return operador, numero, err
		}
		return operador, numero, nil
	default:
		return operador, numero, fmt.Errorf("Tipo nao identificado")
	}
}

func calcula() {
	_, primeiroNumero, _ := ReadUserInput("number")
	_, segundoNumero, _ := ReadUserInput("number")
	operador, _, _ := ReadUserInput("operator")

	resultado := operacoes(primeiroNumero, segundoNumero, operador)

	fmt.Println(resultado)

}

func operacoes(primeiroNumero, segundoNumero float32, operador string) float32 {
	switch {
	case operador == "+":
		return primeiroNumero + segundoNumero
	case operador == "-":
		return primeiroNumero - segundoNumero
	case operador == "*":
		return primeiroNumero * segundoNumero
	case operador == "/":
		if segundoNumero == 0 {
			fmt.Println("Não há divisão por 0")
			return 0
		}
		return primeiroNumero / segundoNumero
	default:
		fmt.Println("Operador não Identificado")
		return 0
	}
}
func main() {
	calcula()
}
