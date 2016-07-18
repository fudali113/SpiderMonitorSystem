package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func httpPost() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("出了错：", err)
		}
	}()
	resp, err := http.Post("http://192.168.1.15:9090/message",
		"application/json;charset=utf-8",
		strings.NewReader(`{"pc_id":"test","ip":"127.0.0.1"}`))
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	if string(body) != "" {
	} else {
		fmt.Print(time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04:05"))
		fmt.Println("-->" + string(body))
	}
}

//func main() {
//		var m = map[string]string{}
//		for k, v := range m {
//			fmt.Println(k + v)
//		}
//		fmt.Println("can")
//		fmt.Print(time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04:05"))
//			t1 := time.NewTimer(time.Millisecond * 1)
//			defer fmt.Print(time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04:05"))
//			for {
//				select {
//				case <-t1.C:
//					httpPost()
//					t1.Reset(time.Millisecond * 1)
//				}
//			}

//	a := make(chan int, 100)
//	for i := range make([]int, 110) {
//		select {
//		case a <- i:
//		default:
//		}
//	}
//	fmt.Println(len(a))
//}
