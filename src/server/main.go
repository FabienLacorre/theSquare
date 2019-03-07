// File: main.go
// File Created: 06 Mar 2019 14:58
// By Maxence Moutoussamy <maxence.moutoussamy1@gmail.com>

package main

import (
	"errors"
	"fmt"
	"os"
	"time"

	bolt "github.com/moutoum/golang-neo4j-bolt-driver"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

func run(ctx *cli.Context) error {

	var (
		conn bolt.Conn
		err  error
	)

	driv := bolt.NewDriver()
	ticker := time.Tick(1 * time.Second)
	limit := time.Now().Add(10 * time.Second)
	for t := <-ticker; t.Before(limit); t = <-ticker {
		if conn, err = driv.OpenNeo("bolt://neo4j:Ab;q*Ile@neo4j:7687"); err != nil {
			logrus.WithError(err).Warn("cannot open neo4j conn")
		}
	}
	if conn == nil {
		return errors.New("cannot open neo4j connection")
	}
	defer conn.Close()

	r, err := conn.QueryNeo("MATCH (n) RETURN n", nil)
	if err != nil {
		return fmt.Errorf("cannot query: %v", err)
	}

	fmt.Println(r)

	return nil
}

func main() {
	app := cli.NewApp()
	app.Name = os.Args[0]
	app.Usage = "Server for the TheSquare project"
	app.Action = run

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "neo4j-bolt-url",
			Usage: "Neo4j url to connect on",
			Value: "bolt://neo4j:7687",
		},
	}

	if err := app.Run(os.Args); err != nil {
		logrus.WithError(err).Error("app cannot run properly")
	}
}
