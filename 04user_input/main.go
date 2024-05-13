package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strconv"
)

func main() {
	reader :=  bufio.NewReader(os.Stdin)
	// comma ok | error ok syntax
	// try , catch := function()
	fmt.Println("user input")
	input, err := reader.ReadString('\n')

	fmt.Println(input,err)
	fmt.Println("user input 2")
	input1, _ := reader.ReadString('\n')
	fmt.Println(input1)


	stringNum := "20";
	parseToNum,err := strconv.ParseInt(stringNum,10,64)  

	if err != nil {
		fmt.Println(err)
		panic(err)
	}else {
		fmt.Println(reflect.TypeOf(parseToNum))
	}
	fmt.Println(parseToNum)

}