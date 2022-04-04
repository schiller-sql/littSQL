package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/mattn/go-sqlite3"
	_ "github.com/mattn/go-sqlite3"
	"github.com/schiller-sql/littSQL/sql_executor"
	"io/ioutil"
	"os"
)

type eRepository struct {
	databasesDir string
}

func NewRepository(databasesDir string) sqlExecutor.Repository {
	return &eRepository{databasesDir}
}

func (e eRepository) ExecuteSQLite(sqlString string) (*[]byte, error) {
	file, err := ioutil.TempFile(e.databasesDir, "*.sqlite3")
	if err != nil {
		panic(err)
	}
	err = file.Close()
	if err != nil {
		panic(err)
	}
	path := file.Name()
	db, err := sql.Open("sqlite3", "file:"+path+"?cache=shared")
	if err != nil {
		panic(err)
	}
	_, err = db.Exec(sqlString)
	if err != nil {
		if errors.Is(err, sqlite3.Error{}) {
			switch err.(sqlite3.Error).Code {
			case 1:
				return nil, fmt.Errorf("error in sql: '%v'", err.Error())
			case 26:
				panic(err)
			}
		}
		return nil, err
	}
	content, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return &content, nil
}
