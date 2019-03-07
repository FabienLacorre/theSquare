package dao

import (
	"io"
	"fmt"

	"github.com/davecgh/go-spew/spew"
	bolt "github.com/moutoum/golang-neo4j-bolt-driver"
)

type DataManager struct {
	conn bolt.Conn
	
}

func NewDataManager(conn bolt.Conn) *DataManager {
	return &DataManager{
		conn: conn,
	}
}

func (b *DataManager) GetProfileWithID(id int) error {
	r, err := b.conn.QueryNeo("MATCH (p:Person) WHERE ID(p) = {id} RETURN p", map[string]interface{}{"id": id})
	defer r.Close()
	if err != nil {
		return fmt.Errorf("cannot query: %v", err)
	}
	for data, _, err := r.NextNeo(); err != io.EOF; data, _, err = r.NextNeo() {
		spew.Dump(data)
	}
	return nil
}

func (b *DataManager) GetProfileWithName(pattern string) error {
	r, err := b.conn.QueryNeo("MATCH (p:Person) WHERE p.name CONTAINS {pattern} RETURN p", map[string]interface{}{"pattern": pattern})
	defer r.Close()
	if err != nil {
		return fmt.Errorf("cannot query: %v", err)
	}
	for data, _, err := r.NextNeo(); err != io.EOF; data, _, err = r.NextNeo() {
		spew.Dump(data)
	}
	return nil
}

// type Profile {

// }

// func main() {
// 	var p Profile

// 	if err := bdd.GetProfileWithID(8, &p); err != nil {

// 	}
// }
