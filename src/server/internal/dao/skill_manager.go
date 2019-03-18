// File: skill_manager.go
// File Created: 18 Mar 2019 14:47
// By Maxence Moutoussamy <maxence.moutoussamy1@gmail.com>

package dao

import (
	"fmt"
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
