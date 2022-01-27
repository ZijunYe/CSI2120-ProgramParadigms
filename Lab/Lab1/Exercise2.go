//go Basic:Looping

/*
A program should print a pattern as follows to the console:


    x
   xx
  xxx
 xxxx
xxxxx
 xxxx
  xxx
   xx
    x
Add two loops to the following program. The first loop should increase the "x" in lineSymb until the number of "x" is equal to
lineWidth
. The second loop should reduce the number of "x" until only one "x" is left.


*/

package main

import "fmt"

func main() {
	lineWidth := 5
	symb := "x"
	lineSymb := symb
	formatStr := fmt.Sprintf("%%%ds\n", lineWidth)

	//if the start is not reach the max
	//lineSymb += symb means ****
	for ; len(lineSymb) < lineWidth; lineSymb += symb {
		fmt.Printf(formatStr, lineSymb)
	}

	for ; len(lineSymb) > 0; lineSymb = lineSymb[:len(lineSymb)-1] {
		fmt.Printf(formatStr, lineSymb)
	}

}
