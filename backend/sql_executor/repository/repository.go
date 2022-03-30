package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/mattn/go-sqlite3"
	_ "github.com/mattn/go-sqlite3"
	"github.com/schiller-sql/littSQL/sql_executor"
	"github.com/spf13/viper"
	"os"
)

type eRepository struct {
}

func NewRepository() sqlExecutor.Repository {
	return &eRepository{}
}

func (e eRepository) ExecuteSQLite(sqlString string) (*[]byte, error) {
	filename := viper.Get("SQLITE_DIR").(string) + "/" + uuid.NewString()
	db, err := sql.Open("sqlite3", "file:"+filename+"?cache=shared")
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
	content, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return &content, nil
}
