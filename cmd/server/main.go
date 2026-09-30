package main

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
	"strings"
)

// Helper to render full pages or HTMX content fragments
func renderTemplate(w http.ResponseWriter, r *http.Request, pageFile string, data interface{}) {
	// If the request comes from HTMX, render ONLY the content block
	if r.Header.Get("HX-Request") == "true" {
		tmpl, err := template.ParseFiles(filepath.Join("views", "pages", pageFile))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		tmpl.ExecuteTemplate(w, "content", data)
		return
	}

	// Otherwise, render the full page with base layout
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

	// 1. Static File Server (CSS, JS, Images)
	fs := http.FileServer(http.Dir("static"))
	mux.Handle("GET /static/", http.StripPrefix("/static/", fs))

	// 2. Page Routes
	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, r, "home.html", nil)
	})

	mux.HandleFunc("GET /ref/linux", func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, r, "cheat_sheet.html", nil)
	})

	// 3. Search Endpoint (HTMX)
	mux.HandleFunc("GET /api/search", func(w http.ResponseWriter, r *http.Request) {
		query := strings.TrimSpace(r.URL.Query().Get("q"))
		if query == "" {
			fmt.Fprint(w, "")
			return
		}
		fmt.Fprintf(w, `<div style="color: var(--accent);">Filtering results for: <strong>%s</strong></div>`, template.HTMLEscapeString(query))
	})

	fmt.Println("Server running on http://127.0.0.1:8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		panic(err)
	}
}
