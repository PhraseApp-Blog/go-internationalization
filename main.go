package main

import (
	"PhraseApp-Blog/go-internationalization/pkg/handlers"
	"PhraseApp-Blog/go-internationalization/pkg/i18n"
	"PhraseApp-Blog/go-internationalization/pkg/model"
	"fmt"
	"log"

	"encoding/json"
	"net/http"
	"time"
)

// Initialize sample data
var speedruns = []model.Speedrun{
	{PlayerName: "Alex", Game: "Super Mario 64", Category: "Any%", Time: "16:58", SubmittedAt: time.Now()},
	{PlayerName: "Theo", Game: "The Legend of Zelda: Ocarina of Time", Category: "Any%", Time: "1:20:41", SubmittedAt: time.Now()},
}

func main() {
	// Initialize i18n package
	if err := i18n.Init(); err != nil {
		log.Fatalf("failed to initialize i18n: %v", err)
	}

	// Define routes
	http.Handle("/", detectLanguageMiddleware(handleIndex))
	http.HandleFunc("/speedruns", handleSpeedruns)
	http.HandleFunc("/speedruns/add", detectLanguageMiddleware(handleSpeedrunForm))
	http.HandleFunc("/speedrun.html", detectLanguageMiddleware(handleSpeedrunForm))

	// Serve static files
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	// Start server
	fmt.Println(i18n.T("Server listening on port %d...", 8080))
	http.ListenAndServe(":8080", nil)
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	handlers.HandleIndex(w, r, speedruns)
}

func handleSpeedruns(w http.ResponseWriter, r *http.Request) {
	// Send speedrun data as JSON response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(speedruns)
}

func handleSpeedrunForm(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// Parse request body to get new speedrun data
		var speedrun model.Speedrun
		err := json.NewDecoder(r.Body).Decode(&speedrun)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Add the current date as the submitted date
		speedrun.SubmittedAt = time.Now()

		// Add new speedrun to the global speedruns slice
		speedruns = append(speedruns, speedrun)

		// Send success response
		fmt.Fprintln(w, i18n.T("Speedrun submitted successfully!"))
	} else {
		handlers.HandleSpeedrun(w, r)
	}
}

func detectLanguageMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get the preferred locale based on the request's Accept-Language header
		lang := i18n.DetectPreferredLocale(r)
		i18n.SetCurrentLocale(lang)
		next.ServeHTTP(w, r)
	}
}
