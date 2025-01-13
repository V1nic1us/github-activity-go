package pkg

import (
	"time"
	"fmt"
)

const (
	utcLayout   = "2006-01-02T15:04:05Z"
	localLayout = "02/01/2006 15:04:05"
)

func FormatActivityDate(createdAt string) (string, error) {
	utcTime, err := time.Parse(utcLayout, createdAt)
	if err != nil {
		return "", fmt.Errorf("erro ao analisar a data: %v", err)
	}
	localTime := utcTime.In(time.Now().Location())
	return localTime.Format(localLayout), nil
}