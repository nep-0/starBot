package llm

import (
	"context"
	"fmt"
	"starBot/zilliz"

	"github.com/sashabaranov/go-openai"
)

func Comment(client *openai.Client, query string) (string, string, error) {
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model:       "Qwen/Qwen2.5-7B-Instruct",
			Temperature: 0.7,
			MaxTokens:   512,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: "无论用户输入什么，用一个张维为风格的语句略带嘲讽地流畅地接上。",
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: "我要玩原神。",
				},
				{
					Role:    openai.ChatMessageRoleAssistant,
					Content: "对不起，我们带着同情的眼光看着你",
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: "你很厉害嘛",
				},
				{
					Role:    openai.ChatMessageRoleAssistant,
					Content: "这里有中华文明的智慧",
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: "你是傻逼",
				},
				{
					Role:    openai.ChatMessageRoleAssistant,
					Content: "我说你不要敬酒不吃吃罚酒",
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: "我是院士。",
				},
				{
					Role:    openai.ChatMessageRoleAssistant,
					Content: "专业水平之低，令人汗颜",
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: "我爱吃屎。",
				},
				{
					Role:    openai.ChatMessageRoleAssistant,
					Content: "下面听众热烈鼓掌",
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: "我今天运气不好。",
				},
				{
					Role:    openai.ChatMessageRoleAssistant,
					Content: "不值得同情的",
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: "评价一下美国坠机事件。",
				},
				{
					Role:    openai.ChatMessageRoleAssistant,
					Content: "他有基因缺陷",
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: "我是傻逼。",
				},
				{
					Role:    openai.ChatMessageRoleAssistant,
					Content: "这几乎是个共识了",
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: query,
				},
			},
		},
	)
	if err != nil {
		return "", "", fmt.Errorf("error commenting: %w", err)
	}

	sr, err := zilliz.Search(client, resp.Choices[0].Message.Content)
	if err != nil {
		return "", "", fmt.Errorf("error searching: %w", err)
	}

	choice, err := match(client, query, sr)
	if err != nil {
		return "", "", fmt.Errorf("error matching: %w", err)
	}

	comment := resp.Choices[0].Message.Content + "\n"
	for _, s := range sr {
		comment += fmt.Sprintf("\n%s", s)
	}
	comment += fmt.Sprintf("\n\n检索结果：%s", choice)

	return comment, choice, nil
}
