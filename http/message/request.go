package message

import "time"

type RequestSendMessage struct {
	KeycloakId string    `json:"keycloak_id"`
	ChatId     int       `json:"chat_id"`
	Text       string    `json:"text"`
	CreatedAt  time.Time `json:"created_at"`
}
