package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

func getShowHandler(c echo.Context) error {
	team := c.QueryParam("team")
	member := c.QueryParam("member")
	return c.String(http.StatusOK, fmt.Sprintf("team:%s - member:%s", team, member))
}
