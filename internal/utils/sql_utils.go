package utils

import "database/sql"

func SqlNonNullString(val string) sql.NullString {
	return sql.NullString{
		String: val,
		Valid:  true,
	}
}

func SqlMaybeNullString(val string) sql.NullString {
	return sql.NullString{
		String: val,
		Valid:  len(val) > 0,
	}
}
