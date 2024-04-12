package main

import (
	"fmt"
	"regexp"
	// "strings"
)

type Stack struct {
	data []interface{}
}

func (s *Stack) Push(val interface{}) {
	fmt.Println("[push]:", val)
	s.data = append(s.data, val)
}

func (s *Stack) Pop() (interface{}, bool) {
	if len(s.data) == 0 {
		return nil, false
	}

	lastIndex := len(s.data) - 1
	lastVal := s.data[lastIndex]
	s.data = s.data[:lastIndex]

	fmt.Println("[pop]: ", lastVal)

	return lastVal, true
}

func (s *Stack) Peek() (interface{}, bool) {
	if len(s.data) == 0 {
		return nil, false
	}
	return s.data[len(s.data)-1], true
}

func GetPrecedence(operator string) int {
	precedences := map[string]int{
		"+": 1,
		"-": 1,
		"*": 2,
		"/": 2,
		"^": 3,
	}
	return precedences[operator]
}

func Tokenize(expression string) []string {
	pattern := `\d+|\+|-|\*|/|\(|\)|\^`
	token_regex, err := regexp.Compile(pattern)

	if err != nil {
		return []string{"ERROR"}
	}

	return token_regex.FindAllString(expression, -1)
}

func IsNumber(str string) bool {
	match, err := regexp.MatchString("^-?[0-9]+(?:\\.[0-9]*)?$", str)
	if err != nil {
		return false
	}
  return match
}

func InfixToPostfix(expression string) string {
	output := &Stack{}
	operators := &Stack{}
	tokens := Tokenize(expression)

	for _, token := range(tokens) {
		fmt.Println("Token:", token)

		// check if token is a valid number and add to stack
		if IsNumber(token) {
			fmt.Println("[number]:", token)
			output.Push(token)
		} else if token == "(" {
			fmt.Println("[operator]:", token)
			operators.Push(token)
		} else if token == ")" {
			fmt.Println("[operator]:", token)
			for operators.data[len(operators.data)] != "(" {
				lastVal, _ := operators.Pop()
				output.Push(lastVal)
			}
		} else {
			fmt.Println("[token]:", token)
			fmt.Println("[stack->length]:", len(operators.data))

			// valid operator order precedence
			if len(operators.data) != 0 {
				for GetPrecedence(operators.data[len(operators.data)-1].(string)) >= GetPrecedence(token) {
					lastVal, _ := operators.Pop()
					output.Push(lastVal)
				}	
			}
			
			operators.Push(token)
		}
	}

	for len(operators.data) >= 0 {
		lastVal, _ := operators.Pop()
		output.Push(lastVal)
	}

	fmt.Println(output.data)

	return ""
	// return strings.Join(output.data, " ")
}

func main() {
	tokens := InfixToPostfix("2 + 3")

	// output := &Stack{}
	// operators := &Stack{}

	// output.Push(2)
	// operators.Push("+")
	// output.Push(3)

	// fmt.Println(output.data)
	// fmt.Println(operators.data)


	fmt.Println(tokens)

}