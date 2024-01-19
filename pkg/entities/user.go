package entities

import "time"

type User struct {
	Id         int       `json:"id"`
	Username   string    `json:"username"`
	CreateDate time.Time `json:"create_date"`
}
