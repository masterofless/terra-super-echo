package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const HELLO = "Hello from Go!!"

var YODELERS = [4]string{"Hans", "Lars", "Bjorn", "Sigfried"}

func main() {
	shouter := echo.New()
	shouter.Use(middleware.Logger())
	shouter.Use(middleware.Recover())

	shouter.GET("/", hello)
	shouter.GET("/yodel", yodel)

	shouter.Logger.Fatal(shouter.Start(":8080")) // Start server
}

func hello(context echo.Context) error {
	return context.String(http.StatusOK, HELLO)
}

type YodelResponse struct {
	Yodel     string `json:"yodel"`
	Yodeler   string
	YodeledAt time.Time
}

func randomYodeler() string {
	var i = rand.Intn(len(YODELERS) - 1)
	return YODELERS[i]
}

func yodel(context echo.Context) error {
	t := time.Now()
	s := fmt.Sprintf("Yodelay Hee Who from go at %s @ %v!!!!!", t.Format(time.Kitchen), t.Location())
	answer := &YodelResponse{Yodel: s, Yodeler: randomYodeler(), YodeledAt: t}
	json, _ := json.Marshal(answer)
	return context.String(http.StatusOK, string(json))
}
