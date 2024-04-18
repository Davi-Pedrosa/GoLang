package main

//Importando "fmt". Ele usado para formatar textos e exibir
import "fmt"

func menu() {
    
	fmt.Println("Qual operação deseja fazer?")

	fmt.Println(("1) Soma? ")
	fmt.Println(("2) Subtração? ")


}

func main() {
	var num1, num2, resposta int
	var soma int = num1 + num2
	var subtracao int = num1 - num2

	fmt.Println("Bem-vindo(a) a calculadora!")

	fmt.Println("Digite o valor do primeiro número: ")
	fmt.Scanln(&num2)

	fmt.Println("Digite o valor do segundo número: ")
	fmt.Scanln(&num2)

	menu()
	fmt.Scanln(&resposta)
	switch resposta {
	case 1:
		fmt.Println(soma)
	case 2:
		fmt.Println(subtracao)
	default:
		fmt.Println("Número inválido!")
	}

}
