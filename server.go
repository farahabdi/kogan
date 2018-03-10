package kogan

import (
	"fmt"
	"encoding/json"
	"net/http"
	"log"
)

type Response struct {
	Objects []Product `json:"objects"`
	Next string `json:"next"`
}

type Product struct {
	Category            string `json:"category"`
	Title               string `json:"title"`
	Weight         		float64 `json:"weight"`
	Sizes				Size   `json:"size"`
}

type Size struct {
	Width               float64 `json:"width"`
	Length              float64 `json:"length"`
	Height        		float64 `json:"height"`
}


{
	products: {
		{
			title: "Red pen"
			price: 2
		},
		{
			title: "Tent"
			price: 12
		}

	},

	next: "/api/products/2"
}

func main() {
	var category []Size

	endpoint := "http://wp8m3he1wt.s3-website-ap-southeast-2.amazonaws.com/api/products/1"
	
	

	req, err := http.NewRequest("GET", endpoint, nil)

	if err != nil {
		log.Fatal("Failed to build request: ", err)
		return
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		return
	}

	defer resp.Body.Close()


	var record Response

	if err := json.NewDecoder(resp.Body).Decode(&record); err != nil {
		log.Println(err)
	}


	for {


	}


fmt.Println("fuck you")
}