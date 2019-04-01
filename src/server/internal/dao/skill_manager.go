// File: skill_manager.go
// File Created: 18 Mar 2019 14:47
// By Maxence Moutoussamy <maxence.moutoussamy1@gmail.com>

package dao

import (
	"fmt"
	"io"
	"reflect"

	"github.com/johnnadratowski/golang-neo4j-bolt-driver/structures/graph"
	bolt "github.com/moutoum/golang-neo4j-bolt-driver"
)

type SkillManager struct {
	pool bolt.DriverPool
}

func NewSkillManager(pool bolt.DriverPool) *SkillManager {
	return &SkillManager{pool}
}

func (m *SkillManager) GetByID(id int64) (*Skill, error) {
	conn, err := m.pool.OpenPool()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	r, err := conn.QueryNeo(`
		MATCH (s:Skill)
		WHERE ID(s) = {id}
		RETURN s
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

	h, ok := results[0][0].(graph.Node)

	if !ok {
		return nil, fmt.Errorf("invalid type, expected `graph.Node` but got `%s`", reflect.TypeOf(results[0][0]))
	}

	skill := &Skill{
		Entity: Entity{
			ID: h.NodeIdentity,
		},
		Name: h.Properties["name"].(string),
	}

	return skill, nil
}

func (m *SkillManager) Search(pattern string) (*SearchResponse, error) {
	conn, err := m.pool.OpenPool()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	r, err := conn.QueryNeo(`
		MATCH (s:Skill)
		WHERE s.name CONTAINS {pattern}
		RETURN s`,
		map[string]interface{}{
			"pattern": pattern,
		})

	if err != nil {
		return nil, fmt.Errorf("cannot query: %v", err)
	}

	defer r.Close()

	response := NewSearchResponse()

	if err != nil {
		return nil, err
	}

	for data, _, err := r.NextNeo(); err != io.EOF; data, _, err = r.NextNeo() {
		d, ok := data[0].(graph.Node)

		if !ok {
			return nil, fmt.Errorf("invalid type, expected `graph.Node` but got `%s`", reflect.TypeOf(data[0]))
		}

		response.Skills = append(response.Skills, &Skill{
			Entity: Entity{
				ID: d.NodeIdentity,
			},
			Name: d.Properties["name"].(string),
		})
	}

	return response, nil
}
