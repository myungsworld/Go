package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type Stats struct {
	Uptime       time.Time      `json:"uptime"`
	RequestCount uint64         `json:"requestCount"`
	Statuses     map[string]int `json:"statuses"`
	mutex        sync.RWMutex
}

func NewStats() *Stats {
	return &Stats{
		Uptime:   time.Now(),
		Statuses: map[string]int{},
	}
}

func indexHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!\n")
}

func getUser(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}

func show(c echo.Context) error {
	team := c.QueryParam("team")
	member := c.QueryParam("member")
	return c.String(http.StatusOK, "team:"+team+", member:"+member)
}

func createUser(c echo.Context) error {
	id := c.FormValue("id")
	email := c.FormValue("email")

	return c.String(http.StatusOK, "id : "+id+" email: "+email)
}

func writeCookie(c echo.Context) error {
	cookie := new(http.Cookie)
	cookie.Name = "song"
	cookie.Value = "jjang"
	cookie.Expires = time.Now().Add(24 * time.Hour)
	c.SetCookie(cookie)
	return c.String(http.StatusOK, "write a cookie")
}

func readCookie(c echo.Context) error {
	cookie, err := c.Cookie("song")
	if err != nil {
		return err
	}
	fmt.Println(cookie.Expires)
	fmt.Println(cookie.Value)
	return c.String(http.StatusOK, "쿠키 값"+cookie.Value)
}

func main() {
	e := echo.New()

	// s := NewStats()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", indexHandler)
	e.GET("/users/:id", getUser)
	e.GET("/show", show)
	e.POST("/usertt", createUser)
	e.GET("/setcookie", writeCookie)
	e.GET("/readcookie", readCookie)

	err := e.Start(":8000")
	if err != nil {
		e.Logger.Fatal(err)
	}

	e.Start(":8000")
}
