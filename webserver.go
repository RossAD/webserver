package main

import (
	"github.com/joho/godotenv"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

// Get path to template directory from env variable
var templatesDir = os.Getenv("TEMPLATES_DIR")

// lovemountaines renders the love-mountains page after passing some data to the HTML
func loveMountains(w http.ResponseWriter, r *http.Request) {
	// Build path to template
	tmplPath := filepath.Join(templatesDir, "love-mountains.html")
	// Load template from disk
	tmpl := template.Must(template.ParseFiles(tmplPath))
	// Inject data into template
	data := "Cadallac Mountain"
	tmpl.Execute(w, data)
}

func main() {
	// Load env variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// Create route to love-mountains web page
	http.HandleFunc("/love-mountains", loveMountains)
	// Launch web server on port 80
	http.ListenAndServe(":80", nil)
}
