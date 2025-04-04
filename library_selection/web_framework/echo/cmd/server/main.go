package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-playground/validator"
	"github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/sky0621/tips-go/library_selection/web_framework/echo/cmd/server/logger"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"time"
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
	e.Use(middleware.RequestID())
	e.Use(middleware.Secure())
	e.Use(middleware.Timeout())
	/*
	 * CORSの設定
	 * フロントエンドからのアクセスを許可する必要がある場合は追加
	 */
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:1323"},
	}))

	/*
	 * 出力ログレベルの設定
	 */
	e.Logger.SetLevel(log.DEBUG)
	/*
	 * 出力ログのフォーマットの設定
	 */
	//if l, ok := e.Logger.(*log.Logger); ok {
	//	l.SetHeader("${time_rfc3339} ${level}")
	//}
	/*
	 * ロガーのカスタムハンドラの登録
	 */
	slog.SetDefault(slog.New(&logger.AppLogHandler{Handler: slog.NewJSONHandler(os.Stdout, nil)}))

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

	/*
	 * CSRF
	 */
	g2 := e.Group("/csrf", middleware.CSRFWithConfig(middleware.CSRFConfig{
		TokenLookup: "header:X-CSRF-Token",
		TokenLength: 32,
		CookieName:  "csrf",
		CookiePath:  "/",
	}))
	g2.GET("/token", func(c echo.Context) error {
		// ミドルウェアによりコンテキストにセットされた CSRF トークンを取得
		csrfToken := c.Get("csrf").(string)
		html := fmt.Sprintf(`
			<html>
			<head>
				<title>CSRF ヘッダーサンプル</title>
			</head>
			<body>
				<p>CSRF トークン: <strong>%s</strong></p>
				<p>このトークンを、POST リクエスト時に HTTP ヘッダー "X-CSRF-Token" に設定してください。</p>
			</body>
			</html>
		`, csrfToken)
		return c.HTML(http.StatusOK, html)
	})
	g2.POST("/submit", func(c echo.Context) error {
		return c.String(200, "submit ok")
	})

	/*
	 * JWT
	 */
	e.POST("/login", loginSample)
	g3 := e.Group("/jwt")
	g3.Use(echojwt.JWT([]byte(jwtSecretKey)))
	g3.GET("/user", jwtSample)

	/*
	 * RequestLogger
	 */
	g4 := e.Group("/request_logger")
	g4.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogHost:      true,
		LogURI:       true,
		LogMethod:    true,
		LogStatus:    true,
		LogLatency:   true,
		LogRequestID: true,
		LogError:     true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			if v.Error == nil {
				slog.LogAttrs(c.Request().Context(), slog.LevelInfo, "REQUEST",
					slog.String("host", v.Host),
					slog.String("path", v.URI),
					slog.String("method", v.Method),
					slog.Int("status", v.Status),
					slog.Int("latency", int(v.Latency)),
					slog.String(logger.LogKeyTraceID, v.RequestID),
				)
			} else {
				slog.LogAttrs(c.Request().Context(), slog.LevelError, "REQUEST_ERROR",
					slog.String("host", v.Host),
					slog.String("path", v.URI),
					slog.String("method", v.Method),
					slog.Int("status", v.Status),
					slog.Int("latency", int(v.Latency)),
					slog.String(logger.LogKeyTraceID, v.RequestID),
					slog.String("error", v.Error.Error()),
				)
			}
			return nil
		},
	}))
	g4.GET("/user/:id", requestLoggerSample)

	/*
	 * ファイルダウンロード
	 */
	g5 := e.Group("/download")
	g5.GET("/", fileDownloadIndex)
	g5.GET("/exec", fileDownloadExec)
	g5.GET("/attachment", fileAttachment)

	/*
	 * ファイルアップロード
	 */
	g6 := e.Group("/upload")
	g6.GET("/", fileUploadIndex)
	g6.POST("/exec", fileUploadExec)

	/*
	 * RequestID
	 */
	e.GET("/request_id", requestIDSample)

	/*
	 * リライト
	 */
	e.Pre(middleware.RewriteWithConfig(middleware.RewriteConfig{
		Rules: map[string]string{
			"/old/*": "/$1",
		},
	}))

	e.GET("/sleep", func(c echo.Context) error {
		fmt.Println("sleep start")
		time.Sleep(5 * time.Second)
		fmt.Println("sleep end")
		return c.String(http.StatusOK, "sleep ok")
	})

	/*
	 * Graceful Shutdown
	 */
	ctxForOsInterrupt, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	go func() {
		if err := e.Start(":1323"); err != nil && !errors.Is(err, http.ErrServerClosed) {
			e.Logger.Fatal(err)
		}
	}()
	// OSシグナルを受信するまで待機
	<-ctxForOsInterrupt.Done()
	fmt.Println("notified os interrupt")

	ctxForTimeout, cancelFunc := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancelFunc()

	if err := e.Shutdown(ctxForTimeout); err != nil {
		e.Logger.Fatal(err)
	}
	fmt.Println("after server shutdown")
}
