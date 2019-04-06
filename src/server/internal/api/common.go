// File: common.go
// File Created: 15 Mar 2019 07:36
// By Maxence Moutoussamy <maxence.moutoussamy1@gmail.com>

package api

import (
	"encoding/json"
	"net/http"
	"server/internal/dao"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

func internalServerError(rw http.ResponseWriter, message string, err error) {
	if err != nil {
		logrus.WithError(err).Error(message)
	}
	http.Error(rw, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func SearchAll(profileManager *dao.ProfileManager, companyManager *dao.CompanyManager, hobbyManager *dao.HobbyManager, jobManager *dao.JobManager, skillManager *dao.SkillManager) func(rw http.ResponseWriter, req *http.Request) {
	return func(rw http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		pattern := vars["pattern"]

		funcs := []func(string) (*dao.SearchResponse, error){
			profileManager.Search,
			companyManager.Search,
			hobbyManager.Search,
			jobManager.Search,
			skillManager.Search,
		}

		response := dao.NewSearchResponse()

		for _, f := range funcs {
			r, err := f(pattern)
			if err != nil {
				internalServerError(rw, "cannot search", err)
				return
			}

			response.Skills = append(response.Skills, r.Skills...)
			response.Companies = append(response.Companies, r.Companies...)
			response.Profiles = append(response.Profiles, r.Profiles...)
			response.Jobs = append(response.Jobs, r.Jobs...)
			response.Hobbies = append(response.Hobbies, r.Hobbies...)
		}

		datas, err := json.Marshal(response)
		if err != nil {
			internalServerError(rw, "cannot marshal SearchResponse type", err)
			return
		}

		rw.Header().Set("Content-type", "application/json")
		rw.Write(datas)
	}
}
