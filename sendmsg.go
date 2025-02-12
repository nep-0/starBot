package main

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

var client = resty.New()

var endpoint string

func initOnebot(url string) {
	endpoint = url
}

func sendMsg(groupId int64, msg string) {
	_, err := client.R().
		SetBody(map[string]any{"group_id": groupId, "message": msg}).
		Post(endpoint + "/send_group_msg")
	if err != nil {
		fmt.Println(err)
	}
}

func sendImg(groupId int64, url string) {
	_, err := client.R().
		SetBody(map[string]any{
			"group_id": groupId,
			"message": map[string]any{
				"type": "image",
				"data": map[string]string{"file": url},
			},
		}).
		Post(endpoint + "/send_group_msg")
	if err != nil {
		fmt.Println(err)
	}
}
