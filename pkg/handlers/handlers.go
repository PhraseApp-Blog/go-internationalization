package handlers

import (
	"PhraseApp-Blog/go-internationalization/pkg/i18n"
	"PhraseApp-Blog/go-internationalization/pkg/model"
	"context"
	"fmt"
	"html/template"
	"net/http"
	"time"

	"golang.org/x/text/language"
)

func HandleTemplate(w http.ResponseWriter, r *http.Request, tmplName string, data interface{}) {
	// Parse the template
	tmpl, err := template.New(tmplName).Funcs(template.FuncMap{
		"T":  i18n.T,
		"TN": i18n.TN,
		"FormatLocalizedDate": func(submittedAt time.Time, currentLanguage i18n.LanguageCode) string {
			return i18n.FormatLocalizedDate(submittedAt, language.Make(currentLanguage.String()))
		},
		"FormatNumber": func(number int64, currentLanguage i18n.LanguageCode) string {
			return i18n.FormatNumber(number, language.Make(currentLanguage.String()))
		},
	}).ParseFiles("static/" + tmplName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Execute the template with the translation function and data
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

type FormattedSpeedrun struct {
	PlayerName  string
	Game        string
	Category    string
	Time        string
	SubmittedAt string
}

func HandleIndex(w http.ResponseWriter, r *http.Request, data []model.Speedrun) {
	fmt.Println(i18n.GetCurrentLanguage())
	HandleTemplate(w, r, "index.html", map[string]interface{}{
		"Title":              "Speedrun Leaderboard",
		"Header":             "Speedrun Leaderboard",
		"PlayerName":         "Player Name",
		"Game":               "Game",
		"Category":           "Category",
		"Time":               "Time",
		"Date":               "Submitted At",
		"Data":               data,
		"Dir":                i18n.LanguageDirectionMap[i18n.GetCurrentLanguage()],
		"CurrentLanguage":    i18n.GetCurrentLanguage(),
		"SupportedLanguages": i18n.GetSupportedLanguages(),
	})
}

func HandleSpeedrun(w http.ResponseWriter, r *http.Request) {
	HandleTemplate(w, r, "speedrun.html", map[string]interface{}{
		"Header":             "Add New Speedrun",
		"Title":              "Add New Speedrun",
		"PlayerName":         "Player Name",
		"Game":               "Game",
		"Category":           "Category",
		"Submit":             "Submit",
		"Time":               "Time",
		"Dir":                i18n.LanguageDirectionMap[i18n.GetCurrentLanguage()],
		"CurrentLanguage":    i18n.GetCurrentLanguage(),
		"SupportedLanguages": i18n.GetSupportedLanguages(),
	})
}

// Middleware to set the current language in the context of the request
func SetCurrentLanguage(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get the language from the request (you may need to adjust this based on your implementation)
		lang := r.URL.Query().Get("lang")

		// Set the language in the context of the request
		ctx := context.WithValue(r.Context(), "lang", lang)

		// Call the next handler in the chain with the new context
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
