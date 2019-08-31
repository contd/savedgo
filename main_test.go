package main

import (
	"fmt"
	"log"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/danielchatfield/go-chalk"
)

var a *App
var db DB
var DBNAME = "wbag_test"
var testLink = &Link{
	Title:          "Test Link Title",
	URL:            "http://example.com/test",
	Tags:           []string{"test", "tag"},
	DomainName:     "example.com",
	PreviewPicture: "/cache/e/0/e021d972/642383e4.png",
	ReadingTime:    1,
	CreatedAt:      "2018-11-23T15:49:36.000Z",
	UpdatedAt:      "2018-11-23T15:49:47.000Z",
	IsStarred:      true,
	IsArchived:     false,
	Content:        "<h1>Test Link Title</h1><p>This is the body</p>",
	Cached:         Cached{Status: "", FileName: ""},
	Updated:        time.Now(),
	Marked:         "#Test Link Title\nThis is the body\n",
}
var testID string

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
	// Use to seed some dummy data
	setupTestCollection()
	// Runs all tests (funcs that start with TestXxx)
	code := m.Run()
	// After all tests, cleanup/drop test collection
	cleanupTestCollection()
	os.Exit(code)
}

// TestNewApp checks default route returns index.html and data which comes from /starred route
func TestNewApp(t *testing.T) {
	e := httptest.New(t, a.Router)
	// Make sure status is OK for / (index.html)
	e.GET("/").Expect().Status(httptest.StatusOK)
	// Make sure content for /starred exists (default content for /)
	e.GET("/starred").Expect().Status(httptest.StatusOK).Body().NotEmpty()
}

func TestNonExistingRoute(t *testing.T) {
	e := httptest.New(t, a.Router)
	e.GET("/nonexist").Expect().Status(httptest.StatusNotFound)
}

func TestAllGetRoutes(t *testing.T) {
	e := httptest.New(t, a.Router)
	e.GET("/tag/golang").Expect().Status(httptest.StatusOK).Body().NotEmpty()
	e.GET("/archived").Expect().Status(httptest.StatusOK).Body().NotEmpty()
	e.GET("/archived/golang").Expect().Status(httptest.StatusOK).Body().NotEmpty()
	e.GET("/starred").Expect().Status(httptest.StatusOK).Body().NotEmpty()
	e.GET("/starred/golang").Expect().Status(httptest.StatusOK).Body().NotEmpty()
	e.GET("/tags").Expect().Status(httptest.StatusOK).Body().NotEmpty()
}

func TestCreateOne(t *testing.T) {
	newID, err := a.DB.CreateOne("links", testLink)
	if err != nil {
		log.Fatalf("CreateOne [test]: %v", err)
	}
	testID = newID
	testLink2 := &Link{
		Title:          "Test2 Link2 Title",
		URL:            "http://example2.com/test2",
		Tags:           []string{"test", "tag"},
		DomainName:     "example2.com",
		PreviewPicture: "/cache/e/0/e021d972/642383e4.png",
		ReadingTime:    1,
		CreatedAt:      "2018-11-23T15:49:36.000Z",
		UpdatedAt:      "2018-11-23T15:49:47.000Z",
		IsStarred:      true,
		IsArchived:     false,
		Content:        "<h1>Test Link Title</h1><p>This is the body</p>",
		Cached:         Cached{Status: "", FileName: ""},
		Updated:        time.Now(),
		Marked:         "#Test Link Title\nThis is the body\n",
	}
	e := httptest.New(t, a.Router)
	r := e.POST("/").WithJSON(testLink2).Expect().Status(httptest.StatusOK).JSON().Object()
	r.Value("title").Equal("Test2 Link2 Title")
	//e.GET(fmt.Sprintf("/id/%s", testID)).Expect().Status(httptest.StatusOK).Body().NotEmpty()
	//e.GET(fmt.Sprintf("/contents/%s", testID)).Expect().Status(httptest.StatusOK).Body().NotEmpty()
}

func TestUpdateOne(t *testing.T) {
	// Make a change to testLink for update
	testLink.URL = "http://example2.com"
	id := testLink.ID
	e := httptest.New(t, a.Router)
	e.PUT(fmt.Sprintf("/id/%s", id)).WithJSON(testLink).Expect().Status(httptest.StatusOK)
	//e.GET(fmt.Sprintf("/id/%s", testID)).Expect().Status(httptest.StatusOK).Body().NotEmpty()
}

func TestDeleteOne(t *testing.T) {
	id := testLink.ID
	e := httptest.New(t, a.Router)
	e.DELETE(fmt.Sprintf("/id/%s", id)).WithJSON(testLink).Expect().Status(httptest.StatusOK)
	//e.GET(fmt.Sprintf("/id/%s", testID)).Expect().Status(httptest.StatusNotFound)
}

func setupTestCollection() {
	// Setup anything needed ahead of all tests
	msg := chalk.Yellow("Setup test collection in")
	fmt.Printf("%s %s\n", msg, DBNAME)
}

func cleanupTestCollection() {
	// Clear/Drop test collection when finished
	err := a.DB.DropCollection("links")
	if err != nil {
		log.Fatalf("Drop [test]: %v", err)
	}
	msg := chalk.Yellow("Droped test collection in")
	fmt.Printf("%s %s\n", msg, DBNAME)
}
