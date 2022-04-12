package postgres

import (
	"database/sql"

	"github.com/golang-migrate/migrate/v4"
	migratepgx "github.com/golang-migrate/migrate/v4/database/pgx"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"

	"otc-backend/config"
	_ "otc-backend/repository/postgres/migrations"
)

func Connect(cfg config.DB) (*sqlx.DB, error) {
	db, err := sqlx.Connect("pgx", cfg.URI)
	if err != nil {
		return nil, err
	}

	if err = migrateDB(db.DB, cfg.MigrationTable); err != nil {
		return nil, err
	}

	return db, nil
}

func migrateDB(db *sql.DB, table string) error {
	driver, err := migratepgx.WithInstance(db, &migratepgx.Config{MigrationsTable: table})
	migrator, err := migrate.NewWithDatabaseInstance("embed://", table, driver)
	if err != nil {
		return err
	}

	err = migrator.Up()
	if err != nil && err.Error() == "no change" { // "no change" is not an error
		err = nil
	}
	return err
}
