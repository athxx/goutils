package utils

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"sort"
	"strings"
	"time"
)

var CurlDebug bool

func CurlGet(router string, header map[string]string) ([]byte, error) {
	return curl(http.MethodGet, router, nil, header)
}

// CurlPost only support form and json, BODY support string, []byte, map[string]string
func CurlPost(router string, body interface{}, header map[string]string) ([]byte, error) {
	return curl(http.MethodPost, router, body, header)
}

func CurlPut(router string, body interface{}, header map[string]string) ([]byte, error) {
	return curl(http.MethodPut, router, body, header)
}

func CurlPatch(router string, body interface{}, header map[string]string) ([]byte, error) {
	return curl(http.MethodPatch, router, body, header)
}

// CurlDelete is curl delete
func CurlDelete(router string, body interface{}, header map[string]string) ([]byte, error) {
	return curl(http.MethodDelete, router, body, header)
}

func curl(method, router string, body interface{}, header map[string]string) ([]byte, error) {
	var reqBody io.Reader
	contentType := "application/json"
	switch v := body.(type) {
	case string:
		reqBody = strings.NewReader(v)
	case []byte:
		reqBody = bytes.NewReader(v)
	case map[string]string:
		val := url.Values{}
		for k, v := range v {
			val.Set(k, v)
		}
		reqBody = strings.NewReader(val.Encode())
		contentType = "application/x-www-form-urlencoded"
	case map[string]interface{}:
		val := url.Values{}
		for k, v := range v {
			val.Set(k, v.(string))
		}
		reqBody = strings.NewReader(val.Encode())
		contentType = "application/x-www-form-urlencoded"
	}
	if header == nil {
		header = map[string]string{"Content-Type": contentType}
	}
	if _, ok := header["Content-Type"]; !ok {
		header["Content-Type"] = contentType
	}
	resp, er := CurlReq(method, router, reqBody, header)
	if er != nil {
		return nil, er
	}
	res, err := ioutil.ReadAll(resp.Body)
	if CurlDebug {
		blob := SerializeStr(body)
		if contentType != "application/json" {
			blob = HttpBuild(body)
		}
		fmt.Printf("\n\n=====================\n[url]: %s\n[time]: %s\n[method]: %s\n[content-type]: %v\n[req_header]: %s\n[req_body]: %#v\n[resp_err]: %v\n[resp_header]: %v\n[resp_body]: %v\n=====================\n\n",
			router,
			time.Now().Format("2006-01-02 15:04:05.000"),
			method,
			contentType,
			HttpBuildQuery(header),
			blob,
			err,
			SerializeStr(resp.Header),
			string(res),
		)
	}
	resp.Body.Close()
	return res, err
}

func CurlReq(method, router string, reqBody io.Reader, header map[string]string) (*http.Response, error) {
	req, err := http.NewRequest(method, router, reqBody)
	if err != nil {
		return nil, err
	}
	for k, v := range header {
		req.Header.Set(k, v)
	}
	client := newHttpClient()
	return client.Do(req)
}

// CurlPostFile 以键值对Form表单形式上传文件
func CurlPostFile(uri string, fields map[string]string, files map[string]string, headers map[string]string) ([]byte, error) {
	body := &bytes.Buffer{}
	//create a multipart writer
	writer := multipart.NewWriter(body)

	if len(fields) > 0 {
		for k, v := range fields {
			writer.WriteField(k, v)
		}
	}
	if len(files) > 0 {
		for k, v := range files {
			file, err := os.Open(v)
			if err != nil {
				return nil, err
			}
			part, err := writer.CreateFormFile(k, v)
			if err != nil {
				return nil, err
			}
			io.Copy(part, file)
			file.Close()
		}
	}
	if err := writer.Close(); err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", uri, body)
	request.Header.Set("Content-Type", writer.FormDataContentType())
	if len(headers) > 0 {
		for k, v := range headers {
			request.Header.Set(k, v)
		}
	}
	clt := newHttpClient()
	defer clt.CloseIdleConnections()
	resp, err := clt.Do(request)
	if err != nil {
		return nil, err
	}
	res, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	return res, err
}

func newHttpClient() *http.Client {
	// 绕过github等可能因为特征码返回503问题
	// https://www.imwzk.com/posts/2021-03-14-why-i-always-get-503-with-golang/
	defaultCipherSuites := []uint16{0xc02f, 0xc030, 0xc02b, 0xc02c, 0xcca8, 0xcca9, 0xc013, 0xc009,
		0xc014, 0xc00a, 0x009c, 0x009d, 0x002f, 0x0035, 0xc012, 0x000a}
	return &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
				CipherSuites:       append(defaultCipherSuites[8:], defaultCipherSuites[:8]...),
			},
			//Proxy: http.ProxyFromEnvironment,
			//DialContext: (&net.Dialer{
			//	Timeout:   30 * time.Second,
			//	KeepAlive: 30 * time.Second,
			//}).DialContext,
			//ForceAttemptHTTP2:     true,
			//MaxIdleConns:          100,
			//IdleConnTimeout:       90 * time.Second,
			//TLSHandshakeTimeout:   10 * time.Second,
			//ExpectContinueTimeout: 1 * time.Second,
			//MaxConnsPerHost:       100,
			//MaxIdleConnsPerHost:   100,
		},
		// 获取301重定向
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
}

// HttpBuildQuery build request query,sortAsc true为小到大,false为大到小,nil不排序  a=123&b=321
func HttpBuildQuery(args map[string]string, sortAsc ...bool) string {
	str := ``
	if len(args) == 0 {
		return str
	}
	if len(sortAsc) > 0 {
		keys := make([]string, 0, len(args))
		for k := range args {
			keys = append(keys, k)
		}
		if sortAsc[0] {
			sort.Strings(keys)
		} else {
			sort.Sort(sort.Reverse(sort.StringSlice(keys)))
		}
		for _, k := range keys {
			str += "&" + k + "=" + args[k]
		}
	} else {
		for k, v := range args {
			str += "&" + k + "=" + v
		}
	}
	return str[1:]
}

func HttpBuild(body interface{}, sortAsc ...bool) string {
	params := map[string]string{}
	if args, ok := body.(map[string]interface{}); ok {
		for k, v := range args {
			params[k] = AnyToString(v)
		}
		return HttpBuildQuery(params, sortAsc...)
	}
	if args, ok := body.(map[string]string); ok {
		for k, v := range args {
			params[k] = AnyToString(v)
		}
		return HttpBuildQuery(params, sortAsc...)
	}
	if args, ok := body.(map[string]int); ok {
		for k, v := range args {
			params[k] = AnyToString(v)
		}
		return HttpBuildQuery(params, sortAsc...)
	}
	return AnyToString(body)
}
