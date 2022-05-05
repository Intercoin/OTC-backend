package postgres

import "database/sql"

func NullStringFromString(str string) sql.NullString {
	if str != "" {
		return sql.NullString{
			String: str,
			Valid:  true,
		}
	}
	return sql.NullString{}
}
