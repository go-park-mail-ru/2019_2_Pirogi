package chat

import (
	"github.com/go-park-mail-ru/2019_2_Pirogi/app/domain/model"
	"github.com/go-park-mail-ru/2019_2_Pirogi/configs"
	"go.uber.org/zap"
	"golang.org/x/net/websocket"
	"log"
	"net/http"
)

func (s *Server) Add(c *Client) {
	s.addCh <- c
}

func (s *Server) Del(c *Client) {
	s.delCh <- c
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
		var cookie model.Cookie
		cookieHTTP, err := ws.Request().Cookie(configs.Default.CookieAuthName)
		if err != nil {
			s.errCh <- NewErrorChat(err.Error())
			return
		}
		cookie.CopyFromCommon(cookieHTTP)
		zap.S().Debug("Cookie http", cookieHTTP)
		u, e := s.conn.FindUserByCookie(cookieHTTP)
		if e != nil {
			s.errCh <- NewErrorChat(e.Error)
			return
		}
		client := NewClient(ws, s, u.ID)
		s.Add(client)
		client.Listen()
	}
	http.Handle(s.pattern, websocket.Handler(onConnected))

	for {
		select {
		// Add new a client
		case c := <-s.addCh:
			zap.S().Error("Added client ", c.id.String())
			s.clients[c.id] = c
			s.sendPastMessages(c)

		// del a client
		case c := <-s.delCh:
			zap.S().Error("Delete client ", c.id.String())
			err := c.ws.Close()
			if err != nil {
				log.Fatal(err)
			}
			delete(s.clients, c.id)
		case err := <-s.errCh:
			zap.S().Error(NewErrorChat(err.Message))
		case <-s.doneCh:
			return
		}
	}
}
