package main

import (
	"flag"
	"log"
	"simple-http-probe/config"
)

var (
	configFile string
)
func main() {
	// 传入配置文件路径
	flag.StringVar(&configFile,"c","simple-http-probe.yaml","config file path")

	// 解析yaml
	conf,err:=config.LoadFile(configFile)
	if err != nil {
		log.Printf("config.Load.error: %v\n",err)
		return
	}
	log.Printf("配置是：%v\n",conf)
}
