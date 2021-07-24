package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type Config struct {
	HttpListenAddr	string `yaml:"http_listen_addr"`
	HttpProbeTimeoutSecond	int	`yaml:"http_probe_timeout_second"`
}

// yaml解析
func Load(s []byte) (*Config,error) {
	cfg:=&Config{}
	err:=yaml.Unmarshal(s,cfg)

	if err != nil {
		return nil, err
	}
	return cfg,nil
}

// 读取配置文件
func LoadFile(filename string)(*Config,error)  {
	content,err :=ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	cfg,err := Load(content)
	if err != nil {
		log.Printf("load.yaml.error: %v\n",err)
		return nil, err
	}
	return cfg,nil
}

