package dao

import (
	"fmt"
	"io"

	"github.com/johnnadratowski/golang-neo4j-bolt-driver/structures/graph"

	bolt "github.com/moutoum/golang-neo4j-bolt-driver"
)

type CompanyManager struct {
	pool bolt.DriverPool
}

func NewCompanyManager(pool bolt.DriverPool) *CompanyManager {
	return &CompanyManager{pool}
}

func (m *CompanyManager) GetByID(id int64) (*Company, error) {
	conn, err := m.pool.OpenPool()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	r, err := conn.QueryNeo(`
		MATCH (c:Company)-[:Attached]->(d:Domain)
		WHERE ID(c) = {id}
		RETURN c, d.name`,
		map[string]interface{}{
			"id": id,
		})

	if err != nil {
		return nil, fmt.Errorf("cannot query: %v", err)
	}

	defer r.Close()

	i := 0
	company := &Company{}
	for data, _, err := r.NextNeo(); err != io.EOF; data, _, err = r.NextNeo() {
		if i == 0 {
			c, _ := data[0].(graph.Node)
			company.Entity.ID = c.NodeIdentity
			company.Name = c.Properties["name"].(string)
			company.Siret = c.Properties["siret"].(string)
			company.Siren = c.Properties["siren"].(string)
			company.Description = c.Properties["description"].(string)
		}

		d, _ := data[1].(string)
		company.Domains = append(company.Domains, d)
		i++
	}

	return company, nil
}
