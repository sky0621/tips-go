package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func fileDownloadIndex(c echo.Context) error {
	html := `<!doctype html>
<html lang="en">
<head>
    <meta charset="utf-8">
    <title>File download</title>
</head>
<body>
    <p><a href="/download/exec">File download</a></p>
    <p><a href="/download/attachment">File attachment</a></p>
</body>
</html>`
	return c.HTML(http.StatusOK, html)
}

const downloadFilePath = "filedownload/free.jpg"

func fileDownloadExec(c echo.Context) error {
	return c.File(downloadFilePath)
}

func fileAttachment(c echo.Context) error {
	return c.Attachment(downloadFilePath, "free.jpg")
}
