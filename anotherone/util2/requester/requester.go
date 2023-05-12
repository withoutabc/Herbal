package requester

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type Getter struct {
	Url    string
	Query  map[string]string
	Header map[string]string
}

func (g Getter) Get() ([]byte, error) {
	var urlSlice []string
	for k, v := range g.Query {
		urlSlice = append(urlSlice, k+"="+v)
	}
	var url = g.Url
	if len(urlSlice) != 0 {
		url = strings.Join(urlSlice, "&")
		url = g.Url + "?" + url
	}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	for k, v := range g.Header {
		req.Header.Add(k, v)
	}
	var client = &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, errors.New("创建请求失败")
	}
	defer resp.Body.Close()
	d, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("读取数据失败")
	}
	//fmt.Println(string(d))
	dst := map[string]interface{}{}
	if err := json.Unmarshal(d, &dst); err != nil {
		return nil, err
	}
	res, ok := dst["status"].(float64)
	if !ok {
		return nil, err
	}
	if res != 200 {
		return nil, errors.New(fmt.Sprintf("status=%.0f", res))
	}

	fmt.Println("Get")
	return d, nil
}
