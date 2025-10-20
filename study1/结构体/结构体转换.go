package main

import "fmt"

type Student struct {
	Age int
}

type Person struct {
	Age int
}

//func calc() {
//	fmt.Println()
//	var s Student = Student{Age: 19}
//	var p Person = Person{Age: 33}
//	s = Student(p)
//	fmt.Println(s)
//	fmt.Println(s.Age)
//}

type Stu Student

func main() {
	fmt.Println()
	var s1 Student
	s1.Age = 100
	var s2 Stu = Stu{99}
	s3 := Stu{Age: 18}
	fmt.Println(s1, s2, s3)

	fmt.Printf("%T\n", s1)
	fmt.Printf("%T\n", s2)

	//s1 = Student(s2)
	//fmt.Printf("%v  %T\n", s1.Age, s1)
	//fmt.Printf("%v  %T\n", s2.Age, s2)

	fmt.Printf("%v  %T\n", s2.Age, s2)
	fmt.Printf("%v  %T\n", s1.Age, s1)

	s2 = Stu(s1)
	fmt.Printf("%v  %T\n", s2.Age, s2)
	fmt.Printf("%v  %T\n", s1.Age, s1)
}
