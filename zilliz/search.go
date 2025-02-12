package zilliz

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/sashabaranov/go-openai"
)

type result struct {
	Text     string  `json:"text"`
	Distance float64 `json:"distance"`
}

type resultBody struct {
	Data []result `json:"data"`
	Code int64    `json:"code"`
	Cost int64    `json:"cost"`
}

var url string
var bearerToken string

func InitZilliz(zillizURL, zillizBearerToken string) {
	url = zillizURL
	bearerToken = zillizBearerToken
}

func Search(client *openai.Client, text string) ([]string, error) {
	resp, err := client.CreateEmbeddings(
		context.Background(),
		openai.EmbeddingRequest{
			Input: text,
			Model: "BAAI/bge-m3",
		},
	)
	if err != nil {
		return nil, err
	}

	payload, err := json.Marshal(map[string]any{
		"collectionName": "vv",
		"data":           [][]float32{resp.Data[0].Embedding},
		"limit":          5,
		"outputFields":   []string{"text", "distance"},
	})
	if err != nil {
		return nil, err
	}

	req, _ := http.NewRequest("POST", url, bytes.NewReader(payload))

	req.Header.Add("Authorization", "Bearer "+bearerToken)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(string(body))

	var results resultBody
	if err := json.Unmarshal(body, &results); err != nil {
		return nil, err
	}

	var texts []string
	for _, r := range results.Data {
		texts = append(texts, r.Text)
	}

	return texts, nil
}
