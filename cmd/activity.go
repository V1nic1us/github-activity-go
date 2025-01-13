package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"github.com/V1nic1us/github-activity-go/models"
	"github.com/V1nic1us/github-activity-go/pkg"
)

func fetchActivity(username string) {
	url := fmt.Sprintf("https://api.github.com/users/%s/events", username)

	resp, err := http.Get(url)

	if resp.StatusCode == 404 {
		fmt.Println("Usuário não encontrado.")
		return

	}

	if err != nil {
		fmt.Print("Erro ao fazer a requisição: ")
		fmt.Println(err)
		return
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	var response []models.Activity

	json.Unmarshal(data, &response)

	pkg.PrintOutput(response)
}