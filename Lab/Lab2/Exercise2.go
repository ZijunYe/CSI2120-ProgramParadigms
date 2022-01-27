package main

import (
    "fmt"
    "time"
    "strconv"
)

var i int

func makeCakeAndSend(cs chan string) {
    i = i + 1
    cakeName := "Strawberry Cake " + strconv.Itoa(i)
    fmt.Println("Making a cake and sending ...", cakeName)
    cs <- cakeName //send a strawberry cake

    //cs is channel
    //DEFAULT channel can contains 1

}

func receiveCakeAndPack(cs chan string) {
    s := <-cs //get whatever cake is on the channel
    //what is <- mean?

    //if channel is not value(empty)
    //it will wait till chennel have value
    //the line that receive the channel, it will block if channel is empty
    fmt.Println("Packing received cake: ", s)
}

func main() {
    cs := make(chan string)
    for i := 0; i<3; i++ {
        go makeCakeAndSend(cs)
        go receiveCakeAndPack(cs)

        //sleep for a while so that the program doesn't exit
        //immediately and output is clear for illustration
        time.Sleep(1 * 1e9) 
    }
}
