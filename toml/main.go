package main

import (
	"flag"
	"fmt"
	"io/ioutil"

	"github.com/influxdata/toml/ast"

	"github.com/influxdata/toml"
)

var serverConfigFile = flag.String("serverConf", "./serverConfig.toml", "The location of the server config file.")

// 1. 用 err = toml.Unmarshal(content, config)
type config struct {
	Server struct { // 自动映射为小写字母
		Endpoint        string `toml:"endpoint"`
		AesKey          string `toml:"aes_key"`
		TokenTTL        int64  `toml:"token_ttl"`
		MysqlURL        string `toml:"mysql_url"`
		MongoURI        string `toml:"mongo_uri"`
		ClientConfFiles string `toml:"client_config_files"`
	}
}

// 2. 用
type serverConfig struct {
	Endpoint        string `toml:"endpoint"`
	AesKey          string `toml:"aes_key"`
	TokenTTL        int64  `toml:"token_ttl"`
	MysqlURL        string `toml:"mysql_url"`
	MongoURI        string `toml:"mongo_uri"`
	ClientConfFiles string `toml:"client_config_files"`
}

func parseServerConfig1(configFilePath string) (*config, error) {
	content, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		return nil, err
	}
	config := &config{}
	err = toml.Unmarshal(content, config)
	if err != nil {
		return nil, err
	}
	return config, nil
}

func parseServerConfig2(configFilePath string) (*serverConfig, error) {
	content, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		return nil, err
	}
	config := &serverConfig{}
	tb, err := toml.Parse(content)
	//fmt.Println(tb)
	fmt.Println(tb.Fields["server"])
	t, _ := tb.Fields["server"].(*ast.Table)
	fmt.Println(t)
	err = toml.UnmarshalTable(t, config)
	return config, nil

}

func main() {
	// 1.
	//flag.Parse()
	//fmt.Println(*serverConfigFile)
	//config1, err := parseServerConfig1(*serverConfigFile)
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//fmt.Println(config1.Server)

	config2, _ := parseServerConfig2(*serverConfigFile)
	fmt.Println(config2)
}
