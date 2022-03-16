package test_db

import (
	"github.com/schiller-sql/littSQL/config"
	"fmt"
)

func main() {
	config.InitPostgresDB()
	fmt.Println("PostgreSQL configured correctly!")
}
