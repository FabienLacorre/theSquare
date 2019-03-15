// File: connect_neo4j.go
// File Created: 07 Mar 2019 16:42
// By Maxence Moutoussamy <maxence.moutoussamy1@gmail.com>

package utils

import (
	"time"

	bolt "github.com/moutoum/golang-neo4j-bolt-driver"
	"github.com/sirupsen/logrus"
)

const (
	timeBetweenRetry        = 1 * time.Second
	timeBeforeLeave         = 10 * time.Second
	maxNeo4jPoolConnections = 64
)

func ConnectNeo4j(url string) (driver bolt.ClosableDriverPool) {
	var err error

	ticker := time.Tick(timeBetweenRetry)
	limit := time.Now().Add(timeBeforeLeave)

retryLoop:
	for t := <-ticker; t.Before(limit); t = <-ticker {
		driver, err = bolt.NewClosableDriverPool(url, maxNeo4jPoolConnections)
		if err != nil {
			logrus.WithError(err).Warning("cannot open neo4j driver pool")
			continue
		} else {
			break retryLoop
		}
	}

	if err != nil {
		logrus.WithError(err).Fatal("cannot open neo4j driver pool")
	}

	return driver
}
