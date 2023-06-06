package builtin

import "os"

func GetEnv(key string) string {
	return os.Getenv(key)
}

func HasEnv(key string) bool {
	_, exists := os.LookupEnv(key)
	return exists
}

func LookupEnv(key string) (string, bool) {
	return os.LookupEnv(key)
}
