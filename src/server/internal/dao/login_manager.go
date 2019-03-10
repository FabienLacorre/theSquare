// File: login_manager.go
// File Created: 09 Mar 2019 12:04
// By Maxence Moutoussamy <maxence.moutoussamy1@gmail.com>

package dao

import (
	"strings"

	bolt "github.com/moutoum/golang-neo4j-bolt-driver"
)

type LoginManager struct {
	connection bolt.Conn
}

func NewLoginManager(connection bolt.Conn) *LoginManager {
	return &LoginManager{connection}
}

func (m *LoginManager) VerifyCreditentials(login, password string) (bool, error) {
	rows, err := m.connection.QueryNeo("MATCH (p:Profile) WHERE p.login = {login} AND p.password = {password} RETURN p LIMIT 1", map[string]interface{}{
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

func (m *LoginManager) IsAccountExists(login string) (bool, error) {
	rows, err := m.connection.QueryNeo("MATCH (p:Profile) WHERE p.login = {login} RETURN p LIMIT 1", map[string]interface{}{
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

func (m *LoginManager) CreateAccount(login, password, name, surname, brithdate, country, city string) error {
	_, err := m.connection.ExecNeo(`
		MERGE (co:Country{ name: {country_name} })
		MERGE (ci:City{ name: {city_name} })
		CREATE (p:Profile{
				login: {login},
				password: {password},
				name: {name},
				surname: {surname},
				birthdate: {birthdate}
			})
		CREATE (p)-[:Lives]->(co)
		CREATE (p)-[:Lives]->(ci)`,
		map[string]interface{}{
			"country_name": strings.Title(country),
			"city_name":    strings.Title(city),
			"login":        login,
			"password":     password,
			"name":         strings.Title(name),
			"surname":      strings.Title(surname),
			"birthdate":    brithdate,
		},
	)

	return err
}
