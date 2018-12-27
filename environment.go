package utils

import (
	"log"
	"os"
	"strconv"
	"time"
)

// GetEnvString returns a string value from the env or the default value, if non-existent
func GetEnvString(key, defaultVal string) string {
	if val, found := os.LookupEnv(key); found {
		return val
	}

	return defaultVal
}

// GetEnvInt returns a int value from the env or the default value, if non-existent
func GetEnvInt(key string, defaultVal int) int {
	if val, found := os.LookupEnv(key); found {
		if i, err := strconv.Atoi(val); err != nil {
			log.Panicf("Env variable with key=%s and value=%s cannot be converted to an Integer. Error=%v", key, val, err)
		} else {
			return i
		}
	}

	return defaultVal
}

// GetEnvInt64 returns a int64 value from the env or the default value, if non-existent
func GetEnvInt64(key string, defaultVal int64) int64 {
	if val, found := os.LookupEnv(key); found {
		if i, err := strconv.ParseInt(val, 10, 64); err != nil {
			log.Panicf("Env variable with key=%s and value=%s cannot be converted to an Int64. Error=%v", key, val, err)
		} else {
			return i
		}
	}

	return defaultVal
}

// GetEnvDuration returns a duration value from the env or the default value, if non-existent
func GetEnvDuration(key string, defaultVal time.Duration) time.Duration {
	if val, found := os.LookupEnv(key); found {
		if i, err := time.ParseDuration(val); err != nil {
			log.Panicf("Env variable with key=%s and value=%s cannot be converted to a Duration. Error=%v", key, val, err)
		} else {
			return i
		}
	}

	return defaultVal
}
