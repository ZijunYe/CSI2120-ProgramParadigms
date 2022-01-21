#Basic Info
1. comments use "//" for one line comments
2. Section can be commented out with "/*"
3. no special character to end a line
4. Exported functions from a package are always capitalized
5. entry point "main" has no arguments and no return
6. ```package main //default package
      import "fmt" // import of a package```

example check: Basic.go

# Variable
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
``` d := 90
```

5. can group variable definitions, most often used with globals
```
var{
  x float64 = 8.8
  b int
}
```
example check: Variable.go

# Functions
- type of arguments and returns are at the end
- multiple arguments and multiple returns
- General syntax

```
func functionName (parameter_list) (return_value_list){
  function-body
  return [return_values]
}
```
  - parameter_list --> (param1 type1, param2 type2, ..)
  - return_value_list --> (ret1 type1, ret2 type2, ..)
  - return_value --> either can be specify if unnamed or can be associated by name

example check: DefiningFunction.go
