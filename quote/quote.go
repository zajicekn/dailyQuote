package quote

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

// Create a struct to hold the JSON data from the API request. The tags
// 'q' and 'a' match the tags of the API used
type quoteInfo struct {
	Quote  string `json:"q"`
	Author string `json:"a"`
}

func Quote() string {
	// Send a request to retrieve a random quote from the API
	// endpoint. Make sure to check for errors and close the response body
	resp, err := http.Get("https://zenquotes.io/api/random")
	if err != nil {
		fmt.Println("Error fetching the quote:", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	// Check the status code of the response. Check to see the response is not OK
	if resp.StatusCode != 200 {
		fmt.Printf("Error: Received response with status code %d: %s\n", resp.StatusCode, resp.Status)
		os.Exit(1)
	}

	// Read the response data and check if there are any errors in the response reading
	responseData, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading the response data:", err)
		os.Exit(1)
	}

	// Create a variable to hold the information about the random quote. Use the built struct
	var quoteJson []quoteInfo
	err = json.Unmarshal(responseData, &quoteJson)
	if err != nil {
		fmt.Println("Error unmarshalling the JSON data:", err)
		os.Exit(1)
	}

	message := ""
	now := time.Now()
	for _, quote := range quoteJson {
		message += fmt.Sprintf("%v\n\nQuote:\n%s\n\nAuthor:\n%s\n", now.Format("Monday, 02 January 2006"), quote.Quote, quote.Author)
	}
	return message
}
