package models

import "time"

type Message struct {
	KeycloakId string
	Text       string
	CreatedAt  time.Time
}
