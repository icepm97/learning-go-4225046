package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "http://services.explorecalifornia.org/json/tours.php"

func main() {
	fmt.Println("Network requests")

	client := http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	checkError(err)
	req.Header.Add("User-Agent", "")

	res, err := client.Do(req)
	checkError(err)
	defer res.Body.Close()

	fmt.Printf("Response Type: %T\n", res)

	bytes, err := io.ReadAll(res.Body)
	checkError(err)

	println(string(bytes))
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
