// main.go

package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

//var host = "go-mux-tutorial-db.cnayavcudnyb.ap-southeast-1.rds.amazonaws.com"
const LOCALHOST_DB = false

func main() {
	host := os.Getenv("APP_DB_HOST")
	if host == "" {
		host = "localhost"
	}

	port := getIntEnv("APP_DB_PORT")
	if port == 0 {
		port = 5432
	}

	a := App{}
	a.Initialize(
		host,
		port,
		os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_NAME"))

	a.Run(":8010")
}

func getIntEnv(key string) int {
	val := getStrEnv(key)
	ret, err := strconv.Atoi(val)
	if err != nil {
		log.Fatal(fmt.Sprintf("Invalid value for environment variable " + key))
	}
	return ret
}

func getStrEnv(key string) string {
	val := os.Getenv(key)
	if val == "" {
		log.Fatal(fmt.Sprintf("Invalid value for environment variable " + key))
	}
	return val
}
