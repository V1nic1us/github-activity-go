package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

type Activity struct {
	Id           string  `json:"id"`
	TypeActivity string  `json:"type"`
	Actor        Actor   `json:"actor"`
	Repo         Repo    `json:"repo"`
	Payload      Payload `json:"payload"`
	Public       bool    `json:"public"`
	CreatedAt    string  `json:"created_at"`
}

type Actor struct {
	Id           int    `json:"id"`
	Login        string `json:"login"`
	DisplayLogin string `json:"display_login"`
	GravatarId   string `json:"gravatar_id"`
	Url          string `json:"url"`
	AvatarUrl    string `json:"avatar_url"`
}

type Repo struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Url  string `json:"url"`
}

type Payload struct {
	Action       string   `json:"action,omitempty"`
	RepositoryId int      `json:"repository_id,omitempty"`
	Push_id      int      `json:"push_id,omitempty"`
	Size         int      `json:"size,omitempty"`
	DistinctSize int      `json:"distinct_size,omitempty"`
	Ref          string   `json:"ref,omitempty"`
	Head         string   `json:"head,omitempty"`
	Before       string   `json:"before,omitempty"`
	Commits      []Commit `json:"commits,omitempty"`
	RefType      string   `json:"ref_type,omitempty"`
	MasterBranch string   `json:"master_branch,omitempty"`
	Description  string   `json:"description,omitempty"`
	PusherType   string   `json:"pusher_type,omitempty"`
}

type Commit struct {
	Sha      string `json:"sha"`
	Author   Author `json:"author"`
	Message  string `json:"message"`
	Distinct bool   `json:"distinct"`
	Url      string `json:"url"`
}

type Author struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

func formatActivityDate(createdAt string) (string, error) {
	utcTime, err := time.Parse(utcLayout, createdAt)
	if err != nil {
		return "", fmt.Errorf("erro ao analisar a data: %v", err)
	}
	localTime := utcTime.In(time.Now().Location())
	return localTime.Format(localLayout), nil
}

const (
	utcLayout   = "2006-01-02T15:04:05Z"
	localLayout = "02/01/2006 15:04:05"
)

func main() {
	eventTypes := map[string]string{
		"CommitCommentEvent":            "Um comentário foi feito em ",
		"CreateEvent":                   "Um repositório, branch ou tag foi criado em",
		"DeleteEvent":                   "Um branch ou tag foi deletado em",
		"ForkEvent":                     "Um repositório foi bifurcado (forked).",
		"GollumEvent":                   "Uma página wiki foi criada ou atualizada em ",
		"IssueCommentEvent":             "Um comentário foi adicionado a uma issue ou pull request em ",
		"IssuesEvent":                   "Uma issue foi aberta, fechada ou reaberta em ",
		"MemberEvent":                   "Um colaborador foi adicionado a um repositório em ",
		"PublicEvent":                   "Um repositório foi tornado público.",
		"PullRequestEvent":              "Uma pull request foi aberta, fechada, mesclada ou sincronizada em ",
		"PullRequestReviewEvent":        "Uma revisão de pull request foi submetida em ",
		"PullRequestReviewCommentEvent": "Um comentário foi feito em uma revisão de pull request em",
		"PushEvent":                     "Commits foram enviados (push) para um branch em ",
		"ReleaseEvent":                  "Uma release foi publicada em ",
		"SponsorshipEvent":              "Uma ação relacionada a patrocínio ocorreu em ",
		"WatchEvent":                    "Um usuário começou a \"observar\" (watch) o repositório ",
	}

	if len(os.Args) < 2 {
		fmt.Println("Por favor, informe o nome de usuário.")
		return
	}

	username := os.Args[1]

	if username == "" {
		fmt.Println("O nome de usuário não pode ser vazio.")
		return
	}

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

	var response []Activity

	json.Unmarshal(data, &response)

	for _, activity := range response {
		time, err := formatActivityDate(activity.CreatedAt)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("\033[31mId: %v\033[0m - \033[32m%v\033[0m - \033[34m%v\033[0m - %v",
			activity.Id,
			eventTypes[activity.TypeActivity],
			activity.Repo.Name,
			time)
		fmt.Println()
	}

}
