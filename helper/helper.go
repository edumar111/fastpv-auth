package helper

import (
	"fmt"
	"os"
)

//AddressExpetise address contrac expertise
var host_redis = getEnv("HOST_REDIS", "localhost")
var port_redis = getEnv("PORT_REDIS", "6380")
var pass_redis = getEnv("PASS_REDIS", "fastpv123")

var RedisHost = "redis://" + pass_redis + "@" + host_redis + ":" + port_redis

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	fmt.Println(value)
	if len(value) == 0 {
		return defaultValue
	}
	return value
}
