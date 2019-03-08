// File: auth.go
// File Created: 08 Mar 2019 11:20
// By Maxence Moutoussamy <maxence.moutoussamy1@gmail.com>

package internal

import (
	"encoding/gob"
	"net/http"

	sessions "github.com/goincremental/negroni-sessions"
)

type key int

const (
	connectedKey key = 0
	loginKey     key = 1
)

func init() {
	gob.Register(key(0))
}

func AuthMiddleware(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	session := sessions.GetSession(r)

	isConnected, ok := session.Get(connectedKey).(bool)
	if !ok || !isConnected {
		login, _, ok := r.BasicAuth()
		if !ok {
			session.Set(connectedKey, false)
			http.Error(rw, "invalid or missing authentication informations", http.StatusUnauthorized)
			return
		}

		// TODO: Check BDD

		session.Set(connectedKey, true)
		session.Set(loginKey, login)
	}

	next(rw, r)
}
