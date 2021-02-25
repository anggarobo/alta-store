package controller

import (
	"alta-store/config"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func InitEcho() *echo.Echo {
	config.InitDB()
	e := echo.New()
	return e
}

func TestDeleteCartOnController(t *testing.T) {

	var testCases = []struct {
		name                 string
		path                 string
		expectStatus         int
		expectBodyStartsWith string
	}{
		{
			name:                 "success",
			path:                 "/carts",
			expectBodyStartsWith: "{\"status\":\"success\",\"carts\":",
			expectStatus:         http.StatusOK,
		},
	}

	e := InitEcho()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	for _, testCase := range testCases {
		c.SetPath(testCase.path)

		if assert.NoError(t, GetCartController(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			body := rec.Body.String()

			fmt.Println(body)
		}
	}
}
