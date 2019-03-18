// File: job.go
// File Created: 18 Mar 2019 14:37
// By Maxence Moutoussamy <maxence.moutoussamy1@gmail.com>

package api

import (
	"encoding/json"
	"net/http"
	"server/internal/dao"
	"strconv"

	"github.com/gorilla/mux"
)

type JobService struct {
	manager *dao.JobManager
}

func NewJobService(manager *dao.JobManager) *JobService {
	return &JobService{manager}
}

func (s *JobService) Get(rw http.ResponseWriter, req *http.Request) {
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
		internalServerError(rw, "cannot marshal job type", err)
		return
	}

	rw.Header().Set("Content-type", "application/json")
	rw.Write(datas)
}
