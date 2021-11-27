package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/YongpingGao/go-simple-bank/util"
	_ "github.com/lib/pq"
)

var testDB *sql.DB
var testQueries *Queries

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../../")
	if err != nil {
		log.Fatal("cannot load config", err)
	}

	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}
