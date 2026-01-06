package fetchpkg 

import (
	"fmt"
	"io"
	"log"
	"net/http"
)	


func Fetch(varName string, c chan string) {

	fmt.Println("fetch called")

	resp, err := http.Get(varName)
	if err != nil {
		log.Fatal("Error making GET: ", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		//String print formatted 
		c <- fmt.Sprintf("error %d for URL %s", resp.StatusCode, varName)

	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("error reading body: ", err)
	}

	c <- string(body)	
}



