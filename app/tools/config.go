package tools

import (
	"encoding/json"
	"gin_zyc/config"
	"io/ioutil"
)

var _cfg *config.App = nil

//解析json文件
func ParseConfig(path string) (*config.App, error) {
	content, err := ioutil.ReadFile("./config/" + path + ".json")
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(content, &_cfg)

	if err != nil {
		return nil, err
	}
	return _cfg, nil
}
