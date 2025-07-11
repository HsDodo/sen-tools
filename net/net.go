package net

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func Ip2binary(ip string) string {
	str := strings.Split(ip, ".")
	var ipstr string
	for _, s := range str {
		i, err := strconv.ParseUint(s, 10, 8)
		if err != nil {
			fmt.Println(err)
		}
		ipstr = ipstr + fmt.Sprintf("%08b", i)
	}
	return ipstr
}

// Match 验证IP地址和地址端是否匹配 变量ip为字符串，
// 例子ip "192.168.56.4" iprange为地址端"192.168.56.64/26"
func Match(ip, iprange string) bool {
	ipb := Ip2binary(ip)
	ipr := strings.Split(iprange, "/")
	masklen, err := strconv.ParseUint(ipr[1], 10, 32)
	if err != nil {
		fmt.Println(err)
		return false
	}
	iprb := Ip2binary(ipr[0])
	return strings.EqualFold(ipb[0:masklen], iprb[0:masklen])
}

func HttpPost(url string, body io.Reader) ([]byte, error) {

	req, err := http.NewRequest(http.MethodPost, url, body)
	if err != nil {
		return nil, err
	}

	return httpClientDo(req)
}

func httpClientDo(req *http.Request) ([]byte, error) {

	client := &http.Client{
		Timeout: time.Second * 15,
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	bs, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	} else {
		_ = resp.Body.Close()
		return bs, nil
	}
}
