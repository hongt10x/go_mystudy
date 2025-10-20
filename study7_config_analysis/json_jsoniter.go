package main

import (
	"fmt"

	jsoniter "github.com/json-iterator/go"
)

// Product _
type Product struct {
	Name      string  `json:"name"`
	ProductID int64   `json:"product_id,string"`
	Number    int     `json:"number,string"`
	Price     float64 `json:"price,string"`
	IsOnSale  bool    `json:"is_on_sale,string"`
	Test      string  `json:"-"`
	OMTest    string  `json:"om_test,omitempty"`
}

var cjson = jsoniter.ConfigCompatibleWithStandardLibrary

func main() {
	p := &Product{
		Name:      "test",
		ProductID: 01,
		Number:    110011,
		Price:     0.01,
		IsOnSale:  true,
		Test:      "test",
	}
	jsonP, v := cjson.Marshal(p)
	fmt.Println(string(jsonP), v)

	p1 := &Product{}
	_ = cjson.Unmarshal(jsonP, p1)
	fmt.Println(*p1)
}

/*json:"-"标签的含义是，当这个结构体被序列化为JSON时，Test字段会被忽略，不会包含在生成的JSON字符串中。同样地，当从JSON反序列化回这个结构体时，Test字段也不会被填充，即使JSON字符串中包含对应的数据。

 */

//
//func main() {
//	p := &Product{
//		Name:      "test",
//		ProductID: 01,
//		Number:    110011,
//		Price:     0.01,
//		IsOnSale:  true,
//		Test:      "test",
//	}
//	jsonP, v := json.Marshal(p)
//	fmt.Println(string(jsonP), v)
//}
//
///*json:"-"标签的含义是，当这个结构体被序列化为JSON时，Test字段会被忽略，不会包含在生成的JSON字符串中。同样地，当从JSON反序列化回这个结构体时，Test字段也不会被填充，即使JSON字符串中包含对应的数据。
//
// */
