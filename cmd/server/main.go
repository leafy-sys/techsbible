package main

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path/filepath"

	"github.com/yuin/goldmark"
)

type PageData struct {
	Content template.HTML
}

func parseMarkdownFile(filename string) (template.HTML, error) {
	path := filepath.Join("content", filename)
	mdData, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	if err := goldmark.Convert(mdData, &buf); err != nil {
		return "", err
	}

	return template.HTML(buf.String()), nil
}

func renderTemplate(w http.ResponseWriter, r *http.Request, pageFile string, data interface{}) {
	if r.Header.Get("HX-Request") == "true" {
		tmpl, err := template.ParseFiles(filepath.Join("views", "pages", pageFile))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		tmpl.ExecuteTemplate(w, "content", data)
		return
	}

	files := []string{
		filepath.Join("views", "layouts", "base.html"),
		filepath.Join("views", "pages", pageFile),
	}
	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.ExecuteTemplate(w, "base.html", data)
}

func main() {
	mux := http.NewServeMux()

	// Static assets
	fs := http.FileServer(http.Dir("static"))
	mux.Handle("GET /static/", http.StripPrefix("/static/", fs))

	// Home page
	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, r, "home.html", nil)
	})

	// Dynamic Markdown Route: match any /ref/{slug} request
	mux.HandleFunc("GET /ref/{slug}", func(w http.ResponseWriter, r *http.Request) {
		slug := r.PathValue("slug") // Extract path variable (e.g., "linux" or "networking")
		mdFileName := fmt.Sprintf("%s.md", slug)

		htmlContent, err := parseMarkdownFile(mdFileName)
		if err != nil {
			// Return a 404 if the requested markdown file doesn't exist
			http.NotFound(w, r)
			return
		}

		data := PageData{Content: htmlContent}
		renderTemplate(w, r, "doc.html", data)
	})

	fmt.Println("Server running on http://127.0.0.1:8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		panic(err)
	}
}
