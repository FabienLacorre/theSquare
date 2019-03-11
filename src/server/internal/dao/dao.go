package dao

import (
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
