package database

import (
	"errors"

	"github.com/nathancastelein/go-course-api/database/connect"
)

func SimpleQuery() (bool, error) {
	query := `SELECT 1 = 1;`

	db, err := connect.SQL()
	if err != nil {
		return false, err
	}

	rows, err := db.Query(query)
	if err != nil {
		return false, err
	}

	if rows.Next() {
		var result bool
		if err := rows.Scan(&result); err != nil {
			return false, err
		}

		return result, nil
	}

	return false, errors.New("no data")
}
