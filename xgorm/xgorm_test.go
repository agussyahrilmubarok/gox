package xgorm_test

import (
	"testing"
	"time"

	"github.com/agussyahrilmubarok/gohelp/xgorm"
	"github.com/stretchr/testify/assert"

	"gorm.io/gorm"
)

// "gorm.io/driver/sqlite" Sqlite driver based on CGO
// "github.com/glebarez/sqlite" Pure go SQLite driver, checkout https://github.com/glebarez/sqlite for details
func newTestDB(opts ...*xgorm.Options) (*gorm.DB, error) {
	return xgorm.NewGorm("sqlite", "file::memory:?cache=shared", opts...)
}

func TestNewGorm_DefaultOptions(t *testing.T) {
	db, err := newTestDB()
	assert.NoError(t, err)
	assert.NotNil(t, db)

	sqlDB, err := db.DB()
	assert.NoError(t, err)

	assert.Equal(t, 20, sqlDB.Stats().MaxOpenConnections)
}

func TestNewGorm_CustomOptions(t *testing.T) {
	opts := &xgorm.Options{
		MaxIdleConns:    10,
		MaxOpenConns:    50,
		ConnMaxLifetime: time.Hour,
	}
	db, err := newTestDB(opts)
	assert.NoError(t, err)
	assert.NotNil(t, db)

	sqlDB, err := db.DB()
	assert.NoError(t, err)

	assert.Equal(t, opts.MaxOpenConns, sqlDB.Stats().MaxOpenConnections)
}

func TestNewGorm_UnsupportedDialect(t *testing.T) {
	db, err := xgorm.NewGorm("unknown", "dsn")
	assert.Nil(t, db)
	assert.Error(t, err)
}

func TestNewGorm_NilOptions(t *testing.T) {
	db, err := newTestDB(nil)
	assert.NoError(t, err)
	assert.NotNil(t, db)
}
