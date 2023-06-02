package main

import (
	"fmt"
	"math/rand"
	"time"
)

func gerarAleatorios() int {
	fonte := rand.NewSource(time.Now().UnixNano())
	gerador := rand.New(fonte)
	return gerador.Intn(100) + 1
}

func adivinhar() int {
	resposta := gerarAleatorios()
	tentativas := 0

	fmt.Println("Bem-vindo ao jogo da adivinhação!")

	for {
		var tentativa int
		fmt.Println("Digite um número entre 1 e 100: ")
		fmt.Scanln(&tentativa)
		tentativas++

		if tentativa > resposta {
			fmt.Println("O número é menor. Tente novamente.")
		} else if tentativa < resposta {
			fmt.Println("O número é maior. Tente novamente.")
		} else {
			fmt.Println("Parabéns, você acertou!")
			break
		}
	}
	fmt.Printf("Você utilizou %d tentativas.\n", tentativas)
	return tentativas
}

func main() {
	var quantidadesTentativas [][]int

	for {
		tentativas := adivinhar()
		quantidadesTentativas = append(quantidadesTentativas, []int{tentativas})

		var novamente string
		fmt.Println("Deseja jogar novamente? (s/n): ")
		fmt.Scanln(&novamente)
		if novamente != "s" && novamente != "S" {
			break
		}
	}

	for i, jogadas := range quantidadesTentativas {
		fmt.Printf("No jogo %d", i+1)

		if len(jogadas) == 0 {
			fmt.Println("Nenhum jogo")
		} else {
			fmt.Printf(" você utilizou %d tentativa(s).\n", jogadas[0])
		}
	}
}
