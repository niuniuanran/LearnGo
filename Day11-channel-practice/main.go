package main

import "log"

func main() {
	channel := make(chan string)
	defer close(channel)
	go printMessage(channel)
	channel <- "anran"
	channel <- "you"
	channel <- "did"
	channel <- "it"
	channel <- "end"
}

func printMessage(channel chan string) {
	for {
		message := <-channel
		if message == "end" {
			return
		}
		log.Println(message)
	}
}
