package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", hello)
	e.GET("/yodel", yodel)

	e.Logger.Fatal(e.Start(":8080")) // Start server
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello from Go!!")
}

type response1 struct {
	Yodel string `json:"yodel"`
}

func yodel(c echo.Context) error {
	t := time.Now()
	s := fmt.Sprintf("Yodelay Hee Who from go at %s @ %v!!!!!", t.Format(time.Kitchen), t.Location())
	answer := &response1{Yodel: s}
	json, _ := json.Marshal(answer)
	return c.String(http.StatusOK, string(json))
}
