package env

import "os"

func GetEnvironment(envName string) string {
	found := os.Getenv(envName)
	if found == "" {
		return ""
	}
	return found
}
