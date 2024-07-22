package utils

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func DoRequest(method string, url string, data io.Reader, headers map[string]string) (string, error) {
	client := &http.Client{}

	var req *http.Request
	var err error
	// fmt.Println("data", data)
	if data != nil {
		req, err = http.NewRequest(method, url, data)
	} else {
		req, err = http.NewRequest(method, url, nil)
	}

	if err != nil {
		return "", err
	}

	// Set custom headers
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	log.Println("Response:", string(respBody))
	return string(respBody), nil
}

func DoRequestGet(method string, url string, data io.Reader, headers map[string]string) ([]byte, error) {
	client := &http.Client{}

	var req *http.Request
	var err error
	// fmt.Println("data", data)
	if data != nil {
		req, err = http.NewRequest(method, url, data)
	} else {
		req, err = http.NewRequest(method, url, nil)
	}

	if err != nil {
		return nil, err
	}

	// Set custom headers
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return respBody, nil
}

func PostBodyRequest(rurl string, data any, headers map[string]string) (string, error) {

	bytesData, _ := json.Marshal(data)
	reqBody := bytes.NewBuffer([]byte(bytesData))
	Response, err := DoRequest("POST", rurl, reqBody, headers)
	if err != nil {
		return "", err
	}

	return Response, nil
}

func PostFormDataRequest(path, filename, rurl string, headers map[string]string) (string, error) {
	ff, err := os.Open(path)
	if err != nil {
		return "", err
	}
	var body bytes.Buffer
	writer := multipart.NewWriter(&body)
	// 使用CreateFormFile方法添加文件
	part, err := writer.CreateFormFile("file", filename)
	if err != nil {
		return "", err
	}
	_, err = io.Copy(part, ff)
	if err != nil {
		return "", err
	}
	err = ff.Close()
	if err != nil {
		return "", err
	}
	err = writer.Close()
	if err != nil {
		return "", err
	}
	headers["Content-Type"] = writer.FormDataContentType()
	Response, err := DoRequest("POST", rurl, &body, headers)
	if err != nil {
		return "", err
	}
	return Response, nil
}

func PostUrlencodedRequest(rurl string, data, headers map[string]string) (string, error) {
	postData := url.Values{}
	for k, v := range data {
		postData.Add(k, v)
	}
	reqBody := strings.NewReader(postData.Encode())
	headers["Content-Type"] = "application/x-www-form-urlencoded"
	Response, err := DoRequest("POST", rurl, reqBody, headers)
	if err != nil {
		return "", err
	}
	return Response, nil
}

func PostFormDataRequestGet(path, filename, rurl string, headers map[string]string) (string, error) {
	ff, err := os.Open(path)
	if err != nil {
		return "", err
	}
	var body bytes.Buffer
	writer := multipart.NewWriter(&body)
	// 使用CreateFormFile方法添加文件
	part, err := writer.CreateFormFile("filename", filename)
	if err != nil {
		return "", err
	}
	_, err = io.Copy(part, ff)
	if err != nil {
		return "", err
	}
	err = ff.Close()
	if err != nil {
		return "", err
	}
	err = writer.Close()
	if err != nil {
		return "", err
	}
	headers["Content-Type"] = writer.FormDataContentType()
	Response, err := DoRequest("POST", rurl, &body, headers)
	if err != nil {
		return "", err
	}
	return Response, nil
}
