package main

import (
	"fmt"
	"strings"
)

type stack struct {
	size int
	top  int
	data []string
}

func main() {

	string := "()"

	output := isValid(string)
	fmt.Println(output)

}

func isValid(str string) bool {
	s := stack{}

	arr := strings.Split(str, "")
	s.size = len(arr)
	for i := 0; i < len(arr); i++ {
		s.push(arr[i])
	}
	fmt.Println(s.data)
	if s.top == 0 {
		return true
	} else {
		return false
	}

}

func (s *stack) isFull() bool {
	if s.top == s.size {
		return true
	} else {
		return false
	}
}

func (s *stack) isEmpty() bool {
	if s.top == 0 {
		return true
	} else {
		return false
	}
}

func (s *stack) push(x string) {
	if !s.isFull() {
		fmt.Println("here")

		if !s.isEmpty() {
			if x == "]" && s.data[s.top-1] == "[" || x == ")" && s.data[s.top-1] == "(" || x == "}" && s.data[s.top-1] == "{" {

				s.pop()
			} else {

				s.data = append(s.data, x)
				s.top++
			}
		} else {

			s.data = append(s.data, x)
			s.top++
		}

	}
	fmt.Println(s.data)
}

func (s *stack) pop() {
	if !s.isEmpty() {
		s.data = append(s.data[:s.top-1], s.data[s.top-1+1:]...)
		s.top--
	}
}
