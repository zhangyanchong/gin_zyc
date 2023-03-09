package tools

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

//解析json文件
func ParseConfig(path string) (map[string]interface{}, error) {
	AppPath, _ := os.Getwd()
	content, err := ioutil.ReadFile(AppPath + "/config/" + path + ".json")
	if err != nil {
		return nil, err
	}
	// Now let's unmarshall the data into `payload`
	var payload map[string]interface{}
	err = json.Unmarshal(content, &payload)

	if err != nil {
		return nil, err
	}
	return payload, nil
}
