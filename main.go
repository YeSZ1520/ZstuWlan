package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	// 如果接收到的参数少于2个，则打印用法并退出
	if len(os.Args) < 3 {
		fmt.Println("Usage: ./ruijie_general.go username password")
		fmt.Println("Example: ./ruijie_general.go 201620000000 123456")
		os.Exit(1)
	}

	// 检查是否已经在线，使用 www.google.cn/generate_204 来检查在线状态
	resp, err := http.Get("http://www.google.cn/generate_204")
	if err != nil {
		fmt.Println("Error checking online status:", err)
		os.Exit(1)
	}
	if resp.StatusCode == http.StatusNoContent {
		fmt.Println("You are already online!")
		os.Exit(0)
	}

	// 如果未在线，则开始锐捷认证
	// 获取锐捷登录页面 URL
	resp, err = http.Get("http://www.google.cn/generate_204")
	if err != nil {
		fmt.Println("Error fetching login page URL:", err)
		os.Exit(1)
	}
	loginPageURL := resp.Request.URL.String()

	// 构造登录 URL
	loginURL := strings.Replace(loginPageURL, "index.jsp", "InterFace.do?method=login", 1)
	fmt.Println(loginURL)

	// 构造查询字符串
	queryString := strings.Split(loginPageURL, "?")[1]
	queryString = strings.ReplaceAll(queryString, "&", "%2526")
	queryString = strings.ReplaceAll(queryString, "=", "%253D")

	// 发送锐捷 eportal 认证请求并输出结果
	client := &http.Client{}
	req, err := http.NewRequest("POST", loginURL, strings.NewReader(fmt.Sprintf("userId=%s&password=%s&service=&queryString=%s&operatorPwd=&operatorUserId=&validcode=&passwordEncrypt=false", os.Args[1], os.Args[2], queryString)))
	if err != nil {
		fmt.Println("Error creating request:", err)
		os.Exit(1)
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.91 Safari/537.36")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	resp, err = client.Do(req)
	if err != nil {
		fmt.Println("Error sending auth request:", err)
		os.Exit(1)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			os.Exit(1)
		}
	}(resp.Body)

	fmt.Println("Auth result:", resp.Status)
}
