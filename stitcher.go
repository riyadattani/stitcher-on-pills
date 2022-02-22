package stitcher_on_pills

import (
	"bytes"
	"io"
	"net/http"
)

func StitchersURL(urls ...string) string {
	var responses []io.Reader
	for _, url := range urls {
		res, _ := http.Get(url)
		defer res.Body.Close()

		responses = append(responses, res.Body)
	}

	return Stitch(responses...)
}

func Stitch(readers ...io.Reader) string {
	var buf bytes.Buffer
	for _, rds := range readers {
		_, _ = io.Copy(&buf, rds)
	}

	return buf.String()
}
