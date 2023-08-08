package helper

import (
	"context"
	"database/sql"
	"fmt"
)

func CheckDuplicate(ctx context.Context, tx *sql.Tx, table string, field string, value interface{}) bool {
	SQL := fmt.Sprintf("SELECT %v FROM %v WHERE %v = ?", field, table, field)
	rows, err := tx.QueryContext(ctx, SQL, value)
	PanicIfError(err)
	return rows.Next()
}
