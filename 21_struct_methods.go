package main

import "fmt"

type Student struct {
	name  string
	marks [5]int
	age   int
}

func (s Student) getAge() int {
	return s.age
}

func (s *Student) setAge(age int) {
	s.age = age
}

func (s Student) getaveragemarks() int {
	sum := 0

	for i := 0; i < len(s.marks); i++ {
		sum += s.marks[i]
	}
	return sum / 5
}
func main() {
	var student Student = Student{"Santhosh", [5]int{7, 7, 7, 7, 7}, 18}
	fmt.Println(student)
	/*
		a := student.getAge()
		fmt.Println(a)
		student.setAge(19)
		fmt.Println(student)*/
	var avg int
	avg = student.getaveragemarks()
	fmt.Println(avg)

}
