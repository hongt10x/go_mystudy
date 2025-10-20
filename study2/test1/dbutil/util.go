package dbutil

import "fmt"

var SCHOOL = "人民大学"

func GetCount() string {
	return fmt.Sprintf("我是util.go中的方法，我有个属性 SCHOOL=%v", SCHOOL)
}
