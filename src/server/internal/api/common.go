// File: common.go
// File Created: 15 Mar 2019 07:36
// By Maxence Moutoussamy <maxence.moutoussamy1@gmail.com>

package api

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

func internalServerError(rw http.ResponseWriter, message string, err error) {
	if err != nil {
		logrus.WithError(err).Error(message)
	}
	http.Error(rw, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}
