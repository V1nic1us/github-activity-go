package cmd

import (
	"fmt"
	"os"
)

func Execute()  {
	if len(os.Args) < 2 {
		fmt.Println("Por favor, informe o nome de usuário.")
		return
	}

	username := os.Args[1]

	if username == "" {
		fmt.Println("O nome de usuário não pode ser vazio.")
		return
	}

	fetchActivity(username)
}