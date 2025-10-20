package main

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)



type Config struct {
	Mysql Mysql `yaml:"mysql"`
	Redis Redis `yaml:"redis"`
}

type Mysql struct {
	Url  string `yaml:"url"`
	Port int    `yaml:"port"`
}

type Redis struct {
	Host interface{} `yaml:"host"`
	Port int         `yaml:"port"`
}

//type Config struct {
//	Mysql Mysql `json:"mysql"`
//	Redis Redis `json:"redis"`
//}
//
//type Mysql struct {
//	Url  string
//	Port int
//}
//
//type Redis struct {
//	Host string
//	Port int
//}

func main() {
	dataBytes, err := os.ReadFile("study7_config_analysis/y2.yaml")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(dataBytes))

	config := Config{}
	err = yaml.Unmarshal(dataBytes, &config)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(config)
	fmt.Println(config.Mysql, config.Redis)
	fmt.Println(config.Mysql.Url, config.Redis.Host, config.Redis.Port)

	fmt.Println("=======================")

	mp := make(map[string]any, 2)
	err = yaml.Unmarshal(dataBytes, mp)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(mp)
}
