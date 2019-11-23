package chat

import (
	"github.com/go-park-mail-ru/2019_2_Pirogi/app/domain/model"
	"golang.org/x/net/websocket"
	"strings"
)

type Message struct {
	Body   string `json:"body"`
}

type MessageNew struct {
	UserID model.ID `json:"user_id"`
	Body   string   `json:"body"`
}

type Chat struct {
	UserID   model.ID  `json:"user_id"`
	Messages []Message `json:"messages"`
}

type Client struct {
	id     int
	ws     *websocket.Conn
	server *Server
	ch     chan *Message
	doneCh chan bool
}

type Server struct {
	pattern   string
	messages  []*Message
	clients   map[int]*Client
	addCh     chan *Client
	delCh     chan *Client
	sendAllCh chan *Message
	doneCh    chan bool
	errCh     chan *ErrorChat
}

type ServerChat interface {
	Add(c *Client)
	Del(c *Client)
	SendAll(msg *Message)
	Done()
	Err(err *ErrorChat)
	sendPastMessages(c *Client)
	Listen()
}

type ClientChat interface {
	Conn() *websocket.Conn
	Write(msg *Message)
	Done()
	Listen()
	listenWrite()
	listenRead()
}

func NewServer(pattern string) *Server {
	var messages []*Message
	clients := make(map[int]*Client)
	addCh := make(chan *Client)
	delCh := make(chan *Client)
	sendAllCh := make(chan *Message)
	doneCh := make(chan bool)
	errCh := make(chan *ErrorChat)

	return &Server{
		pattern,
		messages,
		clients,
		addCh,
		delCh,
		sendAllCh,
		doneCh,
		errCh,
	}
}

type ErrorChat struct {
	Message string
}

func NewErrorChat(messages ...string) *ErrorChat {
	return &ErrorChat{Message: strings.Join(messages, " ")}
}
