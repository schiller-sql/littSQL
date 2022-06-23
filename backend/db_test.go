package main_test

import (
	"github.com/schiller-sql/littSQL/config"
	"testing"
)

func TestInitPostgresDB(t *testing.T) {
	config.InitConfigFile()
	if output := config.InitPostgresDB(); output == nil {
		t.Errorf("Output %v not equal to expected value", output)
	}
}
