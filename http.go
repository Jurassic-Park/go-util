package util

import (
	"bytes"
	"crypto/tls"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

// ContentTypeJSON is contentType for post
const ContentTypeJSON = "application/json;charset=UTF-8"

// Get 发送GET请求
// url：         请求地址
// response：    请求返回的内容
func Get(url string) string {

	// 超时时间：5秒
	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	var buffer [512]byte
	result := bytes.NewBuffer(nil)
	for {
		n, err := resp.Body.Read(buffer[0:])
		result.Write(buffer[0:n])
		if err != nil && err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
	}

	return result.String()
}

// Post 发送POST请求
// url：         请求地址
// data：        POST请求提交的数据
// contentType： 请求体格式，如：application/json
// content：     请求放回的内容
func Post(url string, data []byte, contentType string) (statusCode int, respBody []byte) {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	req.Header.Set("Content-Type", contentType)

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr, Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	result, _ := ioutil.ReadAll(resp.Body)
	return resp.StatusCode, result
}
