package myHttp

import (
	"bufio"
	"fmt"
	"net/http"
)

func HttpClient() {
	resp, err := http.Get("https://gobyexample.com")
	if err != nil {
		fmt.Println("The programme Panicked", err)
	}

	defer resp.Body.Close()

	fmt.Println("Response status", resp.Status)

	scanner := bufio.NewScanner(resp.Body)
	for i := 0; scanner.Scan() && i < 5; i++ {
		fmt.Println("Scanner Text", scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("The system panicked", err)
	}
}
