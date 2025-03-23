package main

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

const (
	jwtSecretKey = "896DC975-0112-404C-8EBE-EA2A1F24D965"
)

func loginSample(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")
	/*
	 * 実際にはここでユーザー名とパスワードをチェックする
	 */
	fmt.Println(username, password)
	/*
	 * JWT Token を生成
	 */
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = username
	claims["admin"] = true
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	/*
	 * JWT用秘密鍵で署名
	 */
	t, err := token.SignedString([]byte(jwtSecretKey))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": t,
	})
}

// Authorization ヘッダーの Bearer <token> から JWT Token を取得できたら、ここに来る
func jwtSample(c echo.Context) error {
	// JWTミドルウェアによって、検証済みのトークンはコンテキストの "user" キーに格納されている
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)

	return c.String(http.StatusOK, "Hello "+name+"!")
}
