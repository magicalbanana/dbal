package dbal

import "database/sql"

// RowsToMap takes the current sql.Rows and maps each column and value to a
// map[string]interface{}.
func RowsToMap(rows *sql.Rows) (map[int]map[string]interface{}, error) {
	columns, _ := rows.Columns()
	count := len(columns)
	values := make([]interface{}, count)
	valuePtrs := make([]interface{}, count)

	finalResult := map[int]map[string]interface{}{}
	resultID := 0
	for rows.Next() {
		for i := range columns {
			valuePtrs[i] = &values[i]
		}
		scanErr := rows.Scan(valuePtrs...)
		if scanErr != nil {
			return nil, scanErr
		}

		tmpStruct := map[string]interface{}{}

		for i, col := range columns {
			var v interface{}
			val := values[i]
			b, ok := val.([]byte)
			if ok {
				v = string(b)
			} else {
				v = val
			}
			tmpStruct[col] = v
		}

		finalResult[resultID] = tmpStruct
		resultID++
	}

	return finalResult, nil
}
