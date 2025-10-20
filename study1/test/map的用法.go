package main

import (
	"fmt"
)

func main() {
	//方式1
	//只声明map时并未分配内存空间，需要等make时才分配
	var map1 map[int]string
	map1 = make(map[int]string, 10)
	map1[202401] = "张三1"
	map1[202402] = "张三2"
	map1[202408] = "'3"
	map1[202403] = "张三3"
	map1[202401] = "李四"
	map1[202405] = "'3"
	fmt.Print(map1)
	fmt.Println()
	//map是无序的

	//方式2
	b := make(map[int]string)
	b[1] = "100"
	b[2] = "200"
	b[3] = "300"
	fmt.Println(b)
	fmt.Println(b[2])
	for k, v := range b {
		fmt.Println(k, v)
	}
	d := make(map[int]interface{})
	d[1] = 100
	d[2] = "sq1s"
	fmt.Println(d)
	fmt.Println("*************")
	f := make(map[interface{}]interface{})
	f["xx"] = 100
	f[99] = 10.199
	fmt.Println(f)
	fmt.Println("%%%%%%%%%%%%")

	//方式3
	c := map[int]string{
		1: "999",
		2: `"100"`,
		3: "200",
	}
	fmt.Println(c)
	fmt.Println(c[2])
	fmt.Printf("xx:%v\n", c[0])
	fmt.Println("-------------")
	fmt.Printf("xx:%q\n", c[3])
	delete(c, 3)
	fmt.Println("3:", c[3])
	fmt.Printf("xx:%q\n", c[3])

	i, j := c[2]
	fmt.Println(i, j)
	k, l := c[3]
	fmt.Println(k, l)
	fmt.Println(l)
	//map的长度
	fmt.Println(len(c))
	//遍历，只有for range
	for k, v := range c {
		fmt.Println(k, v)
	}

	//map中的value为map情况实列
	a1 := make(map[string]map[int]string)
	a1["班级1"] = make(map[int]string)
	a1["班级1"][20001] = "夏鸥"
	a1["班级1"][20002] = "秋菊"
	a1["班级1"][20003] = "冻双"

	a1["班级2"] = make(map[int]string)
	a1["班级2"][50001] = "桑丘"
	a1["班级2"][50002] = "宙斯"
	a1["班级2"][50003] = "咕噜"

	fmt.Println(a1)
	fmt.Println(a1["班级1"])
	for k1, v1 := range a1 {
		fmt.Println(k1, v1)
		for k2, v2 := range v1 {
			fmt.Println(k2, v2)
		}
	}
}
