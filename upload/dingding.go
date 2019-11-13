package upload

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"unsafe"
)

func dingdingAction() error {
	data := map[string]interface{}{
		"msgtype": "link",
		"link": map[string]string{
			"title":      "【" + dingDingInfo.Tag + "】: " + userInfo.Name + " 更新了前端",
			"text":       remoteServerInfo.Host + uploadInfo.RemoteDir + ":   " + dingDingInfo.Message,
			"messageUrl": dingDingInfo.URL,
		},
	}

	bytesData, err := json.Marshal(data)
	if err != nil {
		return errors.New("dingding json marshal error - " + err.Error())
	}

	// 发送
	request, err := http.NewRequest("POST", dingDingInfo.Push, bytes.NewReader(bytesData))
	request.Header.Set("Content-Type", "application/json;charset=UTF-8")
	client := http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return errors.New("dingding push error - " + err.Error())
	}

	// 获取 dingding push response
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.New("dingding push error - " + err.Error())
	}

	str := (*string)(unsafe.Pointer(&respBytes))
	fmt.Println()
	fmt.Println("dingding push response: ", *str)
	return nil
}
