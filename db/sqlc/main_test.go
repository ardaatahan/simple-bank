package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/ardaatahan/simplebank/util"
	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../../")
	if err != nil {
		log.Fatal("cannot load environment variables: ", err)
		return
	}
	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
		return
	}
	testQueries = New(testDB)
	os.Exit(m.Run())
}
