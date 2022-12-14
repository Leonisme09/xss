package main

import (
    "fmt"
    "net/http"
    "regexp"
)

// This function sends a GET request to the given URL and returns the response body
func getResponseBody(url string) string {
    resp, err := http.Get(url)
    if err != nil {
        fmt.Println("Error sending request:", err)
        return ""
    }
    defer resp.Body.Close()

    // Read the response body and return it
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Println("Error reading response body:", err)
        return ""
    }
    return string(body)
}

// This function checks the given URL for potential XSS vulnerabilities
func checkForXss(url string) bool {
    // Send a GET request to the URL and get the response body
    body := getResponseBody(url)
    if body == "" {
        return false
    }

    // Use a regular expression to search the response body for potential XSS vulnerabilities
    xssRegex := regexp.MustCompile(`<script>.*</script>`)
    if xssRegex.MatchString(body) {
        // If the regular expression matches, return true
        return true
    }

    // If no XSS vulnerabilities were found, return false
    return false
}

func main() {
    // Set the URL to scan
    url := "https://www.example.com"

    // Check the URL for potential XSS vulnerabilities
    if checkForXss(url) {
        fmt.Println("XSS vulnerabilities found!")
    } else {
        fmt.Println("No XSS vulnerabilities found.")
    }
}
