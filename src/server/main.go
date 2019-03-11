// File: main.go
// File Created: 06 Mar 2019 14:58
// By Maxence Moutoussamy <maxence.moutoussamy1@gmail.com>

package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	"server/internal"
	"server/internal/api"
	"server/internal/dao"
	"server/internal/utils"

	sessions "github.com/goincremental/negroni-sessions"
	"github.com/goincremental/negroni-sessions/cookiestore"
	"github.com/gorilla/mux"
	negronilogrus "github.com/meatballhat/negroni-logrus"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"github.com/urfave/negroni"
)

var store = cookiestore.New([]byte("secret-session-key"))

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
	url, err := createNeo4jURL(ctx)
	if err != nil {
		return err
	}

	logrus.WithField("url", url).Info("Trying to connect to Neo4j")
	conn, err := utils.ConnectNeo4j(url)
	if err != nil {
		return err
	}
	defer conn.Close()

	profileManager := dao.NewProfileManager(conn)

	connectionService := api.NewConnectionService(profileManager)
	profileService := api.NewProfileService(profileManager)

	router := mux.NewRouter()

	router.HandleFunc("/sign-in", connectionService.SignIn).
		Methods("POST", "PUT").
		HeadersRegexp("Content-Type", "application/json")

	apiRouter := mux.NewRouter()
	apiRouter.HandleFunc("/api/login", connectionService.Login).Methods("POST")
	apiRouter.HandleFunc("/api/logout", connectionService.Logout).Methods("POST")
	apiRouter.HandleFunc("/api/profile/{id:[0-9]+}", profileService.GetByID).Methods("GET")
	router.PathPrefix("/api/").Handler(negroni.New(
		negronilogrus.NewMiddleware(),
		sessions.Sessions("session", store),
		internal.AuthMiddleware(profileManager),
		negroni.Wrap(apiRouter),
	))

	router.PathPrefix("/").Handler(negroni.New(
		negronilogrus.NewMiddleware(),
		negroni.NewStatic(http.Dir("../../public/app"))),
	)

	logrus.Infof("Listening on port 8080...")
	http.ListenAndServe(":8080", router)

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
