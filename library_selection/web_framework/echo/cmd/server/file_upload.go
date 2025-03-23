package main

import (
	"github.com/labstack/echo/v4"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

func fileUploadIndex(c echo.Context) error {
	html := `<!doctype html>
<html lang="en">
<head>
    <meta charset="utf-8">
    <title>File upload</title>
</head>
<body>
	<form action="/upload/exec" method="post" enctype="multipart/form-data">
		<input type="file" name="file" multiple>
		<input type="submit" value="Upload">
	</form>
</body>
</html>`
	return c.HTML(http.StatusOK, html)
}

func fileUploadExec(c echo.Context) error {
	form, err := c.MultipartForm()
	if err != nil {
		return err
	}
	files := form.File["file"]
	for _, file := range files {
		if err := saveFile(file); err != nil {
			return err
		}
	}
	return c.String(http.StatusOK, "upload done")
}

func saveFile(file *multipart.FileHeader) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	dst, err := os.Create(file.Filename)
	if err != nil {
		return err
	}
	defer dst.Close()

	if _, err := io.Copy(dst, src); err != nil {
		return err
	}

	return nil
}
