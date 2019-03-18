package dao

import (
	"fmt"
	"reflect"

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
		MATCH (c:Company)
		WHERE ID(c) = {id}
		RETURN c
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

	c, ok := results[0][0].(graph.Node)

	if !ok {
		return nil, fmt.Errorf("invalid type, expected `graph.Node` but got `%s`", reflect.TypeOf(results[0][0]))
	}

	company := &Company{
		Entity: Entity{
			ID: c.NodeIdentity,
		},
		Name:        c.Properties["name"].(string),
		Siret:       c.Properties["siret"].(string),
		Siren:       c.Properties["siren"].(string),
		Description: c.Properties["description"].(string),
	}

	return company, nil
}

// func (b *DataManager) GetCompanyWithName(pattern string) error {
// 	var results []Company

// 	r, err := b.conn.QueryNeo("MATCH (c:Company) WHERE c.name CONTAINS {pattern} RETURN c", map[string]interface{}{
// 		"pattern": pattern,
// 	})
// 	if err != nil {
// 		return fmt.Errorf("cannot query: %v", err)
// 	}
// 	defer r.Close()

// 	for data, _, err := r.NextNeo(); err != io.EOF; data, _, err = r.NextNeo() {
// 		d := data[0].(graph.Node)

// 		p := Company{
// 			Name:        d.Properties["name"].(string),
// 			Siret:       d.Properties["siret"].(string),
// 			Siren:       d.Properties["siren"].(string),
// 			Description: d.Properties["description"].(string),
// 		}
// 		results = append(results, p)
// 	}

// 	for i := range results {
// 		fmt.Println(results[i].Name)
// 		fmt.Println(results[i].Siret)
// 		fmt.Println(results[i].Siren)
// 		fmt.Println(results[i].Description)
// 	}

// 	return nil
// }

// func (b *DataManager) SetCompany(name string, siret string, siren string, description string) error {
// 	_, err := b.conn.ExecNeo("CREATE (c:Company {name: {name}, siret: {siret}, siren: {siren} , description: {description}})", map[string]interface{}{
// 		"name":        name,
// 		"siret":       siret,
// 		"siren":       siren,
// 		"description": description,
// 	})
// 	if err != nil {
// 		return fmt.Errorf("cannot query: %v", err)
// 	}

// 	return nil
// }
