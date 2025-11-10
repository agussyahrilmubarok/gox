package xgorm

import (
	"fmt"
	"strings"
	"time"

	"github.com/glebarez/sqlite"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

// Options defines optional settings for the database connection.
// You can configure GORM options and connection pool settings here.
type Options struct {
	Config          *gorm.Config  // GORM configuration options
	MaxIdleConns    int           // Maximum number of idle connections in the pool
	MaxOpenConns    int           // Maximum number of open connections to the database
	ConnMaxLifetime time.Duration // Maximum lifetime of a connection
}

// NewGorm creates a GORM database connection using the specified dialect and DSN.
// You can provide optional settings through Options; defaults are applied if omitted.
func NewGorm(dialect, dsn string, opts ...*Options) (*gorm.DB, error) {
	var ops *Options
	if len(opts) > 0 && opts[0] != nil {
		ops = opts[0]
	} else {
		// Use default options if none provided
		ops = &Options{
			Config:          &gorm.Config{},
			MaxIdleConns:    5,
			MaxOpenConns:    20,
			ConnMaxLifetime: 30 * time.Minute,
		}
	}

	// Ensure Config is not nil
	if ops.Config == nil {
		ops.Config = &gorm.Config{}
	}
	// Set default pool settings if not set
	if ops.MaxIdleConns <= 0 {
		ops.MaxIdleConns = 5
	}
	if ops.MaxOpenConns <= 0 {
		ops.MaxOpenConns = 20
	}
	if ops.ConnMaxLifetime <= 0 {
		ops.ConnMaxLifetime = 30 * time.Minute
	}

	var db *gorm.DB
	var err error

	// Open database connection based on dialect
	switch strings.ToLower(dialect) {
	case "sqlite":
		db, err = gorm.Open(sqlite.Open(dsn), ops.Config)
	case "postgres", "postgresql":
		db, err = gorm.Open(postgres.Open(dsn), ops.Config)
	case "mysql":
		db, err = gorm.Open(mysql.Open(dsn), ops.Config)
	case "sqlserver", "mssql":
		db, err = gorm.Open(sqlserver.Open(dsn), ops.Config)
	default:
		return nil, fmt.Errorf("unsupported database dialect: %s", dialect)
	}

	if err != nil {
		return nil, fmt.Errorf("failed to connect database: %w", err)
	}

	// Apply connection pool settings
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxIdleConns(ops.MaxIdleConns)
	sqlDB.SetMaxOpenConns(ops.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(ops.ConnMaxLifetime)

	return db, nil
}
