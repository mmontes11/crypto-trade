package database

import (
	"database/sql"

	// Clickhouse database driver: https://github.com/golang-migrate/migrate#databases
	_ "github.com/ClickHouse/clickhouse-go"
	"github.com/golang-migrate/migrate/v4"
	migrateDB "github.com/golang-migrate/migrate/v4/database"
	migrateCH "github.com/golang-migrate/migrate/v4/database/clickhouse"

	// Filesystem migration source driver: https://github.com/golang-migrate/migrate#migration-sources
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

// ClickHouse database
type ClickHouse struct {
	Config
	DB *sql.DB
}

const clickhouse = "clickhouse"

// NewClickHouse creates a new ClickHouse instance
func NewClickHouse(config Config) DB {
	return &ClickHouse{
		Config: config,
	}
}

// Connect establish connection
func (ch *ClickHouse) Connect() error {
	db, err := sql.Open(clickhouse, ch.URL)
	if err != nil {
		return err
	}
	if err := db.Ping(); err != nil {
		return err
	}

	ch.DB = db

	return nil
}

// MigrateUp runs migrations
func (ch *ClickHouse) MigrateUp() error {
	if ch.DB == nil {
		return ErrConnectionNotEstablished
	}

	m, err := ch.getMigrate()
	if err != nil {
		return err
	}

	return runMigration(m.Up)
}

// MigrateDown rollbacks migrations
func (ch *ClickHouse) MigrateDown() error {
	if ch.DB == nil {
		return ErrConnectionNotEstablished
	}

	m, err := ch.getMigrate()
	if err != nil {
		return err
	}

	return runMigration(m.Down)
}

func runMigration(migration func() error) error {
	err := migration()
	if err != migrate.ErrNoChange {
		return err
	}
	return nil
}

func (ch *ClickHouse) getMigrate() (*migrate.Migrate, error) {
	driver, err := ch.getMigrateDriver()
	if err != nil {
		return nil, err
	}

	return migrate.NewWithDatabaseInstance(ch.MigrationFiles, clickhouse, driver)
}

func (ch *ClickHouse) getMigrateDriver() (migrateDB.Driver, error) {
	return migrateCH.WithInstance(ch.DB, &migrateCH.Config{})
}
