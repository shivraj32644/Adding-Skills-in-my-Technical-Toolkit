package main
import "fmt"


func main()  {
	fmt.Println("hello")
	var username string = "shivraj"
	fmt.Print(username)
	var arr = [3] int {1,2,3}
	var arr2 = [...] int {1,2,3,4,5,6,7}
	fmt.Println(arr)
	fmt.Println(arr2)
	arr3 := [2] int {1,2}
	arr4 := [...] int {1,2}
	fmt.Println(arr3)
	fmt.Println(arr4)


	arr5 := [5]int{} //not initialized
  	arr6 := [5]int{1,2} //partially initialized
  	arr7 := [5]int{1,2,3,4,5} //fully initialized
	arr8 := [8] int {1:20,4:50}
  	fmt.Println(arr5)
  	fmt.Println(arr6)
  	fmt.Println(arr7)
  	fmt.Println(arr8)

	arr9 := [4]string{"Volvo", "BMW", "Ford"}
	fmt.Println(len(arr9))

}