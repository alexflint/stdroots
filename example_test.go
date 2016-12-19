package stdroots

import (
	"fmt"
	"log"
)

func ExampleClient() {
	resp, err := Client.Get("https://www.google.com")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp.Status)
	// Output: 200 OK
}
