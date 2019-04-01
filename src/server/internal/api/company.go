// File: company.go
// File Created: 18 Mar 2019 14:07
// By Maxence Moutoussamy <maxence.moutoussamy1@gmail.com>

package api

import (
	"encoding/json"
	"net/http"
	"server/internal/dao"
	"strconv"

	"github.com/gorilla/mux"
)

type CompanyService struct {
	manager *dao.CompanyManager
}

func NewCompanyService(manager *dao.CompanyManager) *CompanyService {
	return &CompanyService{manager}
}

func (s *CompanyService) Get(rw http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id, _ := strconv.ParseInt(vars["id"], 10, 64)

	company, err := s.manager.GetByID(id)
	if err != nil {
		if err == dao.ErrNotFound {
			http.NotFound(rw, req)
			return
		}
		internalServerError(rw, "cannot GetByID", err)
		return
	}

	datas, err := json.Marshal(company)
	if err != nil {
		internalServerError(rw, "cannot marshal company type", err)
		return
	}

	rw.Header().Set("Content-type", "application/json")
	rw.Write(datas)
}

func (s *CompanyService) Search(rw http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	pattern := vars["pattern"]

	companies, err := s.manager.Search(pattern)
	if err != nil {
		internalServerError(rw, "cannot Search Company", err)
		return
	}

	datas, err := json.Marshal(companies)
	if err != nil {
		internalServerError(rw, "cannot marshal company type", err)
		return
	}

	rw.Header().Set("Content-type", "application/json")
	rw.Write(datas)
}

func (s *CompanyService) GetLikers(rw http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	companyID, _ := strconv.ParseInt(vars["id"], 10, 64)

	profiles, err := s.manager.GetLikers(companyID)
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
