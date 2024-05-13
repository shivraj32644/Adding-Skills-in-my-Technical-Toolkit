# Go Slices

Slices are similar to arrays, but are more powerful and flexible.

Like arrays, slices are also used to store multiple values of the same type in a single variable.

However, unlike arrays, the length of a slice can grow and shrink as you see fit.

- `len()` function - returns the length of the slice (the number of elements in the slice)
- `cap()` function - returns the capacity of the slice (the number of elements the slice can grow or shrink to)


In Go, there are several ways to create a slice:

- Using the []datatype{values} format
```
    myslice := []int{}
    sliceOne := []int{1, 2, 3, 4}
	fmt.Println(sliceOne)
```


- Create a slice from an array
```
    var myarray = [length]datatype{values} // An array
    myslice := myarray[start:end]


```


- Using the make() function

```
    slice_name := make([]type, length, capacity)
    
```