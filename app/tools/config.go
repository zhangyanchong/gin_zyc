package tools

import (
	"encoding/json"
	"gin_zyc/config"
	"io/ioutil"
	"os"
)

var _cfg *config.App = nil

//解析json文件
func ParseConfig(path string) (*config.App, error) {
	AppPath, _ := os.Getwd()
	content, err := ioutil.ReadFile(AppPath + "/config/" + path + ".json")
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(content, &_cfg)

	if err != nil {
		return nil, err
	}
	return _cfg, nil
}
