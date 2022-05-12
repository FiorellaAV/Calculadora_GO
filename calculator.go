package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type calc struct {
}

//reciver function (calc ), is to assing methods/functions to a struct
func (calc) operate(input string, operator string) int {

	values := strings.Split(input, operator)

	firstValue := parse(values[0])
	secondValue := parse(values[1])

	switch operator {

	case "+":
		result := firstValue + secondValue
		fmt.Println(result)
		return result

	case "-":
		result := firstValue - secondValue
		fmt.Println(result)
		return result

	case "*":
		result := firstValue * secondValue
		fmt.Println(result)
		return result

	case "/":
		result := firstValue / secondValue
		fmt.Println(result)
		return result

	default:
		fmt.Println("Invalid operation")
		return 0
	}

}

//function that converts into integer the string
func parse(input string) int {
	value, _ := strconv.Atoi(input)
	return value
}

//function that reads an input and returns it
func readInput() string {
	//we create a scanner to read the input
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	//then we store that value in a variable
	return scanner.Text()
}

func main() {

	input := readInput()
	//the second time we call the readInput() to know the operator
	operator := readInput()

	c := calc{}
	c.operate(input, operator)
}
