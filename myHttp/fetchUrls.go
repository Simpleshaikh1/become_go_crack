package myHttp

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func Fetch(url string, ch chan<- string, filename string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	defer resp.Body.Close()

	file, err := os.Create(filename)
	if err != nil {
		ch <- fmt.Sprintf("while creating file %s: %v", filename, err)
		return
	}
	defer file.Close()

	writer := io.MultiWriter(file, io.Discard)
	nbytes, err := io.Copy(writer, resp.Body)
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}

	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s -> saved to %s", secs, nbytes, url, filename)
}
