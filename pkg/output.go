package pkg

import (
	"fmt"
	"github.com/V1nic1us/github-activity-go/models"
)

func PrintOutput(response []models.Activity) {
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

	for _, activity := range response {
		time, err := FormatActivityDate(activity.CreatedAt)
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
