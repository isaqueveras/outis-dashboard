package config

import (
	"errors"
)

// Database configure the settings for a database
type Database struct {
	Name string `env:"OUTIS_DB_NAME"`
	User string `env:"OUTIS_DB_USER"`
	Pass string `env:"OUTIS_DB_PASS"`
	Host string `env:"OUTIS_DB_HOST"`
	Port string `env:"OUTIS_DB_PORT"`

	MaxOpenConns    int `env:"OUTIS_DB_OPEN_CONNS"`
	MaxIdleConns    int `env:"OUTIS_DB_IDLE_CONNS"`
	Timeout         int `env:"OUTIS_DB_TIMEOUT"`
	ConnMaxLifetime int `env:"OUTIS_DB_CONN_MAX_LIFETIME"`
	ConnMaxIdleTime int `env:"OUTIS_DB_MAX_IDLE_TIME"`

	Cert SSL
}

// Validate validate data structure settings
func (db *Database) Validate() error {
	if db.Name == "" {
		return errors.New("defines the database name")
	}

	if db.User == "" {
		return errors.New("defines the database user")
	}

	if db.Pass == "" {
		return errors.New("defines the database password")
	}

	if db.Host == "" {
		return errors.New("defines the database host")
	}

	if db.Port == "" {
		return errors.New("defines the database port")
	}

	if db.MaxOpenConns == 0 {
		return errors.New("defines the maximum number of database connections")
	}

	if db.MaxIdleConns == 0 {
		return errors.New("set the number of idle database connections")
	}

	if db.ConnMaxLifetime == 0 {
		return errors.New("set the maximum database connection time")
	}

	if db.ConnMaxIdleTime == 0 {
		return errors.New("set the maximum time for idle database connections")
	}

	if err := db.Cert.Validar(); err != nil {
		return err
	}

	return nil
}
