package main

import (
	"fmt"
	"reflect"
)

type MyStruct struct {
	Name string
	Age  int
}

//// 代码示例：从JSON配置文件动态加载设置
//type Config struct {
//	Field1 string `json:"field1"`
//	Field2 int    `json:"field2"`
//}
//
//func LoadConfig(jsonStr string, config *Config) {
//	v := reflect.ValueOf(config).Elem()
//	t := v.Type()
//	// 省略JSON解析步骤
//	// 动态设置字段值
//	for i := 0; i < t.NumField(); i++ {
//		field := t.Field(i)
//		jsonTag := field.Tag.Get("json")
//		// 使用jsonTag从JSON数据中获取相应的值，并设置到结构体字段
//	}
//}
//
//func calc() {
//
//}

func main() {
	var myStruct = MyStruct{"alex", 18}
	fmt.Println(myStruct)
	typeof := reflect.TypeOf(myStruct)
	valueof := reflect.ValueOf(myStruct)
	fmt.Println(typeof, valueof)

}
