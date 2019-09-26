package store

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/winjeg/hamster/log"
	"github.com/winjeg/hamster/util"
	"sync"
)

const (
	createTableSql = `
	CREATE TABLE sites(
	   id       INTEGER PRIMARY KEY   AUTOINCREMENT,
	   author   CHAR(50) NOT NULL,
	   dir      CHAR(128)  NOT NULL,
	   mapping    CHAR(128)  NOT NULL,
	   created  INTEGER NOT NULL,
	   updated  INTEGER NOT NULL
	)`

	createIdxDirSql     = `CREATE UNIQUE INDEX idx_dir on sites (dir)`
	createIdxMappingSql = `CREATE UNIQUE INDEX idx_mapping on sites (mapping)`

	/*
		// may need later
		createAuthorIndexSql  = `CREATE INDEX idx_sites_author ON sites (author);`
		createDirIndexSql     = `CREATE INDEX idx_sites_dir ON sites (dir);`
		createMappingIndexSql = `CREATE INDEX idx_sites_mapping ON sites (mapping);`

	*/
)

var (
	once sync.Once
	lite *sql.DB
)

func GetDb() *sql.DB {
	if lite != nil {
		return lite
	}
	fileNameWithPath := util.GetCurrentPath("hamster.db")
	once.Do(func() {
		db, err := sql.Open("sqlite3", fileNameWithPath)
		lite = db
		// create table
		stmt, err := db.Prepare(createTableSql)
		if stmt != nil {
			_, err = stmt.Exec()
			_, i1err := db.Exec(createIdxDirSql)
			log.LogErr(i1err)
			_, i2err := db.Exec(createIdxMappingSql)
			log.LogErr(i2err)
			log.LogErr(err)
		}
	})
	return lite
}
