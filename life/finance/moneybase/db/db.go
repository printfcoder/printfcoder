package db

import (
	"context"
	"database/sql"
	"fmt"
	"sync"

	"github.com/stack-labs/stack/config"
	log "github.com/stack-labs/stack/logger"
)

var (
	c  DBConfig
	db *sql.DB

	s sync.Mutex
)

type DBConfig struct {
	DB struct {
		Dialect string   `sc:"dialect"`
		Postgre PGConfig `sc:"postgre"`
	} `sc:"db"`
}

type PGConfig struct {
	DBName            string `sc:"dbName"`            // The name of the database to connect to
	User              string `sc:"user"`              // the user to sign in as
	Password          string `sc:"password"`          // The user's password
	Host              string `sc:"host"`              // The host to connect to.Values that start with / are for unix domain sockets. (default is localhost)
	Port              int    `sc:"port"`              // The port to bind to.(default is 5432)
	SSLMode           string `sc:"sslMode"`           // Whether or not to use SSL (default is require, this is not the default for libpq)
	ConnectTimeout    int    `sc:"connectTimeout"`    // Maximum wait for connection, in seconds.Zero or not specified means wait indefinitely.
	SSLCert           string `sc:"sslCert"`           // Cert file location. The file must contain PEM encoded data.
	SSLKey            string `sc:"sslKey"`            // Key file location.The file must contain PEM encoded data.
	SSLRootCert       string `sc:"sslRootCert"`       // The location of the root certificate file.The file must contain PEM encoded data.)
	MaxOpenConnection int    `sc:"maxOpenConnection"` // use the default 0
	MaxIdleConnection int    `sc:"maxIdleConnection"` // use the default 0
}

func init() {
	config.RegisterOptions(&c)
}

func Init(ctx context.Context) error {
	s.Lock()
	defer s.Unlock()

	log.Info("begin to init db.")

	if db != nil {
		log.Warnf("db has been inited")
	}

	log.Infof("the dialect is %s.", c.DB.Dialect)
	if c.DB.Dialect == "postgre" {
		err := initPostgre(ctx)
		if err != nil {
			log.Errorf("load post err: %s")
			return fmt.Errorf("db init err: %s", err)
		}
	}

	return nil
}

func DB() *sql.DB {
	if db == nil {
		log.Fatalf("db is nil")
	}
	return db
}
