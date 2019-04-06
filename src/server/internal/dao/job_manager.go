package dao

import (
	"fmt"
	"io"
	"reflect"

	"github.com/johnnadratowski/golang-neo4j-bolt-driver/structures/graph"
	bolt "github.com/moutoum/golang-neo4j-bolt-driver"
)

type JobManager struct {
	pool bolt.DriverPool
}

func NewJobManager(pool bolt.DriverPool) *JobManager {
	return &JobManager{pool}
}

func (m *JobManager) GetByID(id int64) (*Job, error) {
	conn, err := m.pool.OpenPool()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	r, err := conn.QueryNeo(`
		MATCH (h:Job)
		WHERE ID(h) = {id}
		RETURN h
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

	j, ok := results[0][0].(graph.Node)

	if !ok {
		return nil, fmt.Errorf("invalid type, expected `graph.Node` but got `%s`", reflect.TypeOf(results[0][0]))
	}

	job := &Job{
		Entity: Entity{
			ID: j.NodeIdentity,
		},
		Name:        j.Properties["name"].(string),
		GrossWage:   j.Properties["grossWage"].(string),
		Description: j.Properties["description"].(string),
	}

	return job, nil
}

func (m *JobManager) Search(pattern string) (*SearchResponse, error) {
	conn, err := m.pool.OpenPool()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	r, err := conn.QueryNeo(`
		MATCH (j:Job)
		WHERE
			j.name CONTAINS {pattern} OR
			j.description CONTAINS {pattern}
		RETURN j`,
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

		response.Jobs = append(response.Jobs, &Job{
			Entity: Entity{
				ID: d.NodeIdentity,
			},
			Name:        d.Properties["name"].(string),
			GrossWage:   d.Properties["grossWage"].(string),
			Description: d.Properties["description"].(string),
		})
	}

	return response, nil
}

// func (b *DataManager) GetJobWithName(pattern string) error {
// 	var results []Job

// 	r, err := b.conn.QueryNeo("MATCH (j:Job) WHERE j.name CONTAINS {pattern} RETURN j", map[string]interface{}{
// 		"pattern": pattern,
// 	})
// 	if err != nil {
// 		return fmt.Errorf("cannot query: %v", err)
// 	}
// 	defer r.Close()

// 	for data, _, err := r.NextNeo(); err != io.EOF; data, _, err = r.NextNeo() {
// 		d := data[0].(graph.Node)

// 		p := Job{
// 			Name:        d.Properties["name"].(string),
// 			Description: d.Properties["description"].(string),
// 			GrossWage:   d.Properties["grossWage"].(string),
// 		}
// 		results = append(results, p)
// 	}

// 	for i := range results {
// 		fmt.Println(results[i].Name)
// 		fmt.Println(results[i].Description)
// 		fmt.Println(results[i].GrossWage)
// 	}

// 	return nil
// }

// func (b *DataManager) SetJob(name string, description string, grossWage string) error {
// 	_, err := b.conn.ExecNeo("CREATE (j:Job {name: {name}, description: {description}, grossWage: {grossWage}})", map[string]interface{}{
// 		"name":        name,
// 		"description": description,
// 		"grossWage":   grossWage,
// 	})
// 	if err != nil {
// 		return fmt.Errorf("cannot query: %v", err)
// 	}

// 	return nil
// }
