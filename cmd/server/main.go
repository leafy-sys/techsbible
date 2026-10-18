package main

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
    "github.com/jackc/pgx/v5/pgxpool"
)

type App struct {
    db *pgxpool.Pool
}

var (
	// Matches [[Target Page|Display Text]]
	wikiLinkAliasRegex = regexp.MustCompile(`\[\[([^\]\|]+)\|([^\]]+)\]\]`)
	// Matches [[Target Page]]
	wikiLinkSimpleRegex = regexp.MustCompile(`\[\[([^\]]+)\]\]`)
	// Matches Obsidian Callouts like > [!info]- Title
	calloutRegex = regexp.MustCompile(`(?m)^>\s*\[!(\w+)\][-+]?\s*(.*)`)
)

func preprocessObsidianMarkdown(md []byte) []byte {
	// 1. Convert [[Target Page|Display Text]] -> [Display Text](/ref/target_page)
	md = wikiLinkAliasRegex.ReplaceAllFunc(md, func(match []byte) []byte {
		sub := wikiLinkAliasRegex.FindSubmatch(match)
		target := strings.ToLower(strings.ReplaceAll(string(sub[1]), " ", "_"))
		display := string(sub[2])
		return []byte(fmt.Sprintf("[%s](/ref/%s)", display, target))
	})

	// 2. Convert [[Target Page]] -> [Target Page](/ref/target_page)
	md = wikiLinkSimpleRegex.ReplaceAllFunc(md, func(match []byte) []byte {
		sub := wikiLinkSimpleRegex.FindSubmatch(match)
		target := strings.ToLower(strings.ReplaceAll(string(sub[1]), " ", "_"))
		display := string(sub[1])
		return []byte(fmt.Sprintf("[%s](/ref/%s)", display, target))
	})

	// 3. Convert Obsidian callouts > [!info] -> Standard Blockquotes with bold titles
	md = calloutRegex.ReplaceAll(md, []byte("> **${1}:**${2}"))

	return md
}

// 1. Defined PageData struct to fix compiler error
type PageData struct {
	Content template.HTML
}

var mdParser = goldmark.New(
	goldmark.WithExtensions(
		extension.GFM, // Enables Markdown tables
	),
	goldmark.WithParserOptions(
		parser.WithAutoHeadingID(),
	),
)

func parseMarkdownFile(filename string) (template.HTML, error) {
	path := filepath.Join("content", filename)
	mdData, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	// Preprocess Obsidian / Wikilink syntax
	cleanMdData := preprocessObsidianMarkdown(mdData)

	var buf bytes.Buffer
	if err := mdParser.Convert(cleanMdData, &buf); err != nil {
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
    ctx := context.Background()
	mux := http.NewServeMux()

	// Static assets
	fs := http.FileServer(http.Dir("static"))
	mux.Handle("GET /static/", http.StripPrefix("/static/", fs))

	// Home page - added /{$} so it strictly matches the root domain
	mux.HandleFunc("GET /{$}", func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, r, "home.html", nil)
	})

	// Dynamic Markdown Route: match any /ref/{slug} request
	mux.HandleFunc("GET /ref/{slug}", func(w http.ResponseWriter, r *http.Request) {
		slug := r.PathValue("slug") // Extract path variable
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

    app := &App{db: pool}
    _ = app
}
