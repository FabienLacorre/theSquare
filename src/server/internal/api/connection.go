// File: connection.go
// File Created: 09 Mar 2019 09:06
// By Maxence Moutoussamy <maxence.moutoussamy1@gmail.com>

package api

import (
	"net/http"
	"server/internal"

	sessions "github.com/goincremental/negroni-sessions"
)

type ConnectionService struct {
}

func (s *ConnectionService) Login(rw http.ResponseWriter, req *http.Request) {
	rw.WriteHeader(http.StatusOK)
}

func (s *ConnectionService) Logout(rw http.ResponseWriter, req *http.Request) {
	session := sessions.GetSession(req)
	session.Set(internal.ConnectedKey, false)
	rw.WriteHeader(http.StatusOK)
}
