package entities

import (
	"time"
)

type Message struct {
	Id         int       `json:"id"`
	Username   string    `json:"username"`
	CreateDate time.Time `json:"create_date"`
	Subject    string    `json:"subject"`
	Body       string    `json:"body"`
}
