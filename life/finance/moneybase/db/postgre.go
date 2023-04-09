package db

import (
	"context"
	"database/sql"
	"fmt"

	log "github.com/stack-labs/stack/logger"
)

func initPostgre(ctx context.Context) error {
	log.Infof("[initPG] init postgreSQL")

	var err error

	pgDB, err := sql.Open("postgres", parseConnectStr())
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	pgDB.SetMaxOpenConns(c.DB.Postgre.MaxOpenConnection)
	pgDB.SetMaxIdleConns(c.DB.Postgre.MaxIdleConnection)

	if err = pgDB.Ping(); err != nil {
		log.Fatal(err)
	}

	db = pgDB

	log.Infof("[initPG] pg connected")

	return nil
}

func parseConnectStr() string {
	str := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?search_path=%s", c.DB.Postgre.User, c.DB.Postgre.Password, c.DB.Postgre.Host,
		c.DB.Postgre.Port, c.DB.Postgre.DBName, c.DB.Postgre.SearchPath)

	log.Infof("[initPG] pg connected %s", str)

	str = fmt.Sprintf("%s&sslmode=%s", str, c.DB.Postgre.SSLMode)

	if c.DB.Postgre.SSLMode != "disable" {
		if c.DB.Postgre.SSLCert != "" {
			str += "&sslcert=" + c.DB.Postgre.SSLCert
		}

		if c.DB.Postgre.SSLKey != "" {
			str += "&sslkey=" + c.DB.Postgre.SSLKey
		}

		if c.DB.Postgre.SSLRootCert != "" {
			str += "&sslrootcert=" + c.DB.Postgre.SSLRootCert
		}
	} else {
		// do something
	}

	return str
}
