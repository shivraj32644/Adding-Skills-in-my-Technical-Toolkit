var variable_name data_type = data;

## go arrays

syntax
    first using var
    var arr = [3] int {1,2,3}
	var arr2 = [...] int {1,2,3,4,5,6,7}

    second using :=

    arr3 := [2] int {1,2}
	arr4 := [...] int {1,2}


Array Initialization
If an array or one of its elements has not been initialized in the code, it is assigned the default value of its type.

Tip: The default value for int is 0, and the default value for string is "".

    arr1 := [5]int{} //not initialized   ----> [0 0 0 0 0]
    arr2 := [5]int{1,2} //partially initialized  ----> [1 2 0 0 0]
    arr3 := [5]int{1,2,3,4,5} //fully initialized ----> [1 2 3 4 5]


## Initialize Only Specific Elements
   It is possible to initialize only specific elements in an array.

   arr1 := [5]int{1:10,4:40} [0 10 40 0 0]

   The array above has 5 elements.

        1:10 means: assign 10 to array index 1 (second element).
        4:40 means: assign 40 to array index 4 (third element).


## Find the Length of an Array
The len() function is used to find the length of an array:

```
    arr1 := [4]string{"Volvo", "BMW", "Ford"}
    arr2 := [...]int{1,2,3,4,5,6}

    fmt.Println(len(arr1)) // 4
    fmt.Println(len(arr2)) // 6
```