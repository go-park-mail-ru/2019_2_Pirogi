package chat

import (
	"github.com/go-park-mail-ru/2019_2_Pirogi/app/domain/model"
	"github.com/go-park-mail-ru/2019_2_Pirogi/app/infrastructure/database"
	"golang.org/x/net/websocket"
	"strings"
)

type Client struct {
	id       model.ID
	ws       *websocket.Conn
	server   *Server
	ch       chan model.Message
	doneCh   chan bool
	messages []model.Message
}

type Server struct {
	pattern string
	clients map[model.ID]*Client
	addCh   chan *Client
	delCh   chan *Client
	doneCh  chan bool
	errCh   chan *ErrorChat
	conn    database.Database
}

type ServerChat interface {
	Add(c *Client)
	Del(c *Client)
	SendAll(msg *model.Message)
	Done()
	Err(err *ErrorChat)
	sendPastMessages(c *Client)
	Listen()
}

type ClientChat interface {
	Conn() *websocket.Conn
	Write(msg *model.Message)
	Done()
	Listen()
	listenWrite()
	listenRead()
}

func NewServer(pattern string, conn database.Database) *Server {
	clients := make(map[model.ID]*Client)
	addCh := make(chan *Client)
	delCh := make(chan *Client)
	doneCh := make(chan bool)
	errCh := make(chan *ErrorChat)

	return &Server{
		pattern,
		clients,
		addCh,
		delCh,
		doneCh,
		errCh,
		conn,
	}
}

type ErrorChat struct {
	Message string
}

func NewErrorChat(messages ...string) *ErrorChat {
	return &ErrorChat{Message: strings.Join(messages, " ")}
}
