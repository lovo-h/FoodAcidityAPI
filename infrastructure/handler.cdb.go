package infrastructure

import (
	"database/sql"
	"log"
)

type HandlerCockroach struct {
	DB *sql.DB
}

func (handler *HandlerCockroach) Query(queryStr string, queryParams []interface{}, fnSignal func([][]byte)) error {
	// Execute query
	rows, queryErr := handler.DB.Query(queryStr, queryParams...)

	if queryErr != nil {
		log.Println(queryErr.Error())
		return queryErr
	}

	defer rows.Close()

	// Count columns & create DSes to store results
	cols, colsErr := rows.Columns()

	if colsErr != nil {
		return colsErr
	}

	rawResults := make([][]byte, len(cols))
	dest := make([]interface{}, len(cols))

	for idx := 0; idx < len(cols); idx++ {
		dest[idx] = &rawResults[idx]
	}

	// Read data
	for rows.Next() {
		if scanErr := rows.Scan(dest...); scanErr != nil {
			return scanErr
		}

		// Send string data
		fnSignal(rawResults)
	}

	return nil
}