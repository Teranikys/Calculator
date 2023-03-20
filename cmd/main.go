package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Enter an arithmetic expression to evaluate it.")
		fmt.Println("Enter 0 to exit the program:")
		input, _ := reader.ReadString('\n')
		//input = strings.Replace(input, " ", "", -1)
		input = strings.TrimSpace(input)
		result := EvalArabic(input)
		fmt.Printf("The result is: %v\n", result)
	}
}

func EvalArabic(input string) int {
	exp := toPolandNotation(input)
	var result []int

	for _, val := range exp {
		switch val {
		case "+":
			x, y := result[len(result)-1], result[len(result)-2]
			result = result[2:]
			result = append(result, x+y)
		case "*":
			x, y := result[len(result)-1], result[len(result)-2]
			result = result[2:]
			result = append(result, x*y)
		case "/":
			x, y := result[len(result)-1], result[len(result)-2]
			result = result[2:]
			result = append(result, x/y)
		case "-":
			x, y := result[len(result)-1], result[len(result)-2]
			result = result[2:]
			result = append(result, x-y)
		default:
			x, _ := strconv.Atoi(val)
			result = append(result, x)
		}
	}

	return result[0]
}

func toPolandNotation(input string) []string {
	var exp []string
	exp = strings.Split(input, " ")
	var stack []string
	var queue []string
	precedence := "+-*/"

	for _, val := range exp {
		// If the incoming element is a number, then we add it to the queue
		if matched, _ := regexp.MatchString(`\d`, val); matched {
			num, _ := strconv.Atoi(val)
			if num < 0 || num > 10 {
				fmt.Println("Invalid input")
			}
			queue = append(queue, val)
		} else
		// If the incoming element is an operator (+, -, *, /) then we check:
		if val == "+" || val == "-" || val == "*" || val == "/" {
			// If the stack is empty or contains a left parenthesis at the top, then add the incoming statement
			//to the stack.
			if len(stack) == 0 || stack[len(stack)-1] == "(" {
				stack = append(stack, val)
			} else
			// If the incoming operator has a higher priority than the top, push it onto the stack.
			if strings.Index(precedence, val) > strings.Index(precedence, stack[len(stack)-1]) {
				stack = append(stack, val)
			} else
			// If the incoming operator has lower or equal precedence than the top, pop the top of the stack
			//onto the queue until we see an operator with lower precedence or a left parenthesis at the top,
			//then push the incoming operator onto the stack.
			{
				for strings.Index(precedence, val) <= strings.Index(precedence, stack[len(stack)-1]) {
					queue = append(queue, stack[len(stack)-1])
					stack = stack[:len(stack)-1]
					if len(stack) == 0 {
						break
					}
				}
				stack = append(stack, val)
			}
		} else
		// If the incoming element is a left parenthesis, push it onto the stack.
		if val == "(" {
			stack = append(stack, val)
		} else
		// If the incoming element is a right parenthesis, pop the stack and add its elements to the queue until we see
		//a left parenthesis. Remove the found bracket from the stack.
		if val == ")" {
			for stack[len(stack)-1] != "(" {
				queue = append(queue, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
				if len(stack) != 0 {
					break
				}
			}
			stack = stack[:len(stack)-1]
		}
	}
	//At the end of the expression pop the stack into the queue.
	for len(stack) != 0 {
		queue = append(queue, stack[len(stack)-1])
		stack = stack[:len(stack)-1]
	}

	return queue
}
