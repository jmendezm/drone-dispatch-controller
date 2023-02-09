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
	
	CREATE TABLE drone_load (
    drone_serial_number INTEGER NOT NULL,
    name   TEXT    NOT NULL,
    weight REAL    NOT NULL,
    code   TEXT    NOT NULL,
    image  TEXT);`
	_, err := s.Conn.Exec(create)
	if err != nil {
		log.Fatal("sqlite_db | PrepareDB | creating tables", err)
	}

	data := `
		insert into drone values('AHG1234', 'Lightweight', 200, 98, 'IDLE');
		insert into drone values('AHG1235', 'Middleweight', 300, 53, 'LOADING');
		insert into drone values('AHG1236', 'Cruiserweight', 400, 60, 'DELIVERING');
		insert into drone values('AHG1237', 'Heavyweight', 490, 30, 'RETURNING');
		insert into drone values('AHG1238', 'Lightweight', 200, 23, 'IDLE');
		insert into drone values('AHG1239', 'Middleweight', 300, 41, 'LOADED');
		insert into drone values('AHG1240', 'Heavyweight', 450, 35, 'IDLE');
		insert into drone values('AHG1241', 'Middleweight', 300, 80, 'LOADING');
	`

	_, err = s.Conn.Exec(data)
	if err != nil {
		log.Fatal("sqlite_db | PrepareDB | inserting initial data | ", err)
	}
}
