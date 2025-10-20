package main

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Config1 struct {
	Database struct {
		Name string `yaml:"name"`
		Age  int    `yaml:"age"`
		//Address *Class `yaml:"address"`
	} `yaml:"database"`
}

type Addr struct {
	Num int
	Stu string
}

func main() {
	var config Config1
	//	data := `
	//database:
	//  name: admin
	//  age: 33
	//`
	data, err := os.ReadFile("study7_config_analysis\\y1.yaml")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))
	err1 := yaml.Unmarshal([]byte(data), &config)
	if err1 != nil {
		log.Fatal(err1)
	}

	fmt.Println(config.Database.Name, config.Database.Age)
	config.Database.Name = "Echo Wang"
	fmt.Println(config.Database.Name)
}
