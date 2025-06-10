package main

import (
	"bufio"
	"fmt"
	"html/template"
	"math/rand"
	"net/http"
	"os"
)

var tmplt *template.Template

type QuoteWeb struct {
	Quote string
}

// Bob Ross Go
func main() {

	// Set up the HTTP server
	http.HandleFunc("/", webHandler)  // Handle requests to the root path
	http.ListenAndServe(":8080", nil) // Start the server on port 8080

}

// Handler function to handle HTTP requests
func webHandler(w http.ResponseWriter, r *http.Request) {
	// Set the content type to plain text
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	//Get a random quote from the quotes file
	quote := getQuote() //Get a random quote

	//Generate bubble border around the quote
	bubblequote := buildBubble(quote)

	tmplt, _ = template.ParseFiles("templates/bob.html")
	data := QuoteWeb{Quote: bubblequote} // Create a Quote struct with the quote

	tmplt.Execute(w, data) // Execute the template with the quote data

}

// Get Quote Function
// Open the file quotes.go and chooses a random line from the file.
func getQuote() string {

	file, err := os.Open("data/quotes.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return "Error: Could not open quotes file."
	}

	defer file.Close() //Ensure the file is closed after reading

	reader := bufio.NewScanner(file)
	var quotes []string
	//Read the file line by line
	for reader.Scan() {
		line := reader.Text()         // Get the current line
		quotes = append(quotes, line) // Append the line to the quotes slice
	}

	return quotes[rand.Intn(len(quotes))] // Return a random quote from the slice
}

// Character repeat function to help with bubble width generation
func repeatChar(char rune, count int) string {
	result := ""
	for i := 0; i < count; i++ {
		result += string(char) // Append the character to the result string
	}
	return result // Return the final string with repeated characters
}

// Build Bubble Function
// Puts a quote inside a bubble and prints it
func buildBubble(quote string) string {

	var quotebubble string
	quotelength := len(quote)

	quotebubble += fmt.Sprintf("%s", "  "+repeatChar('_', quotelength+1)+"\n")
	quotebubble += fmt.Sprintf("%s", " / "+repeatChar(' ', quotelength)+"\\\n")
	quotebubble += fmt.Sprintf("%s", "| "+repeatChar(' ', quotelength+2)+"|\n")
	quotebubble += fmt.Sprintf("%s", "| "+quote+"  |\n")
	quotebubble += fmt.Sprintf("%s", "| "+repeatChar(' ', quotelength+2)+"|\n")
	quotebubble += fmt.Sprintf("%s", " \\"+repeatChar('_', quotelength+1)+"/\n")

	return quotebubble // Return the bubble with the quote inside
}
