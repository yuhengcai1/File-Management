package main

import (
	"database/sql"
	"log"

	
)

const (
	dbDriver = "postgres"
	dbSource = "postgres://postgres:postgrespw@postgres:5432"
)

var testQueries *Queries

type Users struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Store struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		Queries: New(db),
		db:      db,
	}
}

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("fall to create database connection", err)
	}
	testQueries = New(conn)
	
	

	
}
