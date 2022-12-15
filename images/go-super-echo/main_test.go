package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestHello(t *testing.T) {
	echoer := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", strings.NewReader("hello"))
	rec := httptest.NewRecorder()
	context := echoer.NewContext(req, rec)

	if assert.NoError(t, hello(context)) {
		assert.Equal(t, http.StatusOK, rec.Code)  // test status
		assert.Equal(t, HELLO, rec.Body.String()) // test response body
	}
}

func TestYodel(t *testing.T) {
	echoer := echo.New()
	var now time.Time
	now = time.Now()
	req := httptest.NewRequest(http.MethodGet, "/", strings.NewReader("yodel"))
	rec := httptest.NewRecorder()
	context := echoer.NewContext(req, rec)

	if assert.NoError(t, yodel(context)) {
		assert.Equal(t, http.StatusOK, rec.Code) // test status
		// test response body
		var yodel YodelResponse
		err := json.Unmarshal(rec.Body.Bytes(), &yodel)
		assert.Nil(t, err)
		assert.NotNil(t, yodel.Yodeler)
		assert.NotNil(t, yodel.YodeledAt)
		assert.True(t, now.Before(yodel.YodeledAt))
		assert.Contains(t, yodel.Yodel, "Yodelay Hee Who")
	}
}
