package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Stu1 struct {
	Name  string `json:"name"`
	Age   int
	HIgh  bool
	sex   string
	Class *Class `json:"class"`
}

type Stu struct {
	Name  string `json:"name"`
	Age   int
	HIgh  bool
	sex   string
	Class Class `json:"class"`
}

type Class struct {
	Name  string
	Grade int
}

func main() {
	//实例化一个数据结构，用于生成json字符串
	stu := Stu{
		Name: "张三",
		Age:  18,
		HIgh: true,
		sex:  "男",
	}

	cla := new(Class)
	cla.Name = "1班"
	cla.Grade = 3
	stu.Class = *cla //如果结构体中定义为*Class（指针类型的），则直接赋值：stu.Class=cla; 否则，需要注意赋值类型

	jsonStu, err := json.Marshal(stu)
	if err != nil {
		fmt.Println("json err:", err)
	}

	fmt.Println(string(jsonStu))
	file, err := os.Create("src/study7_config_analysis/test.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	file.Write(jsonStu)
	//new_stu := new(Stu)
	new_stu := Stu{}
	_ = json.Unmarshal(jsonStu, &new_stu)
	//_ = json.Unmarshal([]byte(jsonStu), &new_stu)
	fmt.Println(new_stu.Class.Name)

}
