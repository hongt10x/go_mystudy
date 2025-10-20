package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type User struct {
	Name  string `json:"name"`
	Age   int    `json:"age,string"` //序列化时将Age（默认int）强转为string，反序列化时将Age转为int
	Email string `json:"email"`
}

type User1 struct {
	Name  string          `json:"name"`
	Age   json.RawMessage `json:"age,string"` //序列化时将Age（默认int）强转为string，反序列化时将Age转为int
	Email string          `json:"email"`
}

func main() {
	//jsonStr := `{"name":"Echo Wang","age":20,"email":"wht@163.com.cn"}`
	user := User{
		Name:  "john",
		Age:   20,
		Email: "john@gmail.com",
	}

	//序列化
	jsonData, err := json.Marshal(user)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(jsonData)
	jsonStr := string(jsonData)
	fmt.Println(jsonStr)
	fmt.Printf("%T\n", jsonStr)

	//反序列化
	newUser := User{}
	fmt.Println(newUser)
	_ = json.Unmarshal([]byte(jsonStr), &newUser)

	fmt.Println(newUser)
	fmt.Println(newUser.Age)
	fmt.Printf("%T\n", newUser.Age)
}
