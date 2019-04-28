package cache

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type CacheDB struct {
	db     *sql.DB
	dbPath string
}

// Create New CacheDB instance
func NewCacheDB(path string) (*CacheDB, error) {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}

	return &CacheDB{
		db:     db,
		dbPath: path,
	}, nil
}

// Close the CacheDB
func (cache *CacheDB) CloseCacheDB() error {
	err := cache.db.Close()
	if err != nil {
		// TODO
		fmt.Printf("Error occurred. %v", err)
	}

	return err
}
