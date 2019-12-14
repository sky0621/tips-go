package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"

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
	flag.Parse()

	res, err := http.Get(*mpath)
	if err != nil {
		return err
	}
	defer func() {
		if res != nil {
			if err := res.Body.Close(); err != nil {
				e = fmt.Errorf("[failed at response.body.close] %w", err)
			}
		}
	}()

	ctx := context.Background()
	client, err := storage.NewClient(ctx, option.WithCredentialsFile(*credentialPath))
	if err != nil {
		return err
	}
	//wc := client.Bucket(bucket).Object(object).NewWriter(ctx)
	//if _, err = io.Copy(wc, res.Body); err != nil {
	//	return err
	//}
	//if err := wc.Close(); err != nil {
	//	return err
	//}
	fmt.Println(client)

	return nil
}
