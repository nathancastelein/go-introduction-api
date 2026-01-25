package connect

import (
	"database/sql"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/stephenafamo/bob"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func getPostgreSQLDSN() string {
	if pgdsn := os.Getenv("PG_DSN"); pgdsn != "" {
		return pgdsn
	}
	return "postgres://gocourse:test@localhost:5432/gocourse?sslmode=disable"
}

func SQL() (*sql.DB, error) {
	return sql.Open("pgx", getPostgreSQLDSN())
}

func SQLx() (*sqlx.DB, error) {
	return sqlx.Open("pgx", getPostgreSQLDSN())
}

func Bob() (bob.DB, error) {
	db, err := sql.Open("pgx", getPostgreSQLDSN())
	if err != nil {
		return bob.DB{}, err
	}
	return bob.NewDB(db), nil
}

func Gorm() (*gorm.DB, error) {
	return gorm.Open(postgres.Open(getPostgreSQLDSN()), &gorm.Config{})
}
