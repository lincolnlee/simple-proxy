package simple_proxy

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
)

type Config struct {
	Server     string `json:"server"`
	ServerPort int    `json:"server_port"`
	Client     string `json:"client"`
	Password   string `json:"password"`
	Proxy      string `json:"proxy"`
	ProxyPort  int    `json:"proxy_port"`
}

func ParseConfig(path string) (config *Config, err error) {
	file, err := os.Open(path) // For read access.
	if err != nil {
		return
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return
	}

	config = &Config{}
	if err = json.Unmarshal(data, config); err != nil {
		return nil, err
	}
	return
}

func UpdateConfig(old, new *Config) {
	newVal := reflect.ValueOf(new).Elem()
	oldVal := reflect.ValueOf(old).Elem()

	for i := 0; i < newVal.NumField(); i++ {
		newField := newVal.Field(i)
		oldField := oldVal.Field(i)

		switch newField.Kind() {
		case reflect.Interface:
			if fmt.Sprintf("%v", newField.Interface()) != "" {
				oldField.Set(newField)
			}
		case reflect.String:
			s := newField.String()
			if s != "" {
				oldField.SetString(s)
			}
		case reflect.Int:
			i := newField.Int()
			if i != 0 {
				oldField.SetInt(i)
			}
		}
	}
}
