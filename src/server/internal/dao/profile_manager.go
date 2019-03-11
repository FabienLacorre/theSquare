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

type Profile struct {
	Login     string
	Password  string
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Birthday  string `json:"birthday"`
	Country   string `json:"country"`
	City      string `json:"city"`
}

type ProfileManager struct {
	conn bolt.Conn
}

func NewProfileManager(conn bolt.Conn) *ProfileManager {
	return &ProfileManager{conn}
}

func (m *ProfileManager) VerifyCreditentials(login, password string) (bool, error) {
	rows, err := m.conn.QueryNeo("MATCH (p:Profile) WHERE p.login = {login} AND p.password = {password} RETURN p LIMIT 1", map[string]interface{}{
		"login":    login,
		"password": password,
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

func (m *ProfileManager) IsAccountExists(login string) (bool, error) {
	rows, err := m.conn.QueryNeo("MATCH (p:Profile) WHERE p.login = {login} RETURN p LIMIT 1", map[string]interface{}{
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

func (b *ProfileManager) GetProfileWithID(id int) (*Profile, error) {
	r, err := b.conn.QueryNeo(`
		MATCH (p:Profile)-[:Lives]->(ci:City)-[:Located]->(co:Country)
		WHERE ID(p) = {id}
		RETURN p, ci, co
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
		return nil, fmt.Errorf("profile with id = %d not found", id)
	}

	p, ok := results[0][0].(graph.Node)

	if !ok {
		return nil, fmt.Errorf("invalid type, expected `graph.Node` but got `%s`", reflect.TypeOf(results[0][0]))
	}

	ci, ok := results[0][1].(graph.Node)
	if !ok {
		return nil, fmt.Errorf("invalid type, expected `graph.Node` but got `%s`", reflect.TypeOf(results[0][0]))
	}

	co, ok := results[0][2].(graph.Node)
	if !ok {
		return nil, fmt.Errorf("invalid type, expected `graph.Node` but got `%s`", reflect.TypeOf(results[0][0]))
	}

	profile := &Profile{
		Login:     p.Properties["login"].(string),
		Password:  p.Properties["password"].(string),
		Firstname: p.Properties["firstname"].(string),
		Lastname:  p.Properties["lastname"].(string),
		Birthday:  p.Properties["birthday"].(string),
		Country:   co.Properties["name"].(string),
		City:      ci.Properties["name"].(string),
	}

	return profile, nil
}

func (b *ProfileManager) Search(pattern string) ([]*Profile, error) {
	r, err := b.conn.QueryNeo(`
		MATCH (p:Profile)-[:Lives]->(ci:City)-[:Located]->(co:Country) 
		WHERE 
			p.firstname CONTAINS {pattern} OR
			p.lastname CONTAINS {pattern} OR
			p.login CONTAINS {pattern}
		RETURN p, ci, co`,
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
		d, ok := data[0].(graph.Node)

		if !ok {
			return nil, fmt.Errorf("invalid type, expected `graph.Node` but got `%s`", reflect.TypeOf(data[0]))
		}

		ci, ok := data[1].(graph.Node)
		if !ok {
			return nil, fmt.Errorf("invalid type, expected `graph.Node` but got `%s`", reflect.TypeOf(data[1]))
		}

		co, ok := data[2].(graph.Node)
		if !ok {
			return nil, fmt.Errorf("invalid type, expected `graph.Node` but got `%s`", reflect.TypeOf(data[2]))
		}

		profiles = append(profiles, &Profile{
			Login:     d.Properties["login"].(string),
			Password:  d.Properties["password"].(string),
			Firstname: d.Properties["firstname"].(string),
			Lastname:  d.Properties["lastname"].(string),
			Birthday:  d.Properties["birthday"].(string),
			Country:   co.Properties["name"].(string),
			City:      ci.Properties["name"].(string),
		})
	}

	return profiles, nil
}

func (m *ProfileManager) Create(p *Profile) error {
	_, err := m.conn.ExecNeo(`
		MERGE (co:Country{ name: {country_name} })
		MERGE (ci:City{ name: {city_name} })
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
			"country_name": strings.Title(p.Country),
			"city_name":    strings.Title(p.City),
			"login":        p.Login,
			"password":     p.Password,
			"fname":        strings.Title(p.Firstname),
			"lname":        strings.Title(p.Lastname),
			"birthday":     p.Birthday,
		},
	)

	return err
}
