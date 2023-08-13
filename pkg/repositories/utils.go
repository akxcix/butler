package repositories

import (
	"fmt"
)

func FormatPostgresDSN(user, pwd, host, port, db string) string {
	return fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s", user, pwd, host, port, db)
}
