package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestHello(t *testing.T) {
	e := echo.New()
	e.GET("/", hello)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set(echo.HeaderOrigin, "http://localhost")

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, hello(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, "Hello!", rec.Body.String())
	}
}
