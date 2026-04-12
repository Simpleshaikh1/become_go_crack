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

//
//func Sender(data []int) <-chan int {
//	fmt.Println("[sender] make channel")
//	ch := make(chan int)
//
//	go func() {
//		defer func() {
//			fmt.Println("[sender] close channel")
//			close(ch)
//		}()
//		for _, item := range data {
//			ch <- item
//		}
//		fmt.Println("[sender] all data sent to channel")
//	}()
//
//	fmt.Println("[sender] return <-channel")
//	return ch
//}
//
//func Receiver(name string, delay time.Duration, ch <-chan int) {
//	for val := range ch {
//		fmt.Printf("[%6s] data = %d\n", name, val)
//		time.Sleep(delay * time.Millisecond)
//	}
//	fmt.Printf("[%6s] stopped\n", name)
//}
