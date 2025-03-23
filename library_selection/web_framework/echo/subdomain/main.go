package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

type Host struct {
	Echo *echo.Echo
}

func main() {
	hosts := map[string]*Host{}

	/*
	 * API
	 */
	api := echo.New()
	api.Use(middleware.Logger())
	api.Use(middleware.Recover())

	hosts["api.localhost:1323"] = &Host{api}

	api.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "API")
	})

	/*
	 * Blog
	 */
	blog := echo.New()
	blog.Use(middleware.Logger())
	blog.Use(middleware.Recover())

	hosts["blog.localhost:1323"] = &Host{blog}

	blog.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Blog")
	})

	/*
	 * Site
	 */
	site := echo.New()
	site.Use(middleware.Logger())
	site.Use(middleware.Recover())

	hosts["localhost:1323"] = &Host{site}

	site.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Site")
	})

	/*
	 * Server
	 */
	e := echo.New()
	e.Any("/*", func(c echo.Context) error {
		req := c.Request()
		res := c.Response()
		host := hosts[req.Host]
		if host == nil {
			return echo.ErrNotFound
		}
		host.Echo.ServeHTTP(res, req)
		return nil
	})
	e.Logger.Fatal(e.Start(":1323"))
}
