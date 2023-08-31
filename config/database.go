package config

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type DbConnection interface {
	Conn() *sql.DB
}

type dbConnection struct {
	db *sql.DB
	config *Config
}

func (d *dbConnection) initDb() error  {
	
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
	d.config.DbConfig.Host,
	d.config.DbConfig.Port,
	d.config.DbConfig.User,
	d.config.DbConfig.Password,
	d.config.DbConfig.Name,
	)
	
	db, err := sql.Open(d.config.DbConfig.Driver, dsn)
	if err != nil {
		return err
	}
	d.db = db
	return nil
}

func (d *dbConnection) Conn() *sql.DB  {
	return d.db
}

func NewDbConnection(config *Config) (DbConnection, error)  {
	conn := &dbConnection{
		config: config,
	}

	err := conn.initDb()
	if err != nil {
		return nil, err
	}

	return conn, nil
}