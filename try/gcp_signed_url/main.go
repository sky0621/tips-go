package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"cloud.google.com/go/storage"

	"golang.org/x/oauth2/google"
)

func main() {
	ctx := context.Background()

	creds, err := google.FindDefaultCredentials(ctx, storage.ScopeReadOnly)
	if err != nil {
		panic(err)
	}

	conf, err := google.JWTConfigFromJSON(creds.JSON, storage.ScopeReadOnly)
	if err != nil {
		panic(err)
	}

	expires, _ := time.Parse(time.RFC3339, "2019-12-27T09:33:00-00:00")
	opts := &storage.SignedURLOptions{
		GoogleAccessID: conf.Email,
		PrivateKey:     conf.PrivateKey,
		Method:         http.MethodGet,
		Expires:        expires,
	}
	url, err := storage.SignedURL("1517c1f4-5ed7-4087-be22-09ba412da599", "sample.mp4", opts)
	if err != nil {
		// サンプル用です。適切にエラーハンドリングしてください。
		panic(err)
	}
	fmt.Println(url)
}
