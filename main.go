package main

import (
	"log"

	"github.com/kataras/iris"
)

func main() {
	const PORT = ":4444"
	const LOGLEVEL = "debug"
	const DBURL = "mongodb://localhost:27017/"
	const DBNAME = "wbag"

	db := &DB{
		URL:  DBURL,
		Name: DBNAME,
	}
	err := db.Connect()
	if err != nil {
		log.Fatal(err)
	}

	app := newApp(db)
	err = app.SetupApp("debug")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Running server on http://localhost%s/\n", PORT)
	//app.Router.Run(iris.Addr(PORT), iris.WithoutServerError(iris.ErrServerClosed))
	app.Router.Run(iris.Addr(PORT))
}
