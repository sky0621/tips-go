package httpdoc

import (
	"bytes"
	"encoding/json"
	"testing"

	"net/http"

	"net/http/httptest"

	"go.mercari.io/go-httpdoc"
)

func TestListMovies(t *testing.T) {
	document := &httpdoc.Document{
		Name:           "Movie API",
		ExcludeHeaders: []string{"Content-Length"},
	}
	defer func() {
		if err := document.Generate("doc/movie.md"); err != nil {
			t.Fatalf("error: %s", err)
		}
	}()

	mux := http.NewServeMux()
	mux.Handle("/v1/movie", httpdoc.Record(&movieHandler{}, document, &httpdoc.RecordOption{
		Description: "List movies",
	}))

	testServer := httptest.NewServer(mux)
	defer testServer.Close()

	req := testNewRequest(t, testServer.URL+"/v1/movie")
	if _, err := http.DefaultClient.Do(req); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func testNewRequest(t *testing.T, urlStr string) *http.Request {
	var buf bytes.Buffer
	encoder := json.NewEncoder(&buf)
	encoder.SetIndent("", " ")
	encoder.Encode(&createUserRequest{
		Name:  "taro",
		Email: "taro@example.com",
		Attribute: attribute{
			Birthday: "1968-11-24",
		},
	})

	req, err := http.NewRequest("POST", urlStr, &buf)
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	req.Header.Add("X-Version", "2")
	return req
}
