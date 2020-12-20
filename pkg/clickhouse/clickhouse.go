package clickhouse

import (
	"database/sql"

	// Clickhouse database driver: https://github.com/golang-migrate/migrate#databases
	_ "github.com/ClickHouse/clickhouse-go"
	"github.com/golang-migrate/migrate/v4"
	migrateCH "github.com/golang-migrate/migrate/v4/database/clickhouse"

	// Filesystem migration source driver: https://github.com/golang-migrate/migrate#migration-sources
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

const clickhouse = "clickhouse"

// Connect establishes connection
func Connect(URL string) (*sql.DB, error) {
	db, err := sql.Open(clickhouse, URL)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

// MigrateUp runs migrations
func MigrateUp(db *sql.DB, migrationsDir string) error {
	m, err := getMigrate(db, migrationsDir)
	if err != nil {
		return err
	}

	return runMigration(m.Up)
}

// MigrateDown rollbacks migrations
func MigrateDown(db *sql.DB, migrationsDir string) error {
	m, err := getMigrate(db, migrationsDir)
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

func getMigrate(db *sql.DB, migrationsDir string) (*migrate.Migrate, error) {
	driver, err := migrateCH.WithInstance(db, &migrateCH.Config{})
	if err != nil {
		return nil, err
	}

	return migrate.NewWithDatabaseInstance(migrationsDir, clickhouse, driver)
}
