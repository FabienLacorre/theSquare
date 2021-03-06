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
	driver := utils.ConnectNeo4j(url)
	defer driver.Close()

	profileManager := dao.NewProfileManager(driver)
	companyManager := dao.NewCompanyManager(driver)
	hobbyManager := dao.NewHobbyManager(driver)
	jobManager := dao.NewJobManager(driver)
	skillManager := dao.NewSkillManager(driver)

	connectionService := api.NewConnectionService(profileManager)

	router := mux.NewRouter()

	// sign in
	router.HandleFunc("/sign-in", connectionService.SignIn).
		Methods("POST", "PUT").
		HeadersRegexp("Content-Type", "application/json")

	apiRouter := mux.NewRouter()

	// auths
	apiRouter.HandleFunc("/api/login", connectionService.Login).Methods("POST")
	apiRouter.HandleFunc("/api/logout", connectionService.Logout).Methods("POST")

	// profile
	profileService := api.NewProfileService(profileManager)
	apiRouter.HandleFunc("/api/profile/{id:[0-9]+}", profileService.Get).Methods("GET")
	apiRouter.HandleFunc("/api/profile/{id:[0-9]+}/companies", profileService.GetCompanies).Methods("GET")
	apiRouter.HandleFunc("/api/profile/{id:[0-9]+}/companies/{company_id:[0-9]+}", profileService.IsLikingCompany).Methods("GET")
	apiRouter.HandleFunc("/api/profile/{id:[0-9]+}/companies/{company_id:[0-9]+}", profileService.PostCompany).Methods("POST")
	apiRouter.HandleFunc("/api/profile/{id:[0-9]+}/companies/{company_id:[0-9]+}", profileService.DeleteCompany).Methods("DELETE")
	apiRouter.HandleFunc("/api/profile/{id:[0-9]+}/hobbies", profileService.GetHobbies).Methods("GET")
	apiRouter.HandleFunc("/api/profile/{id:[0-9]+}/hobbies/{hobby_id:[0-9]+}", profileService.IsLikingHobby).Methods("GET")
	apiRouter.HandleFunc("/api/profile/{id:[0-9]+}/hobbies/{hobby_id:[0-9]+}", profileService.PostHobby).Methods("POST")
	apiRouter.HandleFunc("/api/profile/{id:[0-9]+}/hobbies/{hobby_id:[0-9]+}", profileService.DeleteHobby).Methods("DELETE")
	apiRouter.HandleFunc("/api/profile/{id:[0-9]+}/skills", profileService.GetSkills).Methods("GET")
	apiRouter.HandleFunc("/api/profile/{id:[0-9]+}/skills/{skill_id:[0-9]+}", profileService.IsUsingSkill).Methods("GET")
	apiRouter.HandleFunc("/api/profile/{id:[0-9]+}/skills/{skill_id:[0-9]+}", profileService.PostSkill).Methods("POST")
	apiRouter.HandleFunc("/api/profile/{id:[0-9]+}/skills/{skill_id:[0-9]+}", profileService.DeleteSkill).Methods("DELETE")
	apiRouter.HandleFunc("/api/profile/{id:[0-9]+}/followed", profileService.GetFollowed).Methods("GET")
	apiRouter.HandleFunc("/api/profile/{id:[0-9]+}/follow/{profile_id:[0-9]+}", profileService.IsFollowingProfile).Methods("GET")
	apiRouter.HandleFunc("/api/profile/{id:[0-9]+}/follow/{profile_id:[0-9]+}", profileService.Follow).Methods("POST")
	apiRouter.HandleFunc("/api/profile/{id:[0-9]+}/follow/{profile_id:[0-9]+}", profileService.DeleteFollow).Methods("DELETE")
	apiRouter.HandleFunc("/api/profile/{id:[0-9]+}/jobs", profileService.GetJobs).Methods("GET")
	apiRouter.HandleFunc("/api/profile/{id:[0-9]+}/jobs/{job_id:[0-9]+}", profileService.IsLikingJob).Methods("GET")
	apiRouter.HandleFunc("/api/profile/{id:[0-9]+}/jobs/{job_id:[0-9]+}", profileService.PostJob).Methods("POST")
	apiRouter.HandleFunc("/api/profile/{id:[0-9]+}/jobs/{job_id:[0-9]+}", profileService.DeleteJob).Methods("DELETE")
	apiRouter.HandleFunc("/api/profile/{id:[0-9]+}/propositions/users", profileService.GetPropositionsUsers).Methods("GET")
	apiRouter.HandleFunc("/api/profile/{id:[0-9]+}/propositions/users/hobbies", profileService.GetPropositionsUsersHobbies).Methods("GET")
	apiRouter.HandleFunc("/api/profile/{id:[0-9]+}/propositions/companies", profileService.GetPropositionsCompanies).Methods("GET")
	apiRouter.HandleFunc("/api/profile/search/{pattern:.*?}", profileService.Search).Methods("GET")

	// companies
	companyService := api.NewCompanyService(companyManager)
	apiRouter.HandleFunc("/api/company/{id:[0-9]+}", companyService.Get).Methods("GET")
	apiRouter.HandleFunc("/api/company/search/{pattern:.*?}", companyService.Search).Methods("GET")
	apiRouter.HandleFunc("/api/company/{id:[0-9]+}/likers", companyService.GetLikers).Methods("GET")
	apiRouter.HandleFunc("/api/company/{id:[0-9]+}/jobs", companyService.GetJobs).Methods("GET")

	// hobbies
	hobbyService := api.NewHobbyService(hobbyManager)
	apiRouter.HandleFunc("/api/hobby/{id:[0-9]+}", hobbyService.Get).Methods("GET")
	apiRouter.HandleFunc("/api/hobby/search/{pattern:.*?}", hobbyService.Search).Methods("GET")
	apiRouter.HandleFunc("/api/hobby/{id:[0-9]+}/likers", hobbyService.GetLikers).Methods("GET")

	// jobs
	jobService := api.NewJobService(jobManager)
	apiRouter.HandleFunc("/api/job/{id:[0-9]+}", jobService.Get).Methods("GET")
	apiRouter.HandleFunc("/api/job/search/{pattern:.*?}", jobService.Search).Methods("GET")

	// skills
	skillService := api.NewSkillService(skillManager)
	apiRouter.HandleFunc("/api/skill/{id:[0-9]+}", skillService.Get).Methods("GET")
	apiRouter.HandleFunc("/api/skill/search/{pattern:.*?}", skillService.Search).Methods("GET")

	// common
	apiRouter.HandleFunc("/api/all/search/{pattern:.*?}", api.SearchAll(profileManager, companyManager, hobbyManager, jobManager, skillManager)).Methods("GET")

	router.PathPrefix("/api/").Handler(negroni.New(
		negronilogrus.NewMiddleware(),
		sessions.Sessions("session", store),
		internal.AuthMiddleware(profileManager),
		negroni.Wrap(apiRouter),
	))

	// static files
	router.PathPrefix("/").Handler(negroni.New(
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
