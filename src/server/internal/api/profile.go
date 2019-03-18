// File: profile.go
// File Created: 07 Mar 2019 16:10
// By Maxence Moutoussamy <maxence.moutoussamy1@gmail.com>

package api

import (
	"encoding/json"
	"net/http"
	"server/internal/dao"
	"strconv"

	"github.com/gorilla/mux"
)

type ProfileService struct {
	manager *dao.ProfileManager
}

func NewProfileService(manager *dao.ProfileManager) *ProfileService {
	return &ProfileService{manager}
}

func (s *ProfileService) Get(rw http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id, _ := strconv.ParseInt(vars["id"], 10, 64)

	profile, err := s.manager.GetByID(id)
	if err != nil {
		if err == dao.ErrNotFound {
			http.NotFound(rw, req)
			return
		}
		internalServerError(rw, "cannot GetByID", err)
		return
	}

	datas, err := json.Marshal(profile)
	if err != nil {
		internalServerError(rw, "cannot marshal profile type", err)
		return
	}

	rw.Header().Set("Content-type", "application/json")
	rw.Write(datas)
}

func (s *ProfileService) GetCompanies(rw http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id, _ := strconv.Atoi(vars["id"])

	daoCompanies, err := s.manager.GetCompanies(id)
	if err != nil {
		internalServerError(rw, "cannot GetCompanies", err)
		return
	}

	datas, err := json.Marshal(daoCompanies)
	if err != nil {
		internalServerError(rw, "cannot marshal Company type", err)
		return
	}

	rw.Header().Set("Content-type", "application/json")
	rw.Write(datas)
}

func (s *ProfileService) GetHobbies(rw http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id, _ := strconv.Atoi(vars["id"])

	daoHobbies, err := s.manager.GetHobbies(id)
	if err != nil {
		internalServerError(rw, "cannot GetHobbies", err)
		return
	}

	datas, err := json.Marshal(daoHobbies)
	if err != nil {
		internalServerError(rw, "cannot marshal Skill type", err)
		return
	}

	rw.Header().Set("Content-type", "application/json")
	rw.Write(datas)
}

func (s *ProfileService) GetSkills(rw http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id, _ := strconv.Atoi(vars["id"])

	skills, err := s.manager.GetSkills(id)
	if err != nil {
		internalServerError(rw, "cannot GetSkills", err)
		return
	}

	data, err := json.Marshal(skills)
	if err != nil {
		internalServerError(rw, "cannot marshal Skill type", err)
		return
	}

	rw.Header().Set("Content-type", "application/json")
	rw.Write(data)
}

func (s *ProfileService) GetFollowed(rw http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id, _ := strconv.Atoi(vars["id"])

	followed, err := s.manager.GetFollowed(id)
	if err != nil {
		internalServerError(rw, "cannot GetFollowed", err)
		return
	}

	data, err := json.Marshal(followed)
	if err != nil {
		internalServerError(rw, "cannot marshal Skill type", err)
		return
	}

	rw.Header().Set("Content-type", "application/json")
	rw.Write(data)
}

func (s *ProfileService) GetJobs(rw http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id, _ := strconv.Atoi(vars["id"])

	jobs, err := s.manager.GetJobs(id)
	if err != nil {
		internalServerError(rw, "cannot GetJobs", err)
		return
	}

	data, err := json.Marshal(jobs)
	if err != nil {
		internalServerError(rw, "cannot marshal Skill type", err)
		return
	}

	rw.Header().Set("Content-type", "application/json")
	rw.Write(data)
}
