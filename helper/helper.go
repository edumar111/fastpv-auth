package helper

import (
	"fmt"
	"os"
)

// Redis Variables
var host_redis = getEnv("HOST_REDIS", "localhost")
var port_redis = getEnv("PORT_REDIS", "6380")
var pass_redis = getEnv("PASS_REDIS", "fastpv123")

var RedisHost = "redis://" + pass_redis + "@" + host_redis + ":" + port_redis


// mysql variables
var host_mysql = getEnv("HOST_MYSQL", "localhost")
var port_mysql = getEnv("PORT_MYSQL", "3306")
var dbUser = getEnv("USER_MYSQL", "root")
var dbPass = getEnv("PASS_MYSQL", "123456")
var dbName   = getEnv("DB_MYSQL", "securitydb")
var MysqlHost = dbUser+":"+dbPass+"@tcp("+host_mysql+":"+port_mysql+")/"+dbName

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	fmt.Println(value)
	if len(value) == 0 {
		return defaultValue
	}
	return value
}
