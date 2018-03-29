//package main
//
//import (
//	"golang.org/x/text/message"
//	"golang.org/x/text/language"
//	"fmt"
//	"golang.org/x/text/message/catalog"
//	"golang.org/x/text/feature/plural"
//	"golang.org/x/text/currency"
//	"golang.org/x/text/number"
//)
//
//func init()  {
//	message.SetString(language.Greek, "%s went to %s.",  "%s πήγε στήν %s.")
//	message.SetString(language.AmericanEnglish, "%s went to %s.",  "%s is in %s.")
//	message.SetString(language.Greek, "%s has been stolen.",  "%s κλάπηκε")
//	message.SetString(language.AmericanEnglish, "%s has been stolen.",  "%s has been stolen.")
//	message.SetString(language.Greek, "How are you?",  "Πώς είστε?.")
//	message.Set(language.Greek, "Hello %s!",
//		catalog.Var("person", catalog.String("Λία")),
//		catalog.String("Γιά σου ${person}!"),
//		)
//	message.Set(language.Greek, "You have %d. problem",
//		plural.Selectf(1, "%d",
//			"=1", "Έχεις ένα πρόβλημα",
//			"=2", "Έχεις %[1]d πρόβληματα",
//			"other", "Έχεις πολλά πρόβληματα",
//		))
//	message.Set(language.Greek, "You have %d days remaining",
//		plural.Selectf(1, "%d",
//			"one", "Έχεις μία μέρα ελεύθερη",
//			"other", "Έχεις %[1]d μέρες ελεύθερες",
//		))
//	message.Set(language.Greek, "You are %d minute(s) late.",
//		catalog.Var("minutes", plural.Selectf(1, "%d", "one", "λεπτό")),
//		catalog.String("Αργήσατε %[1]d ${minutes}."))
//}
//
//func main()  {
//	p := message.NewPrinter(language.Greek)
//	p.Printf("You have %d. problem", 1)
//	fmt.Println()
//	p.Printf("You have %d. problem", 2)
//	fmt.Println()
//	p.Printf("You have %d. problem", 5)
//	fmt.Println()
//	p.Printf("You have %d days remaining", 1)
//	fmt.Println()
//	p.Printf("You have %d days remaining", 10)
//	fmt.Println()
//	p.Printf("You are %d minute(s) late.", 1)
//	p = message.NewPrinter(language.AmericanEnglish)
//	p.Printf("%s went to %s.", "Peter", "England")
//	fmt.Println()
//	p.Printf("%s has been stolen.", "The Gem")
//	fmt.Println()
//
//
//	p.Printf("%d", currency.Symbol(currency.USD.Amount(0.1)))
//	fmt.Println()
//	p.Printf("%d", currency.NarrowSymbol(currency.JPY.Amount(1.6)))
//	fmt.Println()
//	p.Printf("%d", currency.ISO.Kind(currency.Cash)(currency.EUR.Amount(12.255)))
//	fmt.Println()
//}
