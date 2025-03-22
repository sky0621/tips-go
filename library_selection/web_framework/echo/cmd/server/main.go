package main

import (
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

type CustomContext struct {
	echo.Context
}

func (c *CustomContext) GetAppName() string {
	return "MyApp"
}

func main() {
	e := echo.New()
	/*
	 * サーバ起動時のコンソールへのバナーとポートの表示をオフ
	 */
	e.HideBanner = true
	e.HidePort = true
	/*
	 * カスタムコンテキストの登録（他のミドルウェアより前に登録する）
	 */
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &CustomContext{c}
			return next(cc)
		}
	})
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	/*
	 * 出力ログレベルの設定
	 */
	e.Logger.SetLevel(log.DEBUG)
	/*
	 * 出力ログのフォーマットの設定
	 */
	if l, ok := e.Logger.(*log.Logger); ok {
		l.SetHeader("${time_rfc3339} ${level}")
	}
	/*
	 * カスタムバリデーションの登録
	 */
	e.Validator = &CustomValidator{validator: validator.New()}

	e.GET("/", func(c echo.Context) error {
		cc := c.(*CustomContext)
		return cc.String(200, cc.GetAppName())
	})
	e.GET("/path/:id", pathParameterSample)
	e.GET("/query", queryParameterSample)
	e.POST("/users", postFormUrlEncodedSample)
	e.POST("/users2", handlingRequestSample)
	e.GET("/cookie", readCookie)
	e.POST("/cookie", writeCookie)
	e.POST("/custom_data_bind", customDataBindSample)
	e.POST("/validate", validateSample)

	/*
	 * ルートのグルーピング
	 */
	g := e.Group("/admin", middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		if username == "joe" && password == "secret" {
			return true, nil
		}
		return false, nil
	}))
	g.GET("/ping", func(c echo.Context) error {
		return c.String(200, "pong")
	})

	e.Logger.Fatal(e.Start(":1323"))
}
