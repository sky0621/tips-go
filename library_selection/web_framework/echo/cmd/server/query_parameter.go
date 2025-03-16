package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

func queryParameterSample(c echo.Context) error {
	team := c.QueryParam("team")
	member := c.QueryParam("member")
	tags := c.QueryParam("tags")
	return c.String(http.StatusOK, fmt.Sprintf("team:%s - member:%s - tags:%s", team, member, tags))
}
