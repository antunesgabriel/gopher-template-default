package helper

import "database/sql"

func GetStringFromNullString(s *sql.NullString) string {
	if s.Valid {
		return s.String
	}

	return ""
}
