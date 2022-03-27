package test_db

import (
	"fmt"

	"github.com/schiller-sql/littSQL/config"
)

func main() {
	config.InitPostgresDB()
	fmt.Println("PostgreSQL configured correctly!")
}
