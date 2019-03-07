// File: connect_neo4j.go
// File Created: 07 Mar 2019 16:42
// By Maxence Moutoussamy <maxence.moutoussamy1@gmail.com>

package utils

import (
	"errors"
	"time"

	bolt "github.com/moutoum/golang-neo4j-bolt-driver"
	"github.com/sirupsen/logrus"
)

const (
	timeBetweenRetry = 1 * time.Second
	timeBeforeLeave  = 10 * time.Second
)

func ConnectNeo4j(url string) (bolt.Conn, error) {
	var (
		conn bolt.Conn
		err  error
	)

	driv := bolt.NewDriver()
	ticker := time.Tick(timeBetweenRetry)
	limit := time.Now().Add(timeBeforeLeave)

	for t := <-ticker; t.Before(limit); t = <-ticker {
		if conn, err = driv.OpenNeo(url); err != nil {
			conn = nil
			logrus.WithError(err).Warn("cannot open neo4j conn")
		}
	}
	if conn == nil {
		return nil, errors.New("cannot open neo4j connection")
	}

	return conn, nil
}
