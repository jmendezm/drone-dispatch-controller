package infra

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	log "github.com/sirupsen/logrus"
)

type SqliteDB struct {
	Conn *sql.DB
}

func (s *SqliteDB) InitSqliteDB() {
	var err error
	s.Conn, err = sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatal(err)
	}
}

func (s *SqliteDB) PrepareDB() {
	create := `CREATE TABLE drone (
    serial_number    TEXT (0, 100) PRIMARY KEY NOT NULL UNIQUE,
    model            TEXT          NOT NULL,
    weight_limit     REAL          NOT NULL,
    battery_capacity REAL          NOT NULL,
    state            TEXT          NOT NULL);
	
	CREATE TABLE medication (
    id     INTEGER PRIMARY KEY AUTOINCREMENT UNIQUE NOT NULL,
    name   TEXT    NOT NULL,
    weight REAL    NOT NULL,
    code   TEXT    NOT NULL,
    image  TEXT);`
	_, err := s.Conn.Exec(create)
	if err != nil {
		log.Fatal("sqlite_db | PrepareDB | creating tables", err)
	}
}
