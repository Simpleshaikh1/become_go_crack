package main

import (
	"GO-CRACK/array_slice"
	"GO-CRACK/myHttp"
	"fmt"
	"os"
	"time"
)

func main() {
	//myHttp.HttpClient()
	//fmt.Println("Http Response was successfully generated")

	//fetch Url
	//start := time.Now()
	ch := make(chan string)
	timestamp := time.Now().Format("20060102_150405")

	for i, url := range os.Args[1:] {
		filename := fmt.Sprintf("output_%d_%s.txt", i, timestamp)
		go myHttp.Fetch(url, ch, filename)
	}
	for range os.Args[1:] {
		fmt.Println(<-ch)
	}
	//fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())

	//using the append slice function

	var x, y []int
	for i := 0; i < 10; i++ {
		y = array_slice.AppendSlice(x, i)
		fmt.Printf("%d cap=%d\t%v\n", i, cap(y), y)
		x = y
	}
	//data := []int{1, 2, 3, 4, 5, 6}
	//ch := myHttp.Sender(data)
	//
	//fmt.Println("[  main] receivers")
	//go myHttp.Receiver("first", 10, ch)
	//go myHttp.Receiver("second", 3, ch)
	//
	//fmt.Println("[  main] waiting for execution")
	//time.Sleep(time.Second)
}
