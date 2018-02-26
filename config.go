package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

// Config get config info from config.yml
type Config struct {
	Database struct {
		DbType  string `yaml:"type"`     //database type :mysql,sqlite,postgesql ...
		DbName  string `yaml:"name"`     //database name
		DbUser  string `yaml:"user"`     //database user
		DbPwd   string `yaml:"password"` //database user password
		DbHost  string `yaml:"host"`     //database host
		DbPort  string `yaml:"port"`     //database port
		DbModel string `yaml:"model"`    //database sslmodel only for postgresql
		DbChart string `yaml:"chart"`    //database character
	} `yaml:"database"` //数据库配置信息
}

// GetConfig get config info from current path file config.yml
func GetConfig() *Config {
	config := Config{}
	if configDir, err := filepath.Abs(filepath.Dir(os.Args[0])); err == nil {
		fmt.Println(configDir)

		if data, err := ioutil.ReadFile(filepath.Join(configDir, "config.yml")); err == nil {
			err = yaml.Unmarshal([]byte(data), &config)
		}

	}
	return &config
}
