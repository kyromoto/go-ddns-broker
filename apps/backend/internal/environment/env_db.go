package environment

import "os"

func DbSqliteFilename(defaultFilename string) string {
	filename, isSet := os.LookupEnv("DB_SQLITE_FILENAME")

	if !isSet {
		return defaultFilename
	}

	return filename
}
