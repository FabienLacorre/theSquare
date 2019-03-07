// File: main.go
// File Created: 06 Mar 2019 14:58
// By Maxence Moutoussamy <maxence.moutoussamy1@gmail.com>

package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/davecgh/go-spew/spew"
	bolt "github.com/moutoum/golang-neo4j-bolt-driver"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

func createNeo4jURL(ctx *cli.Context) (string, error) {
	login := ctx.GlobalString("neo4j-login")
	if len(login) == 0 {
		return "", errors.New("missing neo4j login")
	}

	password := ctx.GlobalString("neo4j-password")
	if len(password) == 0 {
		return "", errors.New("missing neo4j password")
	}

	url := ctx.GlobalString("neo4j-bolt-url")
	if len(url) == 0 {
		return "", errors.New("missing neo4j url")
	}

	return fmt.Sprintf("bolt://%s:%s@%s", login, password, url), nil
}

func run(ctx *cli.Context) error {

	var (
		conn bolt.Conn
		err  error
	)

	url, err := createNeo4jURL(ctx)
	if err != nil {
		return err
	}

	driv := bolt.NewDriver()
	ticker := time.Tick(1 * time.Second)
	limit := time.Now().Add(10 * time.Second)

	for t := <-ticker; t.Before(limit); t = <-ticker {
		if conn, err = driv.OpenNeo(url); err != nil {
			conn = nil
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

	spew.Config.Indent = "\t"

	for data, _, err := r.NextNeo(); err != io.EOF; data, _, err = r.NextNeo() {
		spew.Dump(data)
	}

	return nil
}

func main() {
	app := cli.NewApp()
	app.Name = os.Args[0]
	app.Usage = "Server for the TheSquare project"
	app.Action = run

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "neo4j-bolt-url",
			Usage:  "Neo4j url to connect on",
			Value:  "bolt://neo4j:7687",
			EnvVar: "NEO4J_BOLT_URL",
		},
		cli.StringFlag{
			Name:   "neo4j-login",
			Usage:  "Account login",
			EnvVar: "NEO4J_LOGIN",
		},
		cli.StringFlag{
			Name:   "neo4j-password",
			Usage:  "Account password",
			EnvVar: "NEO4J_PASSWORD",
		},
	}

	if err := app.Run(os.Args); err != nil {
		logrus.WithError(err).Error("app cannot run properly")
	}
}
