package llm

import (
	"context"
	"fmt"

	"github.com/sashabaranov/go-openai"
)

func match(client *openai.Client, query string, candidates []string) (string, error) {
	msgs := make([]openai.ChatCompletionMessage, 0, 3+len(candidates))
	msgs = append(msgs, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleSystem,
		Content: "你是幽默回复机器。用户输入被“<UserInput>”“</UserInput>”包裹；可选的回复被“<Response>”“</Response>”包裹。请你从可选的回复中挑选最幽默且通顺的**一项**。",
	})
	msgs = append(msgs, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleAssistant,
		Content: "<UserInput>" + query + "</UserInput>",
	})
	for _, c := range candidates {
		msgs = append(msgs, openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleAssistant,
			Content: "<Response>" + c + "</Response>",
		})
	}
	msgs = append(msgs, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: "现在请你挑选回复最幽默且通顺的一项。仅回复一个 Response 的文本内容，不要包含标签“<”“>”。",
	})

	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model:       "Qwen/Qwen2.5-7B-Instruct",
			Temperature: 0.7,
			MaxTokens:   512,
			Messages:    msgs,
		},
	)
	if err != nil {
		return "", fmt.Errorf("error matching: %w", err)
	}

	return resp.Choices[0].Message.Content, nil

}

func directMatch(client *openai.Client, query string, candidates []string) (string, error) {
	return candidates[0], nil
}
