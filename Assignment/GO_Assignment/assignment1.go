package main

import (
	"bytes"
	"fmt"
	"strings"
	"sync"
	"unicode"
)

func CaesarCipher(m string, shift int) (result string) {
	//turn string all to uppercase
	m = strings.ToUpper(m)

	//string to char
	//char to integer, then store char
	//char to string

	var c []rune //c stores all the char
	for _, value := range m {
		if unicode.IsLetter(value) {
			c = append(c, value) //add character to slice of unicode
		}
	}
	for i := 0; i < len(c); i++ {
		check := int(c[i]) + shift
		if check > 90 {
			check = 64 + (check - 90)
		}
		c[i] = rune(check)
	}
	var buffer bytes.Buffer
	for _, v := range c {
		buffer.WriteRune(v)
	}
	result = buffer.String()
	return result

}

func CaesarCipherList(src []string, shift int, wg *sync.WaitGroup, result []string) {
	for i, v := range src {
		result[i] = CaesarCipher(v, shift)
	}
	wg.Done()
}

func main() {
	//Exercise1
	fmt.Println("EXERCISE1: ")
	fmt.Printf(CaesarCipher("I love CS!", 5))

	//Exercise2
	fmt.Println()
	fmt.Println()
	fmt.Println("EXERCISE2: ")
	messages := []string{"Csi2520", "Csi2120", "3 Paradigms",
		"Go is 1st", "Prolog is 2nd", "Scheme is 3rd",
		"uottawa.ca", "csi/elg/ceg/seg", "800 King Edward"}

	result := make([]string, len(messages))
	//Create channels
	var wg sync.WaitGroup
	wg.Add(1) //number of threads

	//call go function
	go CaesarCipherList(messages[:], 2, &wg, result[:])

	wg.Wait()
	//print results
	for _, v := range result {
		fmt.Println(v)
	}
	//add synchronisation

	fmt.Println()
	fmt.Println("EXERCISE3: ")

	var wg2 sync.WaitGroup
	wg2.Add(3) //number of threads
	result2 := make([]string, len(messages))

	size := len(messages)

	//split array into three subarray
	if size < 3 {
		fmt.Println("The array impossible to have 3 subarray!")
	} else {
		if size%3 == 0 { //the length of array are able divided by 3
			k := size / 3 //first array
			go CaesarCipherList(messages[:k], 2, &wg2, result2[:k])
			go CaesarCipherList(messages[k:2*k], 2, &wg2, result2[k:2*k])
			go CaesarCipherList(messages[2*k:], 2, &wg2, result2[2*k:])
		} else { //unable to divided by 3
			k := int(size / 3)
			go CaesarCipherList(messages[:k], 2, &wg2, result2[:k])
			go CaesarCipherList(messages[k:2*k], 2, &wg2, result2[k:2*k])
			go CaesarCipherList(messages[2*k:], 2, &wg2, result2[2*k:])
		}

	}

	wg2.Wait()

	for _, v := range result2 {
		fmt.Println(v)
	}

}
