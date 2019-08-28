package main

import (
	"fmt"
	"log"
	"os"
	"testing"
)

var a *App
var db DB
var DBNAME = "wbag_test"

func TestMain(m *testing.M) {
	db := &DB{
		URL:  "mongodb://localhost:27017/",
		Name: DBNAME,
	}
	err := db.Connect()
	if err != nil {
		log.Fatal(err)
	}
	a = newApp(db)
	a.Initialize("debug")

	setupTestCollection()
	code := m.Run()
	cleanupTestCollection()
	os.Exit(code)
}

func setupTestCollection() {
	// Need to seed with a few dummy links
	fmt.Printf("Setup test collection in %s\n", DBNAME)
}

func cleanupTestCollection() {
	// Clear/Drop test collection when finished
	fmt.Printf("Dropping test collection in %s\n", DBNAME)
}
