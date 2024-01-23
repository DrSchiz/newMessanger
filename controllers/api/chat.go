package controllers

import (
	"encoding/json"
	"messanger/http/message"
	"messanger/models"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"

	"github.com/gin-gonic/gin"
)

func SendMessage(c *gin.Context) {

	Rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	var requestMessage message.RequestSendMessage
	json.NewDecoder(c.Request.Body).Decode(&requestMessage)

	var messages []models.Message
	messagesJson := Rdb.Get(c, "chat:"+strconv.Itoa(requestMessage.ChatId)).Val()

	if messagesJson != "" {
		err := json.Unmarshal([]byte(messagesJson), &messages)
		if err != nil {
			panic(err)
		}
	}

	messages = append(messages, models.Message{
		KeycloakId: requestMessage.KeycloakId,
		Text:       requestMessage.Text,
		CreatedAt:  time.Now(),
	})

	marshalMessages, err := json.Marshal(messages)
	if err != nil {
		panic(err)
	}

	err = Rdb.Set(c, "chat:"+strconv.Itoa(requestMessage.ChatId), marshalMessages, 0).Err()
	if err != nil {
		panic(err)
	}
}
