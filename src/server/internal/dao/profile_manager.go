// File: profile_manager.go
// File Created: 11 Mar 2019 07:50
// By Maxence Moutoussamy <maxence.moutoussamy1@gmail.com>

package dao

import (
	"fmt"
	"io"
	"reflect"
	"strings"

	"github.com/johnnadratowski/golang-neo4j-bolt-driver/structures/graph"
	bolt "github.com/moutoum/golang-neo4j-bolt-driver"
)

type ProfileManager struct {
	pool bolt.DriverPool
}

func NewProfileManager(pool bolt.DriverPool) *ProfileManager {
	return &ProfileManager{pool}
}

func (m *ProfileManager) VerifyCreditentials(login, password string) (int64, error) {
	conn, err := m.pool.OpenPool()
	if err != nil {
		return -1, err
	}
	defer conn.Close()

	rows, err := conn.QueryNeo("MATCH (p:Profile) WHERE p.login = {login} AND p.password = {password} RETURN p LIMIT 1", map[string]interface{}{
		"login":    login,
		"password": password,
	})

	if err != nil {
		return -1, err
	}

	defer rows.Close()

	results, _, err := rows.All()
	if err != nil {
		return -1, err
	}

	if len(results) != 1 {
		return -1, nil
	}

	p, ok := results[0][0].(graph.Node)
	if !ok {
		return -1, fmt.Errorf("invalid type, expected `graph.Node` but got `%s`", reflect.TypeOf(results[0][0]))
	}

	return p.NodeIdentity, nil
}

func (m *ProfileManager) IsAccountExists(login string) (bool, error) {
	conn, err := m.pool.OpenPool()
	if err != nil {
		return false, err
	}
	defer conn.Close()

	rows, err := conn.QueryNeo("MATCH (p:Profile) WHERE p.login = {login} RETURN p LIMIT 1", map[string]interface{}{
		"login": login,
	})

	if err != nil {
		return false, err
	}

	defer rows.Close()

	results, _, err := rows.All()
	if err != nil {
		return false, err
	}

	return len(results) == 1, nil
}

func (b *ProfileManager) GetByID(id int64) (*Profile, error) {
	conn, err := b.pool.OpenPool()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	r, err := conn.QueryNeo(`
		MATCH (e:Education)<-[:Studies]-(p:Profile)-[:Lives]->(ci:City)-[:Located]->(co:Country)
		WHERE ID(p) = {id}
		RETURN p, ci, co, e
		LIMIT 1`,
		map[string]interface{}{
			"id": id,
		})

	if err != nil {
		return nil, fmt.Errorf("cannot query: %v", err)
	}

	defer r.Close()

	results, _, err := r.All()

	if err != nil {
		return nil, err
	}

	if len(results) == 0 {
		return nil, ErrNotFound
	}

	return rowToProfile(results[0]), nil
}

func (b *ProfileManager) Search(pattern string) ([]*Profile, error) {
	conn, err := b.pool.OpenPool()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	r, err := conn.QueryNeo(`
		MATCH (e:Education)<-[:Studies]-(p:Profile)-[:Lives]->(ci:City)-[:Located]->(co:Country) 
		WHERE 
			p.firstname CONTAINS {pattern} OR
			p.lastname CONTAINS {pattern} OR
			p.login CONTAINS {pattern}
		RETURN p, ci, co, e`,
		map[string]interface{}{
			"pattern": pattern,
		},
	)

	if err != nil {
		return nil, fmt.Errorf("cannot query: %v", err)
	}

	defer r.Close()

	var profiles []*Profile

	for data, _, err := r.NextNeo(); err != io.EOF; data, _, err = r.NextNeo() {
		profiles = append(profiles, rowToProfile(data))
	}

	return profiles, nil
}

func (m *ProfileManager) Create(p *Profile) error {
	conn, err := m.pool.OpenPool()
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.ExecNeo(`
		MERGE (co:Country{ name: {country_name} })
		MERGE (ci:City{ name: {city_name} })
		MERGE (e:Education{ level: {education_level} })
		CREATE (p:Profile{
				login: {login},
				password: {password},
				firstname: {fname},
				lastname: {lname},
				birthday: {birthday}
			})
		CREATE (ci)-[:Located]->(co)
		CREATE (p)-[:Lives]->(ci)`,
		map[string]interface{}{
			"country_name":    strings.Title(p.Country),
			"city_name":       strings.Title(p.City),
			"login":           p.Login,
			"password":        p.Password,
			"fname":           strings.Title(p.Firstname),
			"lname":           strings.Title(p.Lastname),
			"birthday":        p.Birthday,
			"education_level": p.EducationLevel,
		},
	)

	return err
}

