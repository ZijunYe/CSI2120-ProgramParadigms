# Go Basic
## commenting, import, package
1. comments use "//" for one line comments
2. Section can be commented out with "/*"
3. no special character to end a line
4. Exported functions from a package are always capitalized
5. entry point "main" has no arguments and no return
6. ```
      package main //default package
      import "fmt" // import of a package
    ```

example check: Basic.go

## Variable
How to define variable?
1. two line define
```
var a int // make a a default value 0 and initialized
a = 90
```

2. define one line
```
var b int = 10 //make b an initialized
```

3. auto determine, but not recommended
```
var c = 70
```

4. it will not "d:= "second time
```
   d := 90
   d := 100 //this will give error
   name, d := "zijun", 100 //this works because left hand side has uninitialized value

```

5. can group variable definitions, most often used with globals
```
var{
  x float64 = 8.8
  b int
}
```
example check: Variable.go



## Functions
- type of arguments and returns are at the end
- multiple arguments and multiple returns
- General syntax:

```
func functionName (parameter_list) (return_value_list){
  function-body
  return [return_values]
}
```
  - parameter_list --> (param1 type1, param2 type2, ..)
  - return_value_list --> (ret1 type1, ret2 type2, ..)
  - return_value --> either can be specify if unnamed or can be associated by name
  - function that associated to the object, it first is input then it is the function name
example check:
DefiningFunction.go

ifStatement_with_Function.go

Lambda_with_Function.go

Passing_Function_asParameters.go

## Pointers
 - a pointer is the address of that in memory
 - "&" operator returns the **address** of the variable or the function
 - "*" operator returns the data at the address (deference)
 - if the parameter is *pointer, the  it have to passing reference(address)
check Example:

Pointers.go

Pointers_and_structures.go


## If statement, for loop

- if statement can start with a short statement 
```go 
if v := math.Pow(x, n); v < lim {
		return v
	}
```

- for loop the init and post statement are optional
  ```go
      func main() {
      sum := 1
      for ; sum < 1000; {
        sum += sum
      }
      fmt.Println(sum)
    }
  ```

- format of for
```go
  func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}
```

- go's while loop 
```go 
sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
```

- range and loop 
```go 
//range can be treat as index, you can skip the index or value by assign to _ 
func main() {
	pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i) // == 2**i
	}
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}
```

```go 
for i, _:= range pow
for _, value := range pow
for i :=range pow //only want the index, you can omit the second variable
```
## switch
```go
  func main() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
}
```

## DataType
Golang has three categories of data tyoes
1. elementary or primitive
    - (int, float,bool,string)
2. Structured or composite
    - (struct,array, slice, map, channel)
3. Interfaces (only describe the behavior of a type)
