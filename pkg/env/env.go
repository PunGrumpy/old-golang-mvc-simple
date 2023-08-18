package env

import "os"

func GetEnvirontment(envName string) string {
	found := os.Getenv(envName)
	if found == "" {
		return ""
	}
	return found
}
