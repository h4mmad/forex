package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"
)

// ApiResponse represents the overall structure of the response.


func makeGETrequest() CurrencyData{
    resp, err := http.Get(createRequestURL())
    if err != nil {
        log.Fatal("Error making GET request", err)
    }

    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)

    if err != nil{
        fmt.Println("Error reading the response body:", err)
    }

    var target CurrencyData

    unmarshalErr := json.Unmarshal(body, &target)

    if unmarshalErr != nil{
        log.Fatalf("Error unmarshaling JSON: %v", err)
    }

    target.Meta.LastCheckedAt = time.Now().UTC().Format(time.RFC3339)
    fmt.Println("API call made successfully")

    return target
}


// createRequestURL generates a full URL for making an API request to convert currencies.
//
// Parameters:
//   - targetCurrency (CurrencyCode): The target CurrencyCode code for the conversion (e.g., "INR", "SAR").
//   - baseURL (string): The base URL of the API endpoint to make the request (e.g., "https://api.currencyapi.com/v3/latest").
//   - apiKey (string): The API key required to authenticate the request.
//
// Returns:
//   - string: A fully-formed URL with the provided parameters, which can be used to make a GET request.
//
// Example:
//   targetCurrency := "INR"
//   baseURL := "https://api.currencyapi.com/v3/latest"
//   apiKey := "your_api_key"
//   fullURL := createRequestURL(targetCurrency, baseURL, apiKey)
//   // fullURL will contain something like:
//   // "https://api.currencyapi.com/v3/latest?apikey=your_api_key&base_currency=USD&currencies=INR"
//
// Details:
//   - The function uses the `url.Values{}` to build query parameters in a URL-encoded format.
//   - The base CurrencyCode is always set to "USD", as specified in the function.
//   - The `fmt.Sprintf` method is used to combine the base URL and the encoded query parameters into the final URL string.
func createRequestURL() string{
    baseURL:="https://api.currencyapi.com/v3/latest"
    apiKey:="cur_live_jjpfSTEGgby2Cs95zkK2MVXzdBJqHjnMNc3EB9Jn" 
    params := url.Values{}
    params.Add("apikey", apiKey)
    fullURL := fmt.Sprintf("%s?%s", baseURL, params.Encode())
    return fullURL
}