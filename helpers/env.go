package helpers

import (
	"os"
)

// GetEnv returns the value of an environment variable
// by name. If it doesn't exist the default will be returned
func GetEnv(name, def string) string {
	value := os.Getenv(name)
	if len(value) == 0 {
		return def
	}

	return value
}
