// File: structs.go
// File Created: 16 Mar 2019 17:20
// By Maxence Moutoussamy <maxence.moutoussamy1@gmail.com>

package dao

type Entity struct {
	ID int64 `json:"id"`
}

type Company struct {
	Entity
	Name        string   `json:"name"`
	Siret       string   `json:"siret"`
	Siren       string   `json:"siren"`
	Description string   `json:"description"`
	Domains     []string `json:"domains"`
}

type Job struct {
	Entity
	Name        string `json:"name"`
	Description string `json:"description"`
	GrossWage   string `json:"grossWage"`
}

type Profile struct {
	Entity
	Login          string `json:"login"`
	Password       string `json:"-"`
	Firstname      string `json:"firstname"`
	Lastname       string `json:"lastname"`
	Birthday       string `json:"birthday"`
	Country        string `json:"country"`
	City           string `json:"city"`
	EducationLevel int64  `json:"education"`
}

type Skill struct {
	Entity
	Name string `json:"name"`
}

type Hobby struct {
	Entity
	Name string `json:"name"`
}

type SearchResponse struct {
	Companies []*Company `json:"companies"`
	Jobs      []*Job     `json:"jobs"`
	Profiles  []*Profile `json:"profiles"`
	Skills    []*Skill   `json:"skills"`
	Hobbies   []*Hobby   `json:"hobbies"`
}

func NewSearchResponse() *SearchResponse {
	return &SearchResponse{
		Companies: []*Company{},
		Jobs:      []*Job{},
		Profiles:  []*Profile{},
		Skills:    []*Skill{},
		Hobbies:   []*Hobby{},
	}
}
