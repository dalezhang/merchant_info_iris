package config

import (
	"encoding/json"
	"fmt"
	"github.com/dalezhang/merchant_info_iris/utils"
	"io/ioutil"
	"os"
	//"path/filepath"
	"regexp"
	//"strings"
	//"unicode/utf8"
)

var jsonData map[string]interface{}

func initJSON() {
	bytes, err := ioutil.ReadFile("./config.json")
	if err != nil {
		fmt.Println("ReadFile: ", err.Error())
		os.Exit(-1)
	}

	configStr := string(bytes[:])
	reg := regexp.MustCompile(`/\*.*\*/`)

	configStr = reg.ReplaceAllString(configStr, "")
	bytes = []byte(configStr)

	if err := json.Unmarshal(bytes, &jsonData); err != nil {
		fmt.Println("invalid config: ", err.Error())
		os.Exit(-1)
	}
}

type dBConfig struct {
	Database string
	Host     string
	Port     int
	URL      string
}

// DBConfig 数据库相关配置
var DBConfig dBConfig

func initDB() {
	utils.SetStructByJSON(&DBConfig, jsonData["database"].(map[string]interface{}))
	url := fmt.Sprintf("%s:%d",
		DBConfig.Host, DBConfig.Port)
	DBConfig.URL = url
}

type serverConfig struct {
	Env  string
	Port int
}

// ServerConfig 服务器相关配置
var ServerConfig serverConfig

func initServer() {
	utils.SetStructByJSON(&ServerConfig, jsonData["go"].(map[string]interface{}))
}

func init() {
	initJSON()
	initDB()
	initServer()
}
