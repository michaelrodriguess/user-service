package config

import (
	"os"
)

func GetMariaDBURI() string {
	return os.Getenv("MARIADB_URI")
}
