package httpdoc2

import (
	"net/http"
	"testing"

	"net/http/httptest"

	"encoding/json"

	"io"

	"strings"

	"go.mercari.io/go-httpdoc"
)

func TestGetUserList(t *testing.T) {
	// WebAPIドキュメントを表す構造体を生成
	doc := &httpdoc.Document{
		Name: "ユーザリソースへのCRUDを担うWebAPI",
	}
	defer func() {
		// WebAPIドキュメント出力
		err := doc.Generate("user.md")
		if err != nil {
			t.Fatalf("error: %s", err)
		}
	}()

	mux := http.NewServeMux()
	mux.Handle(
		"/v1/user",
		httpdoc.Record(
			&userHandler{},
			doc,
			&httpdoc.RecordOption{
				Description: "List users",
			},
		),
	)

	testServer := httptest.NewServer(mux)
	defer testServer.Close()

	req, err := http.NewRequest("POST", testServer.URL+"/v1/user", reqJson(t))
	if err != nil {
		t.Fatalf("error: %s", err)
	}

	_, err = http.DefaultClient.Do(req)
	if err != nil {
		t.Fatalf("error: %s", err)
	}
}

func reqJson(t *testing.T) io.Reader {
	req := &UserRequest{
		UserID:     1,
		UserName:   "たろー",
		UserStatus: 0,
	}
	userReq, err := json.Marshal(req)
	if err != nil {
		t.Fatalf("error: %s", err)
	}
	return strings.NewReader(string(userReq))
}
