package chat

import (
	"go.uber.org/zap"
	"golang.org/x/net/websocket"
	"net/http"
)

func (s *Server) Add(c *Client) {
	s.addCh <- c
}

func (s *Server) Del(c *Client) {
	s.delCh <- c
}

func (s *Server) SendAll(msg *Message) {
	s.sendAllCh <- msg
}

func (s *Server) Done() {
	s.doneCh <- true
}

func (s *Server) Err(err *ErrorChat) {
	s.errCh <- err
}

func (s *Server) sendPastMessages(c *Client) {
	for _, msg := range s.messages {
		c.Write(msg)
	}
}

func (s *Server) Listen() {
	onConnected := func(ws *websocket.Conn) {
		defer func() {
			err := ws.Close()
			if err != nil {
				s.errCh <- NewErrorChat(err.Error())
			}
		}()

		client := NewClient(ws, s)
		s.Add(client)
		client.Listen()
	}
	http.Handle(s.pattern, websocket.Handler(onConnected))

	for {
		select {

		// Add new a client
		case c := <-s.addCh:
			zap.S().Error(NewErrorChat("Added client ", string(c.id)))
			s.clients[c.id] = c
			s.sendPastMessages(c)

		// del a client
		case c := <-s.delCh:
			zap.S().Error(NewErrorChat("Delete client ", string(c.id)))
			delete(s.clients, c.id)

		case err := <-s.errCh:
			zap.S().Error(NewErrorChat(err.Message))

		case <-s.doneCh:
			return
		}
	}
}
