// File: hobby.ho
// File Created: 18 Mar 2019 14:26
// By Maxence Moutoussamy <maxence.moutoussamy1@gmail.com>

package api

import (
	"encoding/json"
	"net/http"
	"server/internal/dao"
	"strconv"

	"github.com/gorilla/mux"
)

type HobbyService struct {
	manager *dao.HobbyManager
}

func NewHobbyService(manager *dao.HobbyManager) *HobbyService {
	return &HobbyService{manager}
}

func (s *HobbyService) Get(rw http.ResponseWriter, req *http.Request) {
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
		internalServerError(rw, "cannot marshal hobby type", err)
		return
	}

	rw.Header().Set("Content-type", "application/json")
	rw.Write(datas)
}

func (s *HobbyService) Search(rw http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	pattern := vars["pattern"]

	hobbies, err := s.manager.Search(pattern)
	if err != nil {
		internalServerError(rw, "cannot Search Hobbies", err)
		return
	}

	datas, err := json.Marshal(hobbies)
	if err != nil {
		internalServerError(rw, "cannot marshal hobby type", err)
		return
	}

	rw.Header().Set("Content-type", "application/json")
	rw.Write(datas)
}

func (s *HobbyService) GetLikers(rw http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	hobbyID, _ := strconv.ParseInt(vars["id"], 10, 64)

	profiles, err := s.manager.GetLikers(hobbyID)
	if err != nil {
		internalServerError(rw, "cannot get likers", err)
		return
	}

	data, err := json.Marshal(profiles)
	if err != nil {
		internalServerError(rw, "cannont marshal profile slice type", err)
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.Write(data)
}
