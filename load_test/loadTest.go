package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"
)

type res struct {
	hostname string
	time     int64
}

func requestServer(url string, startTime int64) res {
	resp, err := http.Get(url)
	if err != nil {
		return res{fmt.Sprint(err), -1}
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return res{fmt.Sprint(err), -1}
	}
	var data map[string]interface{}
	if err := json.Unmarshal(body, &data); err != nil {
		return res{fmt.Sprint(err), -1}
	}
	return res{data["hostname"].(string), time.Now().UnixNano() - startTime}
}

func work(url string, n, t int64) (int, int, int64, map[string]int, map[string]int) {
	ch := make(chan res, n)
	for i := int64(0); i < n*t; i++ {
		go func() {
			ch <- requestServer(url, time.Now().UnixNano())
		}()
		time.Sleep(time.Microsecond * 1000000 / time.Duration(n))
		if 100*(i+1)/(n*t) > 100*i/(n*t) {
			fmt.Printf("进度: %d%% \r", 100*(i+1)/(n*t))
		}
	}
	missCnt, successCnt := 0, 0
	avgttl := int64(0)
	hostCnt, errCnt := make(map[string]int), make(map[string]int)
	for i := int64(0); i < n*t; i++ {
		req := <-ch
		if 100*(i+1)/(n*t) > 100*i/(n*t) {
			fmt.Printf("进度: %d%% \r", 100*(i+1)/(n*t))
		}
		if req.time < 0 {
			missCnt++
			errCnt[req.hostname]++
		} else {
			successCnt++
			avgttl += req.time
			hostCnt[req.hostname]++
		}
	}
	avgttl /= int64(successCnt)
	return missCnt, successCnt, avgttl, hostCnt, errCnt
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("请输入正确的参数")
		os.Exit(-1)
	}
	fmt.Println("URL:", os.Args[1])
	fmt.Println("Requests per second:", os.Args[2])
	fmt.Println("Duration:", os.Args[3], "s")
	cnt, _ := strconv.ParseInt(os.Args[2], 10, 64)
	t, _ := strconv.ParseInt(os.Args[3], 10, 64)
	missCnt, successCnt, avgttl, hostCnt, errCnt := work(os.Args[1], cnt, t)
	fmt.Printf("\n-------\nmiss: %d, success: %d, 平均访问时间: %f ms\n", missCnt, successCnt, float64(avgttl)/1000000.0)
	fmt.Println("-------\n各个Pod处理请求数（负载均衡实验）：")
	for k, v := range hostCnt {
		fmt.Println(k, ":", v)
	}
	fmt.Println("-------\nErrors：")
	for k, v := range errCnt {
		fmt.Println(k, ":", v)
	}
}
