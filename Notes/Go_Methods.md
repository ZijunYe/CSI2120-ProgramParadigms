# Stream I/O 
1. Console I/O
    - console I/O with package fmt.
    - C-like syntax with scanf and printf in addition to print(println)
    ```go
    package main
    import "fmt"

    func main(){
        var inStr string
        fmt.Printf("Input?")
        fmt.Scanf("%s", &inStr) //take the input from user as reference, inStr is the value
        fmt.Printf("\nOutput: %s\n" ,inStr)
    }
    ```
2. File I/O
    - File I/O follows familar notion of stream I/O
    - relavant packages are os and bufio besides fmt
    ```go
    func fCopy(outN, inN string)(bytesOut int64, err error){
        inF,err:= os.Open(inN) //we open a file for inN
        if err != nil{
            return
        }
        //nil is similar to null that points to nothing
        outF,err := os.Create(outN)//we create a file for outF, for write
        if err != nil{
            inF.close()
            return
        }
        bytesOut, err = io.Copy(outF,inF) //it will return how many bit that the file use
        inF.close(); outF.Close()
        return
    }
    ```


# Final evaluation with defer
- handling stuff that need to be done at the end
- delay the function call before the return(till the end)
  ```go
    func fCopy(outN, inN string)(bytesOut int64, err error){
        inF,err:= os.Open(inN) //we open a file for inN
        if err != nil{
            return
        }
        defer inF.Close()

        outF,err := os.Create(outN)
        if err != nil{
            return
        }
        defer outF.Close()

        bytesOut, err = io.Copy(outF,inF) //it will return how many bit that the file use
        return
    }
    ```
# Errors and Panic
- there is no exception in go, so instead of exception, return error codes of type error
- when there is a servious error occurs, use panic statement
- the features of panic:
    1. function stops immediately
    2. deferred function are executed
    3. stack unwinding occurs
        - return to the calling function
        - start execute deferred functions until main exits
        - can be stopped with recover
    ```go
    package main
    import "fmt"
    func main(){
        defer fmt.Println("Last output")
        fmt.Println("Before Panic")
        panic("Out of here")
        fmt.Println("After panic") //never execute
    }
    ```
  - **Recover** is the statement to regain control after panic 
             - to be used in a defeered functon 
             - In the absence of panic, the recover simply return nil 
    
# Method and Receiver 
- Definition of Method: is a function acts on a cetertain type or a structure 
- Definition of Receiver: can be a pointer to an allowed type 
example: 
```go 
type Point struct{
    x.float64 
    y.float64
}
//(pt *Point) is receiver 
//norm is the method of Point 
func (pt *Point) norm() float64{
    return math.Sqrt( pt.x*pt.x + pt.y*pt.y)
}

func main(){
    a := Point{2,4}
    n := a.norm()
}
```

# Interface and Polymorphism 
- Definition: interface define a set of method, no implementation of the method 
- name convention: interface name should end in "er" 
- need to use ```type name_of_interface interface{}```
Example: 
```go 
type ColorPt struct{
    pt Point 
    color string 
}

type Box struct{
    weight float64 
    color string 
}

//interface to be implemented by ColorPt and Box 
type Color interface{
    setColor(String)//method1
    Color() string //method2 
}


//interface implementation for type ColorPt 
func (p *ColorPt) color() string{return p.color}
func (p *ColorPt) setColor(col string){p.color = col}
//interface implementation for type Box
func (p *Box) color() string{return p.color}
func (p *Box) setColor(col string){p.color = col}

func main(){
    //array with interface type 
    table := [...] Color{
        &ColorPt{Point{1.1,2.2},"red"}
        &Box{32.4,"yellow"}
        &ColorPt{Point{3.1,3.2},"blue"}
    }
   for _,element := range table{
        fmt.Printf("color =%s\n", element.Color()
    }
  }
```
# Binary tree and Generic Queue 

```go 
// tree
package main

import (
	"fmt"
)

type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

func main() {
	t := NewTree(5)
	t.Insert(7)
	t.Insert(9)
	t.Insert(2)
	t.InOrder()

}

func NewTree(v int) *Tree {

	return &Tree{nil, v, nil}
}

func (t *Tree) InOrder() {

	if t.Left != nil {
		t.Left.InOrder()
	}
	fmt.Println(t.Value)

	if t.Right != nil {
		t.Right.InOrder()
	}
}
func (t *Tree) Insert(v int) *Tree {

	if v < t.Value {
		if t.Left == nil {
			t.Left = NewTree(v)
			return t.Left
		} else {
			return t.Left.Insert(v)
		}
	} else {
		if t.Right == nil {
			t.Right = NewTree(v)
			return t.Right
		} else {
			return t.Right.Insert(v)
		}
	}
	return t
}
```
## Generic Go queue 
- use the empty interface, that interface can hold any value 

```go
// queue
package main

import (
	"fmt"
)

type Queue struct {
	items []interface{} //empty interface 
}

// Enqueue adds a value to the
// end of the queue
func (s *Queue) Enqueue(t interface{}) {
	s.items = append(s.items, t)
}

// Dequeue removes a value from the
// start of the queue
func (s *Queue) Dequeue() interface{} {
	item := s.items[0]
	s.items = s.items[1:len(s.items)]
	return item
}

// Front returns the item next in
// the queue, without removing it
func (s *Queue) Front() *interface{} {
	item := s.items[0]
	return &item
}

// IsEmpty returns true if the
// queue is empty
func (s *Queue) IsEmpty() bool {
	return len(s.items) == 0
}

// Size returns the number of elements
func (s *Queue) Size() int {
	return len(s.items)
}
func main() {

	var queue Queue

	queue.Enqueue(10)
	queue.Enqueue(20)
	queue.Enqueue(30)

	t := queue.Dequeue()

	fmt.Println("value=", t)

	queue.Enqueue(40)
	queue.Enqueue(50)
	queue.Enqueue(60)

	t = queue.Dequeue()
	t = queue.Dequeue()
	t = queue.Dequeue()
	fmt.Println("value(2)=", t)
	fmt.Println("size=", queue.Size())
	fmt.Println("front=", *queue.Front())
}

```











