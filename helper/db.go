package helper

import (
	"context"
	"database/sql"
	"fmt"
)

func CheckDuplicate(ctx context.Context, db *sql.DB, table string, field string, value interface{}) bool {
	SQL := fmt.Sprintf("SELECT %v FROM %v WHERE %v = ?", field, table, field)
	rows, err := db.QueryContext(ctx, SQL, value)
	PanicIfError(err)
	return rows.Next()
}
