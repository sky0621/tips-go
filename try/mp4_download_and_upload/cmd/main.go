package main

import (
	"context"
	"flag"
	"io"
	"log"
	"net/http"

	"github.com/pkg/errors"

	"google.golang.org/api/option"

	"cloud.google.com/go/storage"
)

func main() {
	if err := exec(); err != nil {
		log.Fatal(err)
	}
}

func exec() (e error) {
	credentialPath := flag.String("c", "credential file path", "~/keyfile/gcp.json")
	mpath := flag.String("m", "download url path", "http://localhost:9000/sample.mp4")
	writeBucket := flag.String("b", "write bucket", "http://localhost:7000/bucket")
	objectName := flag.String("o", "object name", "sample.mp4")
	flag.Parse()

	res, err := http.Get(*mpath)
	if err != nil {
		return err
	}
	defer func() {
		e = close(res.Body, e)
	}()

	ctx := context.Background()
	client, err := storage.NewClient(ctx, option.WithCredentialsFile(*credentialPath))
	if err != nil {
		return err
	}
	wc := client.Bucket(*writeBucket).Object(*objectName).NewWriter(ctx)
	if _, err = io.Copy(wc, res.Body); err != nil {
		return err
	}
	defer func() {
		e = closeWriter(wc, e)
	}()

	return nil
}

func close(c io.Closer, e error) error {
	if c == nil {
		return e
	}
	err := c.Close()
	if err == nil {
		return e
	}
	if e == nil {
		return err
	}
	return errors.Wrap(err, e.Error())
}

func closeWriter(w *storage.Writer, e error) error {
	if w == nil {
		return e
	}
	err := w.Close()
	if err == nil {
		return e
	}
	if e == nil {
		return err
	}
	return errors.Wrap(err, e.Error())
}
