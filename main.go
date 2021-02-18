// main.go

package main

import "os"

const host = "go-mux-tutorial-db.cnayavcudnyb.ap-southeast-1.rds.amazonaws.com"
const port = 5432

func main() {
	a := App{}
	a.Initialize(
		host,
		port,
		os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_NAME"))

	a.Run(":8010")
}
