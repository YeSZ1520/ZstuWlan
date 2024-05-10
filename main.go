package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("参数错误：\nlogin 账号 密码")
		os.Exit(1)
	}
	fmt.Println(fmt.Sprintf("账号：[%v]", os.Args[1]))
	fmt.Println(fmt.Sprintf("密码：[%v]", os.Args[2]))
	loginUrl, queryString := getLoginUrl()
	result := login(loginUrl, os.Args[1], os.Args[2], queryString)
	fmt.Println(string(result))
}

func getLoginUrl() (string, string) {
	// 发送 GET 请求到 204 生成器
	resp, err := http.Get("http://www.google.cn/generate_204")
	if err != nil {
		errHanding("204 请求失败，请检查互联网连接")
	}
	defer resp.Body.Close()

	if resp.StatusCode == 204 {
		errHanding("已登录校园网")
		return "", ""
	} else {
		html, err := io.ReadAll(resp.Body)
		if err != nil {
			errHanding("读取 Http 相应失败")
		}
		loginUrl := regexp.MustCompile(`href='([^']*)index`).FindStringSubmatch(string(html))
		queryStrings := regexp.MustCompile(`index.jsp\?([^']*)'</script`).FindStringSubmatch(string(html))
		queryString := strings.Replace(strings.Replace(queryStrings[1], "&", "%2526", -1), "=", "%253D", -1)
		return loginUrl[1] + "InterFace.do?method=login", queryString
	}
}

func login(loginUrl, acc, passwd, queryString string) []byte {
	client := &http.Client{}
	loginPostData := fmt.Sprintf("userId=%v&password=%v&service=&queryString=%v&operatorPwd=&operatorUserId=&validcode=&passwordEncrypt=false", acc, passwd, queryString)
	request, err := http.NewRequest(http.MethodPost, loginUrl, strings.NewReader(loginPostData))
	if err != nil {
		errHanding("无法创建req")
	}

	request.Header.Set("Accept", "*/*")
	request.Header.Set("Accept-Encoding", "gzip, deflate")
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	request.Header.Set("Connection", "keep-alive")
	request.Header.Set("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6")
	request.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.0.0 Safari/537.36 Edg/124.0.0.0")

	resp, err := client.Do(request)
	if err != nil {
		errHanding("登录请求失败")
	}
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("读取登录结果失败")
	}
	return respBody
}

func errHanding(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
