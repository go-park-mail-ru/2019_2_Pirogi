package chat

import (
	"fmt"
	"github.com/go-park-mail-ru/2019_2_Pirogi/app/domain/model"
	"go.uber.org/zap"
	"io"
	"log"

	"golang.org/x/net/websocket"
)

const channelBufSize = 100

// Create new chat client.
func NewClient(ws *websocket.Conn, server *Server, userID model.ID) *Client {
	if ws == nil {
		log.Fatal("ws can not be nil")
	}

	if server == nil {
		log.Fatal("server cannot be nil")
	}

	ch := make(chan *Message, channelBufSize)
	doneCh := make(chan bool)

	return &Client{userID, ws, server, ch, doneCh}
}

func (c *Client) Conn() *websocket.Conn {
	return c.ws
}

func (c *Client) Write(msg *Message) {
	select {
	case c.ch <- msg:
	default:
		c.server.Del(c)
		err := fmt.Errorf("client %d is disconnected.", c.id)
		c.server.Err(NewErrorChat(err.Error()))
	}
}

func (c *Client) Done() {
	c.doneCh <- true
}

// Listen Write and Read request via chanel
func (c *Client) Listen() {
	go c.listenWrite()
	c.listenRead()
}

// Listen write request via chanel
func (c *Client) listenWrite() {
	for {
		select {

		// send message to the client
		case msg := <-c.ch:
			log.Println("Send:", msg)
			websocket.JSON.Send(c.ws, msg)

		// receive done request
		case <-c.doneCh:
			c.server.Del(c)
			c.doneCh <- true // for listenRead method
			return
		}
	}
}

// Listen read request via chanel
func (c *Client) listenRead() {
	for {
		select {
		case <-c.doneCh:
			c.server.Del(c)
			c.doneCh <- true
			return

		default:
			var msg Message
			err := websocket.JSON.Receive(c.ws, &msg)
			if err == io.EOF {
				c.doneCh <- true
			} else if err != nil {
				c.server.Err(NewErrorChat(err.Error()))
			}
			if msg.Body != "" {
				zap.S().Debug(msg.Body)
				e := c.server.conn.Upsert(model.MessageNew{
					UserID: c.id,
					Body:   msg.Body,
				})
				if e != nil {
					c.server.Err(NewErrorChat(e.Error))
				}
			}
		}
	}
}
