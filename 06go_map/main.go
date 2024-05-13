package main

import "fmt"

func main() {
	myMap := make(map[int]string)

	myMap[1] = "value_one"
	myMap[5] = "value_five"
	myMap[8] = "value_eight"
	myMap[1] = "value_new_one"
	myMap[2] = "value_two"

	fmt.Println(myMap)
	fmt.Println("myMap[1] ",myMap[1])
	fmt.Println("myMap[8] ",myMap[8])

	delete(myMap,1)
	fmt.Println("After delete")
	fmt.Println(myMap)
	fmt.Println("myMap[1] ",myMap[1])



	for key,value := range(myMap){
		fmt.Printf("for key %v, value is %v \n",key,value)
	}
}