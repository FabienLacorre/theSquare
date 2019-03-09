// File: auth.go
// File Created: 08 Mar 2019 11:20
// By Maxence Moutoussamy <maxence.moutoussamy1@gmail.com>

package internal

import (
	"encoding/gob"
	"net/http"
	"server/internal/dao"

	"github.com/sirupsen/logrus"

	"github.com/urfave/negroni"

	sessions "github.com/goincremental/negroni-sessions"
)

type key int

const (
	ConnectedKey key = 0
	LoginKey     key = 1
)

func init() {
	gob.Register(key(0))
}

func AuthMiddleware(manager *dao.LoginManager) negroni.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		session := sessions.GetSession(r)

		isConnected, ok := session.Get(ConnectedKey).(bool)
		if !ok || !isConnected {
			login, password, ok := r.BasicAuth()
			if !ok {
				session.Set(ConnectedKey, false)
				http.Error(rw, "invalid or missing authentication informations", http.StatusUnauthorized)
				return
			}

			ok, err := manager.VerifyCreditentials(login, password)
			if err != nil {
				session.Set(ConnectedKey, false)
				http.Error(rw, "cannot verify creditentials", http.StatusInternalServerError)
				logrus.WithError(err).Error("cannot verify creditentials")
				return
			}

			if !ok {
				session.Set(ConnectedKey, false)
				http.Error(rw, "invalid creditentials informations", http.StatusUnauthorized)
				return
			}

			session.Set(ConnectedKey, true)
			session.Set(LoginKey, login)
		}

		next(rw, r)
	}
}
