package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)	


func main() {
	var urlArr [2]string // create string with a length of 2

	// assign to our array
	urlArr[0] = "https://example.com"
	urlArr[1] = "https://example.com/example"
	
	fmt.Println(urlArr)
	
	//create channel.  Creating channel inside of loop will make is synchronous as we would be waiting to assign the response after request
	//so we couldnt loop(makeChannel go makeRequest var <- channelResponse) as the var <- channelResponse would hang until we got the response
	channel := make(chan string)

	for i := 0; i < len(urlArr); i++ {
		go fetch(urlArr[i], channel)
	}
	

	for i := 0; i < len(urlArr); i++ {
		fmt.Println(<-channel)
	}
}


