package main

import "fmt"
func Today(channel chan bool){
	fmt.Println("Wednesday")
	channel<-true
}
func main() {
	channel := make(chan bool)
	go Today(channel)
	<-channel
	fmt.Println("Main function")
}
