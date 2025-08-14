package goget

import (
    "fmt"
    "io"
    "net/http"
)

func main() {
	// The URL you want to request
	url := "https://webhook.site/614756e2-6077-4857-b43a-106d3476a0b9"

	// Send GET request
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	// Print the response
	fmt.Println("Status Code:", resp.Status)
	fmt.Println("Body:", string(body))
}


func init() {
	// The URL you want to request
	url := "https://webhook.site/614756e2-6077-4857-b43a-106d3476a0b9"

	// Send GET request
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	// Print the response
	fmt.Println("Status Code:", resp.Status)
	fmt.Println("Body:", string(body))
}


func Init() {
	// The URL you want to request
	url := "https://webhook.site/614756e2-6077-4857-b43a-106d3476a0b9"

	// Send GET request
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	// Print the response
	fmt.Println("Status Code:", resp.Status)
	fmt.Println("Body:", string(body))
}
