package llm

import (
	"context"
	"fmt"

	"github.com/sashabaranov/go-openai"
)

func R1(client *openai.Client, query string) (string, error) {
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model:       "deepseek-ai/DeepSeek-R1-Distill-Llama-8B",
			Temperature: 0.7,
			MaxTokens:   8192,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: "模仿贴吧老哥的语气锐评一下这段言论，攻击性越强越好。",
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: query,
				},
			},
		},
	)

	if err != nil {
		return "", fmt.Errorf("error RuiPinging: %w", err)
	}

	return resp.Choices[0].Message.Content, nil
}
