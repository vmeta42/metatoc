package myHttp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type Http struct {
	Client http.Client
	Method string
	Url    string
	Params []map[string]interface{}
	Data   map[string]interface{}
	Header map[string]interface{}
}

func New() *Http {
	return &Http{
		Client: http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (h *Http) SendRequest(result interface{}) error {
	// params
	paramsString := ""
	if len(h.Params) > 0 {
		for index, value := range h.Params {
			if index == 0 {
				paramsString += "?"
			}
			paramsString += fmt.Sprintf("%s=%v&", value["key"], value["value"])
		}
		paramsString += fmt.Sprintf("_=%d", time.Now().Unix())
	}

	// data
	buf := bytes.NewBuffer(nil)
	encoder := json.NewEncoder(buf)
	if err := encoder.Encode(h.Data); err != nil {
		return err
	}

	// new request
	fmt.Println("h.Method:", h.Method)
	fmt.Println("h.Url:", fmt.Sprintf("%s%s", h.Url, paramsString))
	fmt.Println("h.Data:", h.Data)
	request, err := http.NewRequest(h.Method, fmt.Sprintf("%s%s", h.Url, paramsString), buf)
	if err != nil {
		return err
	}

	// header
	request.Header.Add("Content-Type", "application/json")
	if len(h.Header) > 0 {
		for key, value := range h.Header {
			request.Header.Add(key, value.(string))
		}
	}

	response, err := h.Client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}
	newBody := strings.Replace(string(body), "\"", "", -1)
	newBody = strings.Replace(newBody, "\\", "\"", -1)
	//fmt.Println("newBody:", newBody)

	if err = json.Unmarshal([]byte(newBody), &result); err != nil {
		return err
	}

	return nil
}
