package i18n

import (
	"github.com/leonelquinteros/gotext"
)

// LanguageCode represents a language code.
type LanguageCode string

// Constants representing supported language codes.
const (
	GR LanguageCode = "el_GR" // Greek language code
	EN LanguageCode = "en_US" // English language code
	AR LanguageCode = "ar_SA" // Arabic language code
)

// langMap stores Locale instances for each language code.
var langMap = make(map[LanguageCode]*gotext.Locale)

// LanguageDirectionMap maps language codes to their typical text directions
var LanguageDirectionMap = map[LanguageCode]string{
	"el_GR": "ltr", // Greek language code
	"en_US": "ltr", // English language code
	"ar_SA": "rtl", // Arabic language code
}

// String returns the string representation of a LanguageCode.
func (l LanguageCode) String() string {
	return string(l)
}

// T returns the translated string for the given key in the specified language.
func (l LanguageCode) T(s string) string {
	// Check if a Locale exists for the specified language code
	if lang, ok := langMap[l]; ok {
		// Retrieve the translated string from the Locale
		return lang.Get(s)
	}
	// Return the original string if no translation is available
	return s
}
