package main

import (
	"encoding/json"
	"fmt"

	"gopkg.in/yaml.v3"
)

func main() {
	//YamlTest()
	fmt.Println("json")
	info := Info{Id: "110", Age: 15}
	jsonstr, _ := json.Marshal(&info)
	fmt.Println(string(jsonstr))

	fmt.Println("\nyaml")
	yamlstr, _ := yaml.Marshal(&info)
	fmt.Println(string(yamlstr))
}

type Info struct {
	Id   string `json:"id" yaml:"id"`
	Name string
	Age  int    `json:"age" yaml:"age,omitempty"`
	Num  string `json:"num,omitempty" yaml:"-"`
}
