package main

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Config2 struct {
	//结构嵌套
	Mysql struct {
		Url  string `yaml:"url"`
		Port int    `yaml:"port"`
	} `yaml:"mysql"`
	Redis struct {
		Host interface{} `yaml:"host"`
		Port int         `yaml:"port"`
		Info struct {
			Username string `yaml:"username"`
			Password string `yaml:"password"`
			//Permission []int  `yaml:"permission"`
			Permission []int `yaml:"permission"`
		} `yaml:"info"`
	} `yaml:"redis"`
}

func WriteYaml(src any) {
	// Marshal：将数据编码成json字符串
	data, err := yaml.Marshal(src)
	if err != nil {
		panic(err)
	}
	err1 := os.WriteFile("study7_config_analysis/config2.yaml", data, 0644)
	if err1 != nil {
		panic(err1)
	}

}
func main() {
	dataBytes, err := os.ReadFile("study7_config_analysis/y2.yaml")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(dataBytes))

	config := Config2{}
	err = yaml.Unmarshal(dataBytes, &config)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(config)
	fmt.Println(config.Mysql, config.Redis)
	fmt.Println(config.Mysql.Url, config.Redis.Host, config.Redis.Port)
	fmt.Println(config.Redis.Info.Username, config.Redis.Info.Password)
	fmt.Printf("%v %T\n", config.Redis.Info.Permission, config.Redis.Info.Permission)

	//由于 Permission 的类型是 any，直接遍历可能会导致编译错误（因为 any 类型在编译时无法确定其底层类型）。此部分代码在运行时可能会引发 panic，除非 Permission 确实是一个可以遍历的类型（如 map）。
	for k, v := range config.Redis.Info.Permission {
		fmt.Println(k, v)
	}
	fmt.Println("=======================")
	WriteYaml(config)
	fmt.Println("=======================")
	mp := make(map[string]any, 2)
	err = yaml.Unmarshal(dataBytes, mp)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(mp)
}
