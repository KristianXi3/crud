package handler

import (
	"database/sql"
	"fmt"

	_ "github.com/denisenkom/go-mssqldb"
)

type dbstruct struct {
	Connstr string
	SqlDb   *sql.DB
}

func ConnectSQL(connstr string) *dbstruct {
	con := dbstruct{
		Connstr: connstr,
	}

	db, err := sql.Open("sqlserver", con.Connstr)
	if err != nil {
		fmt.Printf("Fail to connect SQL Server: %v", err)
	}

	con.SqlDb = db
	con.SqlDb.SetMaxIdleConns(255)
	con.SqlDb.SetMaxOpenConns(255)

	return &con
}
