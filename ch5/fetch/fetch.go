package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
)

// Fetch downloads the URL and returns the
// name and length of the local file
func fetch(url string) (filename string, n int64, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()

	local := path.Base(resp.Request.URL.Path)
	if local == "/" {
		local = "index.html"
	}
	f, err := os.Create(local)
	if err != nil {
		return "", 0, err
	}
	n, err = io.Copy(f, resp.Body)
	// Close file, but prefer error from Copy, if any
	if closeErr := f.Close(); err == nil {
		err = closeErr
	}

	return local, n, err
}

func main() {
	filename, length, err := fetch(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	fmt.Printf("filename: %s, file length: %d\n", filename, length)
}
