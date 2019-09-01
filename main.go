package main

import (
	"log"
	"os"
)

func main() {
	var PORT = os.Getenv("SERVER_PORT")  // ":4444"
	var LOGLEVEL = os.Getenv("LOGLEVEL") // "debug"
	var DBURL = os.Getenv("MONGO_ADDRS") // "mongodb://localhost:27017/"
	var DBNAME = os.Getenv("MONGO_DB")   // "wbag"

	if PORT == "" {
		PORT = ":4444"
	}
	if LOGLEVEL == "" {
		LOGLEVEL = "debug"
	}
	if DBURL == "" {
		DBURL = "mongodb://localhost:27017/"
	}
	if DBNAME == "" {
		DBNAME = "wbag"
	}

	db := &DB{
		URL:  DBURL,
		Name: DBNAME,
	}
	err := db.Connect()
	if err != nil {
		log.Fatal(err)
	}

	app := newApp(db)
	err = app.Initialize(LOGLEVEL)
	if err != nil {
		log.Fatal(err)
	}
	app.Run(PORT)
}
