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
