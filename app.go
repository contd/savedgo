package main

import (
	"log"

	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
)

// App struct to make app testable
type App struct {
	Router *iris.Application
	DB     *DB
}

func (app *App) getItem(id string) (*Link, error) {
	link, err := app.DB.GetOne("links", id)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return link, nil
}

func newApp(db *DB) *App {
	app := &App{
		Router: iris.Default(),
		DB:     db,
	}
	return app
}

// Initialize is the App initializer
func (app *App) Initialize(loglevel string) error {
	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
	})
	requestLogger := logger.New(logger.Config{
		// Status displays status code
		Status: true,
		// IP displays request's remote address
		IP: true,
		// Method displays the http method
		Method: true,
		// Path displays the request path
		Path: true,
		// Query appends the url query to the Path.
		Query: true,

		// if !empty then its contents derives from `ctx.Values().Get("logger_message")
		// will be added to the logs.
		MessageContextKeys: []string{"logger_message"},

		// if !empty then its contents derives from `ctx.GetHeader("User-Agent")
		MessageHeaderKeys: []string{"User-Agent"},
	})
	app.Router.Logger().SetLevel(loglevel)
	app.Router.Use(recover.New())
	app.Router.Use(requestLogger)
	app.Router.Use(crs)

	// Method:   GET
	// Resource: http://localhost:4444/
	assetHandler := app.Router.StaticHandler("./public", true, true)
	app.Router.SPA(assetHandler)

	// app.Router.Get("/", func(ctx iris.Context) {
	// 	links, err := app.DB.GetStarred("links")
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	} else {
	// 		ctx.JSON(links)
	// 	}
	// })

	// Method:   GET
	// Resource: http://localhost:4444/id/:id
	app.Router.Get("/id/{id:string max(255)}", func(ctx iris.Context) {
		id := ctx.Params().Get("id")
		link, _ := app.getItem(id)
		ctx.JSON(link)
	})

	// Method:   GET
	// Resource: http://localhost:4444/contents/:id{int64}
	app.Router.Get("/contents/{id:string max(255)}", func(ctx iris.Context) {
		id := ctx.Params().Get("id")
		link, _ := app.getItem(id)
		ctx.JSON(link)
	})

	// Method:   GET
	// Resource: http://localhost:4444/tag/:tag
	app.Router.Get("/tag/{tag:string max(255)}", func(ctx iris.Context) {
		tag := ctx.Params().Get("tag")
		link, err := app.DB.GetTagged("links", tag)
		if err != nil {
			log.Fatal(err)
		} else {
			ctx.JSON(link)
		}
	})

	// Method:   GET
	// Resource: http://localhost:4444/archived
	app.Router.Get("/archived", func(ctx iris.Context) {
		links, err := app.DB.GetArchived("links")
		if err != nil {
			log.Fatal(err)
		} else {
			ctx.JSON(links)
		}
	})

	// Method:   GET
	// Resource: http://localhost:4444/archived/:tag
	app.Router.Get("/archived/{tag:string max(255)}", func(ctx iris.Context) {
		tag := ctx.Params().Get("tag")
		links, err := app.DB.GetArchivedTag("links", tag)
		if err != nil {
			log.Fatal(err)
		} else {
			ctx.JSON(links)
		}
	})

	// Method:   GET
	// Resource: http://localhost:4444/starred
	app.Router.Get("/starred", func(ctx iris.Context) {
		links, err := app.DB.GetStarred("links")
		if err != nil {
			log.Fatal(err)
		} else {
			ctx.JSON(links)
		}
	})

	// Method:   GET
	// Resource: http://localhost:4444/starred/:tag
	app.Router.Get("/starred/{tag:string max(255)}", func(ctx iris.Context) {
		tag := ctx.Params().Get("tag")
		links, err := app.DB.GetStarredTag("links", tag)
		if err != nil {
			log.Fatal(err)
		} else {
			ctx.JSON(links)
		}
	})

	// Method:   GET
	// Resource: http://localhost:4444/tags
	app.Router.Get("/tags", func(ctx iris.Context) {
		tags, err := app.DB.GetAllTags("links")
		if err != nil {
			log.Fatal(err)
		} else {
			ctx.JSON(tags)
		}
	})

	/* SEARCH
	   ******* NOT WORKING YET *************
	// Method:   GET
	// Resource: http://localhost:4444/search?q=some+words+or+phrase
	app.Router.Get("/search", func(ctx iris.Context) {
		q := ctx.URLParam("q")
		links, err := app.DB.SearchFor("links", q)
		if err != nil {
			log.Fatal(err)
		} else {
			ctx.JSON(links)
		}
	})
	*/

	// Method:   POST
	// Resource: http://localhost:4444/
	app.Router.Post("/", func(ctx iris.Context) {
		link := new(Link)
		err := ctx.ReadJSON(link)
		if err != nil {
			log.Fatalf("ReadJSON: %v\n", err)
		}
		_, err = app.DB.CreateOne("links", link)
		if err != nil {
			log.Fatalf("CreateOne: %v", err)
		}
		ctx.JSON(link)
	})

	// Method:   POST
	// Resource: http://localhost:4444/:id

	// Method:   PUT
	// Resource: http://localhost:4444/:id

	// Method:   DELETE
	// Resource: http://localhost:4444/:id

	return nil
}

// Run is the actual call to run (iris.Application.Run)
func (app *App) Run(PORT string) {
	log.Printf("Running server on http://localhost%s/\n", PORT)
	app.Router.Run(iris.Addr(PORT))
}
