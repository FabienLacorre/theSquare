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

func (s *ProfileService) PostCompany(rw http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id, _ := strconv.Atoi(vars["id"])
	companyID, _ := strconv.Atoi(vars["company_id"])

	if err := s.manager.PostCompany(id, companyID); err != nil {
		internalServerError(rw, "cannot PostCompany", err)
		return
	}

	rw.WriteHeader(http.StatusOK)
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

func (s *ProfileService) PostHobby(rw http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id, _ := strconv.Atoi(vars["id"])
	hobbyID, _ := strconv.Atoi(vars["hobby_id"])

	if err := s.manager.PostHobby(id, hobbyID); err != nil {
		internalServerError(rw, "cannot PostHobby", err)
		return
	}

	rw.WriteHeader(http.StatusOK)
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

func (s *ProfileService) PostSkill(rw http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id, _ := strconv.Atoi(vars["id"])
	skillID, _ := strconv.Atoi(vars["skill_id"])

	if err := s.manager.PostSkill(id, skillID); err != nil {
		internalServerError(rw, "cannot PostSkill", err)
		return
	}

	rw.WriteHeader(http.StatusOK)
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

func (s *ProfileService) Follow(rw http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id, _ := strconv.Atoi(vars["id"])
	profileID, _ := strconv.Atoi(vars["profile_id"])

	if err := s.manager.Follow(id, profileID); err != nil {
		internalServerError(rw, "cannot Follow", err)
		return
	}

	rw.WriteHeader(http.StatusOK)
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

func (s *ProfileService) PostJob(rw http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id, _ := strconv.Atoi(vars["id"])
	jobID, _ := strconv.Atoi(vars["job_id"])

	if err := s.manager.PostJob(id, jobID); err != nil {
		internalServerError(rw, "cannot PostJob", err)
		return
	}

	rw.WriteHeader(http.StatusOK)
}

func (s *ProfileService) GetPropositionsUsers(rw http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	profileID, _ := strconv.Atoi(vars["id"])

	profiles, err := s.manager.GetPropositionsUsers(profileID)
	if err != nil {
		internalServerError(rw, "cannot GetPropositionsUsers", err)
		return
	}

	data, err := json.Marshal(profiles)
	if err != nil {
		internalServerError(rw, "cannot marshal Profile type", err)
		return
	}

	rw.Header().Set("Content-type", "application/json")
	rw.Write(data)
}

func (s *ProfileService) GetPropositionsCompanies(rw http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	profileID, _ := strconv.Atoi(vars["id"])

	companies, err := s.manager.GetPropositionsCompanies(profileID)
	if err != nil {
		internalServerError(rw, "cannot GetPropositionsCompanies", err)
		return
	}

	data, err := json.Marshal(companies)
	if err != nil {
		internalServerError(rw, "cannot marshal Company type", err)
		return
	}

	rw.Header().Set("Content-type", "application/json")
	rw.Write(data)
}

func (s *ProfileService) GetPropositionsUsersHobbies(rw http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	profileID, _ := strconv.Atoi(vars["id"])

	hobbies, err := s.manager.GetPropositionsUsersHobbies(profileID)
	if err != nil {
		internalServerError(rw, "cannot GetPropositionsUsersHobbies", err)
		return
	}

	data, err := json.Marshal(hobbies)
	if err != nil {
		internalServerError(rw, "cannot marshal Hobby type", err)
		return
	}

	rw.Header().Set("Content-type", "application/json")
	rw.Write(data)
}

func (s *ProfileService) DeleteHobby(rw http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id, _ := strconv.Atoi(vars["id"])
	idUnliked, _ := strconv.Atoi(vars["hobby_id"])

	if err := s.manager.DeleteHobby(id, idUnliked); err != nil {
		internalServerError(rw, "cannot delete", err)
		return
	}
}

func (s *ProfileService) DeleteCompany(rw http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id, _ := strconv.Atoi(vars["id"])
	idUnliked, _ := strconv.Atoi(vars["company_id"])

	if err := s.manager.DeleteCompany(id, idUnliked); err != nil {
		internalServerError(rw, "cannot delete", err)
		return
	}
}

func (s *ProfileService) DeleteFollow(rw http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id, _ := strconv.Atoi(vars["id"])
	idUnliked, _ := strconv.Atoi(vars["profile_id"])

	if err := s.manager.DeleteFollow(id, idUnliked); err != nil {
		internalServerError(rw, "cannot delete", err)
		return
	}
}

func (s *ProfileService) DeleteSkill(rw http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id, _ := strconv.Atoi(vars["id"])
	idUnliked, _ := strconv.Atoi(vars["skill_id"])

	if err := s.manager.DeleteSkill(id, idUnliked); err != nil {
		internalServerError(rw, "cannot delete", err)
		return
	}
}

func (s *ProfileService) IsLikingCompany(rw http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	profileID, _ := strconv.Atoi(vars["id"])
	companyID, _ := strconv.Atoi(vars["company_id"])

	isLiking, err := s.manager.IsLikingCompany(profileID, companyID)
	if err != nil {
		internalServerError(rw, "cannot get company", err)
		return
	}

	rw.Header().Set("Content-type", "application/json")
	rw.Write([]byte(strconv.FormatBool(isLiking)))
}

func (s *ProfileService) IsLikingHobby(rw http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	profileID, _ := strconv.Atoi(vars["id"])
	hobbyID, _ := strconv.Atoi(vars["hobby_id"])

	isLiking, err := s.manager.IsLikingHobby(profileID, hobbyID)
	if err != nil {
		internalServerError(rw, "cannot get company", err)
		return
	}

	rw.Header().Set("Content-type", "application/json")
	rw.Write([]byte(strconv.FormatBool(isLiking)))
}

func (s *ProfileService) IsUsingSkill(rw http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	profileID, _ := strconv.Atoi(vars["id"])
	skillID, _ := strconv.Atoi(vars["skill_id"])

	isUsing, err := s.manager.IsUsingSkill(profileID, skillID)
	if err != nil {
		internalServerError(rw, "cannot get skill", err)
		return
	}

	rw.Header().Set("Content-type", "application/json")
	rw.Write([]byte(strconv.FormatBool(isUsing)))
}

func (s *ProfileService) IsFollowingProfile(rw http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	profileID, _ := strconv.Atoi(vars["id"])
	otherProfileID, _ := strconv.Atoi(vars["profile_id"])

	isFollowing, err := s.manager.IsFollowingProfile(profileID, otherProfileID)
	if err != nil {
		internalServerError(rw, "cannot get profile", err)
		return
	}

	rw.Header().Set("Content-type", "application/json")
	rw.Write([]byte(strconv.FormatBool(isFollowing)))
}

func (s *ProfileService) IsLikingJob(rw http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	profileID, _ := strconv.Atoi(vars["id"])
	jobID, _ := strconv.Atoi(vars["job_id"])

	isLiking, err := s.manager.IsLikingJob(profileID, jobID)
	if err != nil {
		internalServerError(rw, "cannot get job", err)
		return
	}

	rw.Header().Set("Content-type", "application/json")
	rw.Write([]byte(strconv.FormatBool(isLiking)))
}

func (s *ProfileService) DeleteJob(rw http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id, _ := strconv.Atoi(vars["id"])
	idUnliked, _ := strconv.Atoi(vars["job_id"])

	if err := s.manager.DeleteJob(id, idUnliked); err != nil {
		internalServerError(rw, "cannot delete", err)
		return
	}
}

func (s *ProfileService) Search(rw http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	pattern := vars["pattern"]

	profiles, err := s.manager.Search(pattern)
	if err != nil {
		internalServerError(rw, "cannot Search Profile", err)
		return
	}

	datas, err := json.Marshal(profiles)
	if err != nil {
		internalServerError(rw, "cannot marshal profile type", err)
		return
	}

	rw.Header().Set("Content-type", "application/json")
	rw.Write(datas)
}
