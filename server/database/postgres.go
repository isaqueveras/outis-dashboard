package database

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"

	"github.com/isaqueveras/outis-dashboard/server/config"
)

type postgres struct {
	db      *sql.DB
	timeout int
}

// open open a transaction with the database
func (p *postgres) open(cfg *config.Database) error {
	cfgSSL := "?sslmode=disable"
	if cfg.Cert.Mode != "" && cfg.Cert.Mode != "disable" {
		cert := cfg.Cert.Certificate
		cfgSSL = fmt.Sprintf("?sslmode=%s&sslrootcert=%s&sslcert=%s&sslkey=%s", cfg.Cert.Mode, cert.CA, cert.Cert, cert.Key)
	}

	url := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", cfg.User, cfg.Pass, cfg.Host, cfg.Port, cfg.Name)
	pgxPoolConfig, err := pgxpool.ParseConfig(url + cfgSSL)
	if err != nil {
		return err
	}

	pgxPoolConfig.ConnConfig.RuntimeParams = map[string]string{
		"application_name": "outis",
		"DateStyle":        "ISO",
		"IntervalStyle":    "iso_8601",
		"search_path":      "public",
	}

	db := stdlib.OpenDB(*pgxPoolConfig.ConnConfig)
	db.SetMaxOpenConns(cfg.MaxOpenConns)
	db.SetMaxIdleConns(cfg.MaxIdleConns)
	db.SetConnMaxLifetime(time.Duration(cfg.ConnMaxLifetime) * time.Second)
	db.SetConnMaxIdleTime(time.Duration(cfg.ConnMaxIdleTime) * time.Second)

	if err = db.Ping(); err != nil {
		return err
	}

	p.db, p.timeout = db, int(cfg.Timeout)
	return nil
}

func (p *postgres) close() {
	if p.db != nil {
		_ = p.db.Close()
	}
}

func (p *postgres) transaction(ctx context.Context) (*sql.Tx, error) {
	tx := new(sql.Tx)

	ctx, cancel := context.WithCancel(ctx)
	go func() {
		<-time.After(time.Duration(p.timeout+1) * time.Second)
		if tx == nil {
			cancel()
		}
	}()

	var err error
	if tx, err = p.db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelDefault, ReadOnly: false}); err != nil {
		return nil, err
	}

	return tx, nil
}
