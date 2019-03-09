// File: connection.go
// File Created: 09 Mar 2019 09:06
// By Maxence Moutoussamy <maxence.moutoussamy1@gmail.com>

package api

import (
	"net/http"
	"server/internal"

	"github.com/emicklei/go-restful"
	sessions "github.com/goincremental/negroni-sessions"
)

type ConnectionService struct {
}

func (s *ConnectionService) Register(root string) *restful.WebService {
	ws := new(restful.WebService)
	ws.Path(root)
	ws.Route(ws.POST("/login").To(s.login).Doc("Login the user"))
	ws.Route(ws.POST("/logout").To(s.logout).Doc("Logout the user"))
	return ws
}

func (s *ConnectionService) login(req *restful.Request, resp *restful.Response) {
	resp.WriteHeader(http.StatusOK)
}

func (s *ConnectionService) logout(req *restful.Request, resp *restful.Response) {
	session := sessions.GetSession(req.Request)
	session.Set(internal.ConnectedKey, false)
	resp.WriteHeader(http.StatusOK)
}
