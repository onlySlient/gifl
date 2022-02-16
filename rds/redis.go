package rds

import (
	"context"
	"fmt"
	"os"

	"github.com/go-redis/redis/v8"
	"github.com/teambition/gear/logging"
)

var RDB *redis.Client

var addr = func() string {
	return getEnv("RDS_ADDR", "localhost:6379")
}()

func init() {
	// RDB = redis.NewClient(&redis.Options{Addr: addr, Network: "tcp", Password: "", DB: 0})
	url, err := redis.ParseURL(addr)
	if err != nil {
		panic(err)
	}
	RDB = redis.NewClient(url)

	logging.Info(fmt.Sprintf("connection rds: %s", addr))
	err = RDB.Ping(context.Background()).Err()
	if err != nil {
		panic(err)
	}
}

func getEnv(k string, vs ...string) string {
	v := os.Getenv(k)
	if v != "" {
		return v
	}

	if len(vs) > 0 {
		return vs[0]
	}

	return ""
}
