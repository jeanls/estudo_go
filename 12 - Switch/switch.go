package main

import "fmt"

func diaDaSemana(numero int) string {
	switch numero {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda"
	case 3:
		return "Terça"
	case 4:
		return "Quarta"
	case 5:
		return "Quinta"
	case 6:
		return "Sexta"
	case 7:
		return "Sabado"
	default:
		return ""
	}
}

func diaDaSemana2(numero int) string {
	switch {
	case numero == 1:
		return "Domingo"
	default:
		return ""
	}
}

func main() {
	dia := diaDaSemana(1)
	fmt.Println(dia)
}