func (m *ProfileManager) GetCompanies(profileID int) ([]*Company, error) {
	conn, err := m.pool.OpenPool()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	r, err := conn.QueryNeo(`
		MATCH (p:Profile)-[:Likes]->(c:Company)-[:Attached]->(d:Domain)
		WHERE ID(p) = {profileID}
		RETURN c, d.name`,
		map[string]interface{}{
			"profileID": profileID,
		})

	if err != nil {
		return nil, fmt.Errorf("cannot query: %v", err)
	}

	defer r.Close()

	companies := make(map[int64]*Company)
	for data, _, err := r.NextNeo(); err != io.EOF; data, _, err = r.NextNeo() {
		c, _ := data[0].(graph.Node)
		d, _ := data[1].(string)

		if _, ok := companies[c.NodeIdentity]; !ok {
			companies[c.NodeIdentity] = &Company{}
			companies[c.NodeIdentity].Entity.ID = c.NodeIdentity
			companies[c.NodeIdentity].Name = c.Properties["name"].(string)
			companies[c.NodeIdentity].Siret = c.Properties["siret"].(string)
			companies[c.NodeIdentity].Siren = c.Properties["siren"].(string)
			companies[c.NodeIdentity].Description = c.Properties["description"].(string)
		}
		companies[c.NodeIdentity].Domains = append(companies[c.NodeIdentity].Domains, d)
	}

	companiesOutput := make([]*Company, 0, len(companies))
	for _, company := range companies {
		companiesOutput = append(companiesOutput, company)
	}

	return companiesOutput, nil
}

func (m *ProfileManager) PostCompany(profileID, companyID int) error {
	conn, err := m.pool.OpenPool()
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.ExecNeo(`
		MATCH (p:Profile), (c:Company)
		WHERE ID(p) = {profileID} AND ID(c) = {companyID}
		MERGE (p)-[:Likes]->(c)`,
		map[string]interface{}{
			"profileID": profileID,
			"companyID": companyID,
		})

	if err != nil {
		return fmt.Errorf("cannot query: %v", err)
	}

	return nil
}

func (m *ProfileManager) GetHobbies(profileID int) ([]*Hobby, error) {
	conn, err := m.pool.OpenPool()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	r, err := conn.QueryNeo("MATCH (p:Profile)-[:Likes]->(s:Hobby) WHERE ID(p) = {profileID} RETURN s", map[string]interface{}{
		"profileID": profileID,
	})

	if err != nil {
		return nil, fmt.Errorf("cannot query: %v", err)
	}

	defer r.Close()

	hobbies := []*Hobby{}

	for data, _, err := r.NextNeo(); err != io.EOF; data, _, err = r.NextNeo() {
		d, ok := data[0].(graph.Node)

		if !ok {
			return nil, fmt.Errorf("invalid type, expected `graph.Node` but got `%s`", reflect.TypeOf(data[0]))
		}

		hobbies = append(hobbies, &Hobby{
			Entity: Entity{
				ID: d.NodeIdentity,
			},
			Name: d.Properties["name"].(string),
		})
	}

	return hobbies, nil
}

func (m *ProfileManager) PostHobby(profileID, hobbyID int) error {
	conn, err := m.pool.OpenPool()
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.ExecNeo(`
		MATCH (p:Profile), (c:Hobby)
		WHERE ID(p) = {profileID} AND ID(c) = {hobbyID}
		MERGE (p)-[:Likes]->(c)`,
		map[string]interface{}{
			"profileID": profileID,
			"hobbyID":   hobbyID,
		})

	if err != nil {
		return fmt.Errorf("cannot query: %v", err)
	}

	return nil
}

func (m *ProfileManager) GetSkills(profileID int) ([]*Skill, error) {
	conn, err := m.pool.OpenPool()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	r, err := conn.QueryNeo("MATCH (p:Profile)-[:Uses]->(s:Skill) WHERE ID(p) = {profileID} RETURN s", map[string]interface{}{
		"profileID": profileID,
	})

	if err != nil {
		return nil, fmt.Errorf("cannot query: %v", err)
	}

	defer r.Close()

	skills := []*Skill{}

	for data, _, err := r.NextNeo(); err != io.EOF; data, _, err = r.NextNeo() {
		d, ok := data[0].(graph.Node)

		if !ok {
			return nil, fmt.Errorf("invalid type, expected `graph.Node` but got `%s`", reflect.TypeOf(data[0]))
		}

		skills = append(skills, &Skill{
			Entity: Entity{
				ID: d.NodeIdentity,
			},
			Name: d.Properties["name"].(string),
		})
	}

	return skills, nil
}

func (m *ProfileManager) PostSkill(profileID, skillID int) error {
	conn, err := m.pool.OpenPool()
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.ExecNeo(`
		MATCH (p:Profile), (c:Skill)
		WHERE ID(p) = {profileID} AND ID(c) = {skillID}
		MERGE (p)-[:Uses]->(c)`,
		map[string]interface{}{
			"profileID": profileID,
			"skillID":   skillID,
		})

	if err != nil {
		return fmt.Errorf("cannot query: %v", err)
	}

	return nil
}

