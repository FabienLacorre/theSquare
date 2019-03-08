package dao

import (
	"fmt"
	"io"
	//"time"

	//"github.com/davecgh/go-spew/spew"
	"github.com/johnnadratowski/golang-neo4j-bolt-driver/structures/graph"
	bolt "github.com/moutoum/golang-neo4j-bolt-driver"
)

type Profile struct {
	firstName string `json:"firstName"`
	lastName string `json:"lastName"`
	birthday string `json:"birthday"`
	residence string `json:"residence"`
}

type Company struct {
	name string `json:"name"`
	siret string `json:"siret"`
	siren string `json:"siren"`
	description string `json:"description"`
}

type Job struct {
	name string `json:"name"`
	description string `json:"description"`
	grossWage string `json:"grossWage"`
}

type DataManager struct {
	conn bolt.Conn
}

func NewDataManager(conn bolt.Conn) *DataManager {
	return &DataManager{
		conn: conn,
	}
}

func (b *DataManager) GetProfileWithID(id int) error {
	var results []Profile

	r, err := b.conn.QueryNeo("MATCH (p:Person) WHERE ID(p) = {id} RETURN p", map[string]interface{}{
		"id": id,
	})
	if err != nil {
		return fmt.Errorf("cannot query: %v", err)
	}
	defer r.Close()

	for data, _, err := r.NextNeo(); err != io.EOF; data, _, err = r.NextNeo() {
		d := data[0].(graph.Node)

		p := Profile {
	        firstName:    d.Properties["firstName"].(string),
	        lastName:  d.Properties["lastName"].(string),
	        birthday:  d.Properties["birthday"].(string),
	        residence:  d.Properties["residence"].(string), 
	    }
	    results = append(results, p)
	}

	for i := range results {
		fmt.Println(results[i].firstName)
		fmt.Println(results[i].lastName)
		fmt.Println(results[i].birthday)
		fmt.Println(results[i].residence)	
	}

	return nil
}

func (b *DataManager) GetProfileWithName(pattern string) error {
	var results []Profile

	r, err := b.conn.QueryNeo("MATCH (p:Person) WHERE p.firstName CONTAINS {pattern} RETURN p", map[string]interface{}{
		"pattern": pattern,
	})
	if err != nil {
		return fmt.Errorf("cannot query: %v", err)
	}
	defer r.Close()

	for data, _, err := r.NextNeo(); err != io.EOF; data, _, err = r.NextNeo() {
		d := data[0].(graph.Node)

		p := Profile {
	        firstName:    d.Properties["firstName"].(string),
	        lastName:  d.Properties["lastName"].(string),
	        birthday:  d.Properties["birthday"].(string),
	        residence:  d.Properties["residence"].(string), 
	    }
	   	results = append(results, p)
	}

	for i := range results {
		fmt.Println(results[i].firstName)
		fmt.Println(results[i].lastName)
		fmt.Println(results[i].birthday)
		fmt.Println(results[i].residence)	
	}

	return nil
}

func (b *DataManager) GetProfileWithLastName(pattern string) error {
	var results []Profile

	r, err := b.conn.QueryNeo("MATCH (p:Person) WHERE p.lastName CONTAINS {pattern} RETURN p", map[string]interface{}{
		"pattern": pattern,
	})
	if err != nil {
		return fmt.Errorf("cannot query: %v", err)
	}
	defer r.Close()

	for data, _, err := r.NextNeo(); err != io.EOF; data, _, err = r.NextNeo() {
		d := data[0].(graph.Node)

		p := Profile {
	        firstName:    d.Properties["firstName"].(string),
	        lastName:  d.Properties["lastName"].(string),
	        birthday:  d.Properties["birthday"].(string),
	        residence:  d.Properties["residence"].(string), 
	    }
	   	results = append(results, p)
	}

	for i := range results {
		fmt.Println(results[i].firstName)
		fmt.Println(results[i].lastName)
		fmt.Println(results[i].birthday)
		fmt.Println(results[i].residence)	
	}

	return nil
}

func (b *DataManager) GetCompanyWithName(pattern string) error {
	var results []Company

	r, err := b.conn.QueryNeo("MATCH (c:Company) WHERE c.name CONTAINS {pattern} RETURN c", map[string]interface{}{
		"pattern": pattern,
	})
	if err != nil {
		return fmt.Errorf("cannot query: %v", err)
	}
	defer r.Close()

	for data, _, err := r.NextNeo(); err != io.EOF; data, _, err = r.NextNeo() {
		d := data[0].(graph.Node)

		p := Company {
	        name:    d.Properties["name"].(string),
	        siret:  d.Properties["siret"].(string),
	        siren:  d.Properties["siren"].(string),
	        description:  d.Properties["description"].(string), 
	    }
	   	results = append(results, p)
	}

	for i := range results {
		fmt.Println(results[i].name)
		fmt.Println(results[i].siret)
		fmt.Println(results[i].siren)
		fmt.Println(results[i].description)	
	}

	return nil
}

func (b *DataManager) GetJobWithName(pattern string) error {
	var results []Job

	r, err := b.conn.QueryNeo("MATCH (j:Job) WHERE j.name CONTAINS {pattern} RETURN j", map[string]interface{}{
		"pattern": pattern,
	})
	if err != nil {
		return fmt.Errorf("cannot query: %v", err)
	}
	defer r.Close()

	for data, _, err := r.NextNeo(); err != io.EOF; data, _, err = r.NextNeo() {
		d := data[0].(graph.Node)

		p := Job {
	        name:    d.Properties["name"].(string),
	        description:  d.Properties["description"].(string),
	        grossWage:  d.Properties["grossWage"].(string),
	    }
	   	results = append(results, p)
	}

	for i := range results {
		fmt.Println(results[i].name)
		fmt.Println(results[i].description)
		fmt.Println(results[i].grossWage)	
	}

	return nil
}

func (b *DataManager) SetJob(name string, description string, grossWage string) error {
	_, err := b.conn.ExecNeo("CREATE (j:Job {name: {name}, description: {description}, grossWage: {grossWage}})", map[string]interface{}{
		"name": name,
		"description": description,
		"grossWage": grossWage,
	})
	if err != nil {
		return fmt.Errorf("cannot query: %v", err)
	}

	return nil
}

func (b *DataManager) SetCompany(name string, siret string, siren string, description string) error {
	_, err := b.conn.ExecNeo("CREATE (c:Company {name: {name}, siret: {siret}, siren: {siren} , description: {description}})", map[string]interface{}{
		"name": name,
		"siret": siret,
		"siren": siren,
		"description": description,
	})
	if err != nil {
		return fmt.Errorf("cannot query: %v", err)
	}

	return nil
}

func (b *DataManager) SetProfile(firstName string, lastName string, birthday string, residence string) error {
	_, err := b.conn.ExecNeo("CREATE (p:Person {firstName: {firstName}, lastName: {lastName}, birthday: {birthday} , residence: {residence}})", map[string]interface{}{
		"firstName": firstName,
		"lastName": lastName,
		"birthday": birthday,
		"residence": residence,
	})
	if err != nil {
		return fmt.Errorf("cannot query: %v", err)
	}

	return nil
}
