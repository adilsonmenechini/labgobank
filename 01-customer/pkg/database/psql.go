package database

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/adilsonmenechini/golabbank/config"
	"github.com/adilsonmenechini/golabbank/pkg/utils"
	_ "github.com/lib/pq"
)

var (
	db     *sql.DB
	err    error
	logger *utils.Logger
)

func init() {
	logger = utils.NewLogger("database")
}

func dsnPsql() string {
	dbEnv := &config.EnvConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBname:   os.Getenv("DB_NAME"),
	}
	host, port, user, password, dbName := dbEnv.Host, dbEnv.Port, dbEnv.User, dbEnv.Password, dbEnv.DBname
	if host == "" || port == "" || user == "" || password == "" || dbName == "" {
		logger.Errorf("missing required environment variables")
	}
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)
}

// Connect function
func ConnectPSQL() *sql.DB {

	db, err = sql.Open("postgres", dsnPsql())
	if err != nil {
		logger.Errorf("Error %s when opening DB\n", err)
	}

	err = db.Ping()
	if err != nil {
		logger.Errorf("Errors %s pinging DB", err)
	}
	logger.Infof("Connected to database successfully")
	return db
}
