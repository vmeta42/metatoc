package core

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
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

func NewHttp() *Http {
	return &Http{
		Client: http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (server *NATSBridge) sendRequest(messageDataMap map[string]interface{}) error {
	h := NewHttp()
	h.Method = messageDataMap["type"].(string)
	h.Url = messageDataMap["url"].(string)
	//h.Params = messageDataMap["params"].([]map[string]interface{})
	h.Data = messageDataMap["data"].(map[string]interface{})

	// params
	paramsString := ""
	paramsString += fmt.Sprintf("?_=%d", time.Now().Unix())
	//if len(h.Params) > 0 {
	//	for index, value := range h.Params {
	//		if index == 0 {
	//			paramsString += "?"
	//		}
	//		paramsString += fmt.Sprintf("%s=%v&", value["key"], value["value"])
	//	}
	//	paramsString += fmt.Sprintf("_=%d", time.Now().Unix())
	//} else {
	//	paramsString += fmt.Sprintf("?_=%d", time.Now().Unix())
	//}

	// data
	buf := bytes.NewBuffer(nil)
	encoder := json.NewEncoder(buf)
	if err := encoder.Encode(h.Data); err != nil {
		return err
	}

	// new request
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

	server.RecordResponseBody(body, "sendRequest")
	return nil
}

func (server *NATSBridge) RecordResponseBody(body []byte, fileName string) {
	//filePath := fmt.Sprintf("./logs/%s.log", time.Now().Format("20060102"))
	//filePath := fmt.Sprintf("/opt/logs/%s.log", time.Now().Format("20060102"))
	filePath := fmt.Sprintf("./logs/%s.log", fileName)
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		log.Println(err.Error())
	}
	defer file.Close()
	write := bufio.NewWriter(file)
	write.WriteString(fmt.Sprintf("%s - %s \n", time.Now().Format("2006-01-02 15:04:05"), string(body)))
	write.Flush()

	server.logger.Noticef("response body is, %s", string(body))
}
