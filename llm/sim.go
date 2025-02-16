package llm

import (
	"fmt"
	"starBot/zilliz"

	"github.com/sashabaranov/go-openai"
)

func Sim(client *openai.Client, query string) (string, string, error) {
	sr, err := zilliz.Search(client, query)
	if err != nil {
		return "", "", fmt.Errorf("error searching: %w", err)
	}

	comment := "Similarity Search\nDebug: " + query + "\n"
	for i, s := range sr {
		comment += fmt.Sprintf("\nResult %d: %s", i, s)
	}

	choice, err := semanticMatch(client, query, sr)
	if err != nil {
		return "", "", fmt.Errorf("error matching: %w", err)
	}

	return comment, choice, nil
}
