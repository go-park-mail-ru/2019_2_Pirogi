package model

import "time"

type Message struct {
	Datetime time.Time `json:"datetime"`
	Body     string    `json:"body"`
	// 0 - user, 1 - support
	Author bool `json:"author"`
}

type MessageNew struct {
	UserID ID     `json:"user_id"`
	Body   string `json:"body"`
}

type Chat struct {
	UserID   ID        `json:"user_id"`
	Messages []Message `json:"messages"`
}
