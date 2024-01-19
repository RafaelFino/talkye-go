package entities

import (
	"time"
)

type Chat struct {
	Id         int       `json:"id"`
	Subject    string    `json:"subject"`
	CreateDate time.Time `json:"create_date"`
	Messages   []Message `json:"messages"`
}
