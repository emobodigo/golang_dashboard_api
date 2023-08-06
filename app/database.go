package app

import (
	"database/sql"
	"time"

	"github.com/emobodigo/golang_dashboard_api/helper"
)

func NewDB() *sql.DB {
	// dbSource := os.Getenv("DATABASE_URL")
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/dashboard_api")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
