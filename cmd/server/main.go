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
	"context"
	"log"
    "errors"
    "crypto/rand"
    "crypto/sha256"
    "encoding/base64"
    "time"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
    "github.com/jackc/pgx/v5"
    "github.com/jackc/pgx/v5/pgxpool"
    "github.com/jackc/pgx/v5/pgconn"
    "golang.org/x/crypto/bcrypt"
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

func hashToken(token string) []byte {
    sum := sha256.Sum256([]byte(token))
    return sum[:]
}

func (a *App) currentUser(r *http.Request) (int, bool) {
    c, err := r.Cookie("session")
    if err != nil {
        return 0, false
    }
    var userID int
    err = a.db.QueryRow(r.Context(),
        `SELECT user_id FROM sessions WHERE token_hash = $1 AND expires_at > NOW()`,
        hashToken(c.Value)).Scan(&userID)
    if err != nil {
        return 0, false
    }
    return userID, true
}

type SignupData struct {
    Error string
    Email string
    Username string
}

type LoginData struct {
    Error string
    Email string
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

func (a *App) signupPage(w http.ResponseWriter, r *http.Request) {
    renderTemplate(w, r, "signup.html", SignupData{})
}

func (a *App) signupSubmit(w http.ResponseWriter, r *http.Request) {
    email := strings.TrimSpace(r.FormValue("email"))
    username := strings.TrimSpace(r.FormValue("username"))
    password := r.FormValue("password")

    fail := func(msg string) {
        renderTemplate(w, r, "signup.html", SignupData{Error: msg, Email: email, Username: username})
    }
    if len(password) < 8 || len(password) > 72 {
        fail("Password must be 8 to 72 characters.")
        return
    }

    hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        http.Error(w, "something went wrong", http.StatusInternalServerError)
        return
    }

    _, err = a.db.Exec(r.Context(),
        `INSERT INTO users (email, username, password_hash) VALUES ($1, $2, $3)`,
        email, username, string(hash))
    if err != nil {
        var pgErr *pgconn.PgError
        if errors.As(err, &pgErr) && pgErr.Code == "23505" {
            fail("That email or username is already taken.")
            return
        }
        log.Printf("signup insert: %v", err)
        http.Error(w, "something went wrong", http.StatusInternalServerError)
        return
    }

    http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func (a *App) loginPage(w http.ResponseWriter, r *http.Request) {
    renderTemplate(w, r, "login.html", LoginData{})
}

func (a *App) loginSubmit(w http.ResponseWriter, r *http.Request) {
    email := strings.TrimSpace(r.FormValue("email"))
    password := r.FormValue("password")

    fail := func() {
        renderTemplate(w, r, "login.html", LoginData{
            Error: "Invalid email or password.",
            Email: email,
        })
    }

    var userID int
    var hash string
    err := a.db.QueryRow(r.Context(),
        `SELECT user_id, password_hash FROM users WHERE email = $1`,
        email).Scan(&userID, &hash)
    if err != nil {
        if errors.Is(err, pgx.ErrNoRows) {
            fail()
            return
        }
        log.Printf("login lookup: %v", err)
        http.Error(w, "something went wrong", http.StatusInternalServerError)
        return
    }

    if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)); err != nil {
        fail()
        return
    }

    token, err := newSessionToken()
    if err != nil {
        log.Printf("session token: %v", err)
        http.Error(w, "something went wrong", http.StatusInternalServerError)
        return
    }
    expires := time.Now().Add(7 * 24 * time.Hour)

    _, err = a.db.Exec(r.Context(),
        `INSERT INTO sessions (token_hash, user_id, expires_at) VALUES ($1, $2, $3)`,
        hashToken(token), userID, expires)
    if err != nil {
        log.Printf("session insert: %v", err)
        http.Error(w, "something wnet wrong", http.StatusInternalServerError)
        return
    }

    http.SetCookie(w, &http.Cookie{
        Name: "session",
        Value: token,
        Path: "/",
        Expires: expires,
        HttpOnly: true,
        SameSite: http.SameSiteLaxMode,
        Secure: true,
    })

    http.Redirect(w, r, "/", http.StatusSeeOther)
}

func handToken(token string) []byte {
    sum := sha256.Sum256([]byte(token))
    return sum[:]
}

func newSessionToken() (string, error) {
    b := make([]byte, 32)
    if _, err := rand.Read(b); err != nil {
        return "", err
    }
    return base64.RawURLEncoding.EncodeToString(b), nil
}

func main() {
    ctx := context.Background()

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL is not set")
	}

	pool, err := pgxpool.New(ctx, dbURL)
	if err != nil {
		log.Fatalf("unable to create pool: %v", err)
	}
	defer pool.Close()

	if err := pool.Ping(ctx); err != nil {
		log.Fatalf("unable to reach database: %v, err")
	}
	fmt.Println("connected to db")

	app := &App{db: pool}
	_ = app

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

    mux.HandleFunc("GET /signup", app.signupPage)
    mux.HandleFunc("POST /signup", app.signupSubmit)

    mux.HandleFunc("GET /login", app.loginPage)
    mux.HandleFunc("POST /login", app.loginSubmit)

	fmt.Println("Server running on http://127.0.0.1:8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		panic(err)
	}
}
