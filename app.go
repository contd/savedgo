package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// App struct to make app testable
type App struct {
	Router *echo.Echo
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
		Router: echo.New(),
		DB:     db,
	}
	return app
}

// Initialize is the App initializer
func (app *App) Initialize(loglevel string) error {
	app.Router.Use(middleware.Logger())
	app.Router.Use(middleware.Recover())
	app.Router.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	// Method:   GET
	// Resource: http://localhost:4444/
	app.Router.Static("/", "./public")

	// Method:   GET
	// Resource: http://localhost:4444/id/:id
	app.Router.GET("/id/:id", func(ctx echo.Context) error {
		id := ctx.Param("id")
		link, _ := app.getItem(id)
		return ctx.JSON(http.StatusOK, link)
	})

	// Method:   GET
	// Resource: http://localhost:4444/contents/:id{int64}
	app.Router.GET("/contents/:id", func(ctx echo.Context) error {
		id := ctx.Param("id")
		link, _ := app.getItem(id)
		return ctx.JSON(http.StatusOK, link)
	})

	// Method:   GET
	// Resource: http://localhost:4444/tag/:tag
	app.Router.GET("/tag/:tag", func(ctx echo.Context) error {
		tag := ctx.Param("tag")
		link, err := app.DB.GetTagged("links", tag)
		if err != nil {
			log.Fatal(err)
			return err
		}
		return ctx.JSON(http.StatusOK, link)
	})

	// Method:   GET
	// Resource: http://localhost:4444/archived
	app.Router.GET("/archived", func(ctx echo.Context) error {
		links, err := app.DB.GetArchived("links")
		if err != nil {
			log.Fatal(err)
			return err
		}
		return ctx.JSON(http.StatusOK, links)
	})

	// Method:   GET
	// Resource: http://localhost:4444/archived/:tag
	app.Router.GET("/archived/:tag", func(ctx echo.Context) error {
		tag := ctx.Param("tag")
		links, err := app.DB.GetArchivedTag("links", tag)
		if err != nil {
			log.Fatal(err)
			return err
		}
		return ctx.JSON(http.StatusOK, links)
	})

	// Method:   GET
	// Resource: http://localhost:4444/starred
	app.Router.GET("/starred", func(ctx echo.Context) error {
		links, err := app.DB.GetStarred("links")
		if err != nil {
			log.Fatal(err)
			return err
		}
		return ctx.JSON(http.StatusOK, links)
	})

	// Method:   GET
	// Resource: http://localhost:4444/starred/:tag
	app.Router.GET("/starred/:tag", func(ctx echo.Context) error {
		tag := ctx.Param("tag")
		links, err := app.DB.GetStarredTag("links", tag)
		if err != nil {
			log.Fatal(err)
			return err
		}
		return ctx.JSON(http.StatusOK, links)
	})

	// Method:   GET
	// Resource: http://localhost:4444/tags
	app.Router.GET("/tags", func(ctx echo.Context) error {
		tags, err := app.DB.GetAllTags("links")
		if err != nil {
			log.Fatal(err)
			return err
		}
		return ctx.JSON(http.StatusOK, tags)
	})

	/* SEARCH
	   ******* NOT WORKING YET *************
	// Method:   GET
	// Resource: http://localhost:4444/search?q=some+words+or+phrase
	app.Router.GET("/search", func(ctx echo.Context) {
		q := ctx.URLParam("q")
		links, err := app.DB.SearchFor("links", q)
		if err != nil {
			log.Fatal(err)
		} else {
			ctx.JSON(links)
		}
	})
	*/

	// Method:   POST (Create)
	// Resource: http://localhost:4444/
	app.Router.POST("/", func(ctx echo.Context) error {
		link := new(Link)
		err := ctx.Bind(link)
		if err != nil {
			log.Fatalf("ReadJSON: %v\n", err)
			return err
		}
		_, err = app.DB.CreateOne("links", link)
		if err != nil {
			log.Fatalf("CreateOne: %v", err)
			return err
		}
		return ctx.JSON(http.StatusOK, link)
	})

	// Method:   PUT (Update)
	// Resource: http://localhost:4444/:id
	app.Router.PUT("/:id", func(ctx echo.Context) error {
		link := new(Link)
		err := ctx.Bind(link)
		if err != nil {
			log.Fatalf("ReadJSON: %v\n", err)
			return err
		}
		err = app.DB.UpdateOne("links", link)
		if err != nil {
			log.Fatalf("UpdateOne: %v", err)
			return err
		}
		return ctx.JSON(http.StatusOK, link)
	})

	// Method:   DELETE (Delete)
	// Resource: http://localhost:4444/:id
	app.Router.DELETE("/:id", func(ctx echo.Context) error {
		link := new(Link)
		err := ctx.Bind(link)
		if err != nil {
			log.Fatalf("ReadJSON: %v\n", err)
			return err
		}
		err = app.DB.DeleteOne("links", link)
		if err != nil {
			log.Fatalf("DeleteOne: %v", err)
			return err
		}
		return ctx.JSON(http.StatusOK, link)
	})

	return nil
}

// Run is the actual call to run (iris.Application.Run)
func (app *App) Run(PORT string) {
	log.Printf("Running server on http://localhost%s/\n", PORT)
	app.Router.Logger.Fatal(app.Router.Start(PORT))
}
