# Structured or composite datatype

## Struct
Definition: a *struct* is a collection of fields 
```go
type Vertex struct {
	X int
	Y int
}

func main() {
	fmt.Println(Vertex{1, 2})
}
```
- How to access struct field? -- accessed using a dot 
```go 
type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	v.X = 4 //to access the field X in struct Vertex 
	fmt.Println(v.X)
}
```

## Array
Definition: arrays's length cannot be resized 
- How to declare Array? 
```go 
//two way to declare array
//1. with size 
var table = [5] int{1,2,3,4,5}
//2. no specific size 
var table = [...] int{1,2,3,4,5} 
```
## Slice
Definition: array has a fixed size. slice is dynamically-sized,flexible view into the elements of array. (arraylist) 
- How to creating a slice? 
```go 
//1way
var slice [] int = array[start :end] 
//2ways 
slice := make([] int, 10,100) //underline will create a array
- A slice form of two indices, a low and high
```go 
a[low:high] 
//example 
a[1:4] // includes elements 1 through 3 
```
- length and capacity
  -  length: number of elements it contains ```len(s)```
  -  capacity: number of element in underlying array,max ```cap(s)```

- Appending to a slice 
  - append new elements to slice, using built-in ```append``` function 
  - ```func append( s[] T, ....T) ``` first parameter is slice name, rest is new element need to append 
## map
Definition: maps keys to value 
```go 
type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
}
```
## channel 
