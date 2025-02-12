package main

import (
	"fmt"
	"starBot/config"
	"starBot/llm"
	"starBot/zilliz"

	"github.com/gin-gonic/gin"
	openai "github.com/sashabaranov/go-openai"
)

type message struct {
	rawMessage string
	id         int64
	groupId    int64
}

var msgChannel chan message = make(chan message, 40)

func reply(conf *config.Config) {
	clientConfig := openai.DefaultConfig(conf.Openai.ApiKey)
	clientConfig.BaseURL = conf.Openai.BaseURL
	client := openai.NewClientWithConfig(clientConfig)

	for {
		msg := <-msgChannel
		if msg.rawMessage[0:2] == "r1" {
			// r1 bot
			msg.rawMessage = msg.rawMessage[2:]
			reply, err := llm.R1(client, msg.rawMessage)
			if err != nil {
				sendMsg(msg.groupId, "error generating comment: "+err.Error())
			}
			sendMsg(msg.groupId, reply)

		} else if msg.rawMessage[0:3] == " r1" {
			// r1 bot
			msg.rawMessage = msg.rawMessage[3:]
			reply, err := llm.R1(client, msg.rawMessage)
			if err != nil {
				sendMsg(msg.groupId, "error generating comment: "+err.Error())
			}
			sendMsg(msg.groupId, reply)

		} else {
			// vv bot
			_, choice, err := llm.Comment(client, msg.rawMessage)
			if err != nil {
				sendMsg(msg.groupId, "error generating comment: "+err.Error())
				continue
			}
			// sendMsg(msg.groupId, comment)
			sendImg(msg.groupId, conf.Static.VvRoot+choice+".png")
		}
	}

}

func main() {
	conf, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}
	initOnebot(conf.Openai.ApiKey)
	zilliz.InitZilliz(conf.Zilliz.Url, conf.Zilliz.BearerToken)

	// concurrency: 4
	go reply(conf)
	go reply(conf)
	go reply(conf)
	go reply(conf)

	r := gin.Default()
	r.POST("/", func(c *gin.Context) {
		json := make(map[string]interface{})
		c.BindJSON(&json)

		fmt.Println(json)

		if json["post_type"] == "message" {
			if json["message_type"] == "group" {
				rawMessage, ok := json["raw_message"].(string)
				if !ok {
					c.JSON(500, nil)
					return
				}
				id, ok := json["message_seq"].(float64)
				if !ok {
					c.JSON(500, nil)
					return
				}
				groupId, ok := json["group_id"].(float64)
				if !ok {
					c.JSON(500, nil)
					return
				}

				if len(rawMessage) > len(conf.OneBot.QQ)+11 {
					if rawMessage[0:len(conf.OneBot.QQ)+11] == "[CQ:at,qq="+conf.OneBot.QQ+"]" {
						msgChannel <- message{rawMessage: rawMessage[len(conf.OneBot.QQ)+11:], id: int64(id), groupId: int64(groupId)}
					}
				}
				if len(rawMessage) > len(conf.OneBot.Nickname)+1 {
					if rawMessage[0:len(conf.OneBot.Nickname)+1] == "@"+conf.OneBot.Nickname {
						msgChannel <- message{rawMessage: rawMessage[len(conf.OneBot.Nickname)+1:], id: int64(id), groupId: int64(groupId)}
					}
				}

			}
		}
		c.JSON(200, nil)
	})
	r.Run(conf.OneBot.Listen)
}
