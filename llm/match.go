package llm

import (
	"context"
	"fmt"
	"strconv"

	"github.com/sashabaranov/go-openai"
)

func match(client *openai.Client, query string, candidates []string) (string, error) {
	msgs := make([]openai.ChatCompletionMessage, 0, 4+len(candidates))
	msgs = append(msgs, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleSystem,
		Content: "你是幽默回复机器。用户输入被“<UserInput>”“</UserInput>”包裹；可选的回复被“<Response>”“</Response>”包裹。请你从可选的回复中挑选最幽默且通顺的**一项**。",
	})
	msgs = append(msgs, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleAssistant,
		Content: "例如：<UserInput>我喜欢你</UserInput><Index>0</Index><Response>这是中国文化巨大的智慧</Response><Index>1</Index><Response>可以说都感动了</Response><Index>2</Index><Response>这里有中华文明的智慧</Response><Index>3</Index><Response>这一切经得起任何国际比较</Response>你应该选择“可以说都感动了”，回复`1`。",
	})
	msgs = append(msgs, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleAssistant,
		Content: "<UserInput>" + query + "</UserInput>",
	})
	for i, c := range candidates {
		msgs = append(msgs, openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleAssistant,
			Content: "<Index>" + fmt.Sprint(i) + "</Index><Response>" + c + "</Response>",
		})
	}
	msgs = append(msgs, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: "现在请你挑选回复最幽默且通顺的一项。仅回复一个 Index 的**数字**，不要包含文本或标签“<”“>”。",
	})

	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model:       "Qwen/Qwen2.5-7B-Instruct",
			Temperature: 0.7,
			MaxTokens:   10,
			Messages:    msgs,
		},
	)
	if err != nil {
		return "", fmt.Errorf("error matching: %w", err)
	}

	index, err := strconv.ParseInt(resp.Choices[0].Message.Content, 10, 64)
	if err != nil || index < 0 || index >= int64(len(candidates)) {
		return "", fmt.Errorf("error matching: %w", err)
	}
	fmt.Println(candidates[index])

	return candidates[index], nil
}

func directMatch(client *openai.Client, query string, candidates []string) (string, error) {
	return candidates[0], nil
}
