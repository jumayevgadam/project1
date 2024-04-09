package dbcon

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Config struct {
	Host     string
	Port     string
	Username string
	DBName   string
	Password string
	SSLMode  string
}

func ConnectToDB(cfg Config) (*sqlx.DB, error) {
	DB, err := sqlx.Open("postgres", fmt.Sprintf("host = %s port = %s user = %s dbname = %s password = %s sslmode = %s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))
	if err != nil {
		return nil, err
	}

	err = DB.Ping()
	if err != nil {
		return nil, err
	}

	return DB, err

}
