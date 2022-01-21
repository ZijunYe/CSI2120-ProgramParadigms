package main

func main() {
	//Different declearation

	//1. two line declear
	var a int //make a a default value 0 and initialized
	a = 90

	//2. declear one line
	var b int = 10 //make b an initialized

	//3. auto determine, but not recommand
	var c = 70

	//4. it will not "d:= "second time
	d := 90

	//5.can group variable definitions
	// most often used with globals
	var (
		x float64 = 8.8
		b int
	)

}