func (m *ProfileManager) GetFollowed(profileID int) ([]*Profile, error) {
	conn, err := m.pool.OpenPool()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	r, err := conn.QueryNeo(`
		MATCH
			(p:Profile)-[:Follow]->(f:Profile)-[:Lives]->(ci:City)-[:Located]->(co:Country),
			(f:Profile)-[:Studies]->(e:Education)
		WHERE ID(p) = {profileID}
		RETURN f, ci, co, e`,
		map[string]interface{}{
			"profileID": profileID,
		})

	if err != nil {
		return nil, fmt.Errorf("cannot query: %v", err)
	}

	defer r.Close()

	profiles := []*Profile{}

	for data, _, err := r.NextNeo(); err != io.EOF; data, _, err = r.NextNeo() {
		profiles = append(profiles, rowToProfile(data))
	}

	return profiles, nil
}

func (m *ProfileManager) Follow(profileID, profileIDToFollow int) error {
	conn, err := m.pool.OpenPool()
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.ExecNeo(`
		MATCH (p:Profile), (c:Profile)
		WHERE ID(p) = {profileID} AND ID(c) = {profileIDToFollow}
		MERGE (p)-[:Follow]->(c)`,
		map[string]interface{}{
			"profileID":         profileID,
			"profileIDToFollow": profileIDToFollow,
		})

	if err != nil {
		return fmt.Errorf("cannot query: %v", err)
	}

	return nil
}

func (m *ProfileManager) GetJobs(profileID int) ([]*Job, error) {
	conn, err := m.pool.OpenPool()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	r, err := conn.QueryNeo("MATCH (p:Profile)-[:Likes]->(j:Job) WHERE ID(p) = {profileID} RETURN j", map[string]interface{}{
		"profileID": profileID,
	})

	if err != nil {
		return nil, fmt.Errorf("cannot query: %v", err)
	}

	defer r.Close()

	jobs := []*Job{}

	for data, _, err := r.NextNeo(); err != io.EOF; data, _, err = r.NextNeo() {
		d, ok := data[0].(graph.Node)

		if !ok {
			return nil, fmt.Errorf("invalid type, expected `graph.Node` but got `%s`", reflect.TypeOf(data[0]))
		}

		jobs = append(jobs, &Job{
			Entity: Entity{
				ID: d.NodeIdentity,
			},
			Name:        d.Properties["name"].(string),
			GrossWage:   d.Properties["grossWage"].(string),
			Description: d.Properties["description"].(string),
		})
	}

	return jobs, nil
}

func (m *ProfileManager) PostJob(profileID, jobID int) error {
	conn, err := m.pool.OpenPool()
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.ExecNeo(`
		MATCH (p:Profile), (j:Job)
		WHERE ID(p) = {profileID} AND ID(j) = {jobID}
		MERGE (p)-[:Likes]->(j)`,
		map[string]interface{}{
			"profileID": profileID,
			"jobID":     jobID,
		})

	if err != nil {
		return fmt.Errorf("cannot query: %v", err)
	}

	return nil
}

func (m *ProfileManager) GetPropositionsUsers(profileID int) ([]*Profile, error) {
	conn, err := m.pool.OpenPool()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	results, err := conn.QueryNeo(`
		MATCH (me:Profile) WHERE ID(me) = {profileID} WITH me
		MATCH (e:Education)<-[:Studies]-(p:Profile)-[:Lives]->(ci:City)-[:Located]->(co:Country)
		WHERE NOT (me)-[:Follow]->(p) AND p <> me
		WITH me, p, ci, co, e
		MATCH (n)
		WHERE ((me)-->(n) AND (p)-->(n))
		RETURN DISTINCT p, ci, co, e`,
		map[string]interface{}{
			"profileID": profileID,
		})

	if err != nil {
		return nil, fmt.Errorf("cannot query: %v", err)
	}

	var profiles []*Profile
	for row, _, err := results.NextNeo(); err != io.EOF; row, _, err = results.NextNeo() {
		profiles = append(profiles, rowToProfile(row))
	}

	return profiles, nil
}

func rowToProfile(row []interface{}) *Profile {
	p, _ := row[0].(graph.Node)
	ci, _ := row[1].(graph.Node)
	co, _ := row[2].(graph.Node)
	e, _ := row[3].(graph.Node)

	return &Profile{
		Entity: Entity{
			ID: p.NodeIdentity,
		},
		Login:          p.Properties["login"].(string),
		Password:       p.Properties["password"].(string),
		Firstname:      p.Properties["firstname"].(string),
		Lastname:       p.Properties["lastname"].(string),
		Birthday:       p.Properties["birthday"].(string),
		Country:        co.Properties["name"].(string),
		City:           ci.Properties["name"].(string),
		EducationLevel: e.Properties["level"].(int64),
	}
}
