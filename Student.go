package main

import "fmt"

type Student struct {
	Name    string
	ID      string
	Age     int
	Courses []string
}

func StudentConstructor(name string, id string, age int) Student {
	return Student{Name: name, ID: id, Age: age}
}

func (s *Student) Show() {
	fmt.Printf("Hi! My name is %s. I am %d years old.\n", s.Name, s.Age)
}

func (s *Student) EnrollCourses(course string) {
	for _, each_course := range s.Courses {
		if each_course == course {
			return
		}
	}
	s.Courses = append(s.Courses, course)
}

func (s *Student) PrintCourses() {
	fmt.Println("All my courses are:")
	for _, each_course := range s.Courses {
		fmt.Println(each_course)
	}

}

func (s *Student) changeName(newname string) {
	s.Name = newname
}
func (s *Student) changeAge(newage int) {
	s.Age = newage
}
