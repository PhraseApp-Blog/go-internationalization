package main

import (
	"html"
	"log"
	"net/http"
	"flag"
	"time"
	"golang.org/x/text/message"
	"golang.org/x/text/language"

	_ "go-internationalization/catalog"
	"context"
	"fmt"
)

//go:generate gotext -srclang=en update -out=catalog/catalog.go -lang=en,el

var matcher = language.NewMatcher(message.DefaultCatalog.Languages())

type contextKey int

const (
	httpPort  = "8080"
	messagePrinterKey contextKey = 1
)

func withMessagePrinter(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		lang, ok := r.URL.Query()["lang"]

		if !ok || len(lang) < 1 {
			lang = append(lang, language.English.String())
		}
		tag,_, _ := matcher.Match(language.MustParse(lang[0]))
		p := message.NewPrinter(tag)
		ctx := context.WithValue(context.Background(), messagePrinterKey, p)

		next.ServeHTTP(w, r.WithContext(ctx))
	}
}

func PrintMessage(w http.ResponseWriter, r *http.Request) {
	p := r.Context().Value(messagePrinterKey).(*message.Printer)
	p.Fprintf(w,"Hello, %v", html.EscapeString(r.Host))
}

func main() {
	var port string
	flag.StringVar(&port, "port", httpPort, "http port")
	flag.Parse()

	server := &http.Server{
		Addr:           ":" + port,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 16,
		Handler:        http.HandlerFunc(withMessagePrinter(PrintMessage))}

	log.Fatal(server.ListenAndServe())
}

