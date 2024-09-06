package database

import (
	"encoding/json"
	"errors"
	"os"
	"sync"
)

var ErrNotExist = errors.New("resource does not exist")

type DB struct {
	path string
	mu   *sync.RWMutex
}

type DBStructure struct {
	Recipe map[int]Recipe `json:"recipe"`
	User   map[int]User   `json:"user"`
}

func NewDB(path string) (*DB, error) {
	db := &DB{
		path: path,
		mu:   &sync.RWMutex{},
	}

	err := db.ensureDB()
	return db, err
}

func (db *DB) LoadDB() (DBStructure, error) {
	db.mu.Lock()
	defer db.mu.Unlock()

	dbStructure := DBStructure{}
	dat, err := os.ReadFile(db.path)
	if errors.Is(err, os.ErrNotExist) {
		return dbStructure, err
	}

	err = json.Unmarshal(dat, &dbStructure)
	if err != nil {
		return dbStructure, err
	}

	return dbStructure, nil
}

func (db *DB) ResetDB() error {
	err := os.Remove(db.path)
	if errors.Is(err, ErrNotExist) {
		return nil
	}

	return db.ensureDB()
}

func (db *DB) WriteDB(dbStructure DBStructure) error {
	db.mu.Lock()
	defer db.mu.Unlock()

	dat, err := json.Marshal(dbStructure)
	if err != nil {
		return err
	}

	err = os.WriteFile(db.path, dat, 0600)
	if err != nil {
		return err
	}

	return nil
}

func (db *DB) createDB() error {
	dbStructure := DBStructure{
		Recipe: map[int]Recipe{},
		User:   map[int]User{},
	}

	return db.WriteDB(dbStructure)
}

func (db *DB) ensureDB() error {
	_, err := os.ReadFile(db.path)
	if errors.Is(err, os.ErrNotExist) {
		return db.createDB()
	}

	return err
}
