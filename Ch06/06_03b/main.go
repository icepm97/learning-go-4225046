package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

const url = "http://services.explorecalifornia.org/json/tours.php"

func main() {
	content := readHttpContent()
	// fmt.Print(content)
	tours := toursFromJson(content)
	for _, tour := range tours {
		fmt.Printf("%+v\n", tour)
	}
}

func toursFromJson(content string) []Tour {
	decoder := json.NewDecoder(strings.NewReader(content))
	_, err := decoder.Token()
	checkError(err)

	tours := make([]Tour, 0)
	var tour Tour
	for decoder.More() {
		err = decoder.Decode(&tour)
		checkError(err)
		tours = append(tours, tour)
	}
	return tours
}

type Tour struct {
	Name, Price string
}

func readHttpContent() string {
	fmt.Println("Network requests")
	client := http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	checkError(err)

	req.Header.Set("User-Agent", "")

	resp, err := client.Do(req)
	checkError(err)
	defer resp.Body.Close()

	fmt.Printf("Response type:%T\n", resp)

	bytes, err := io.ReadAll(resp.Body)
	checkError(err)
	return string(bytes)

}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
