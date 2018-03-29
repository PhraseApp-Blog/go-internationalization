package main

import (
	"golang.org/x/text/language"
	"golang.org/x/text/feature/plural"
	"golang.org/x/text/message"
	"golang.org/x/text/message/catalog"
)

type entry struct {
	tag, key string
	msg      interface{}
}

var entries = [...]entry{
	{"en", "Hello World", "Hello World"},
	{"el", "Hello World", "Για Σου Κόσμε"},
	{"en", "%d task(s) remaining!", plural.Selectf(1, "%d",
		"=1", "One task remaining!",
		"=2", "Two tasks remaining!",
		"other", "[1]d tasks remaining!",
	)},
	{"el", "%d task(s) remaining!", plural.Selectf(1, "%d",
		"=1", "Μία εργασία έμεινε!",
		"=2", "Μια-δυο εργασίες έμειναν!",
		"other", "[1]d εργασίες έμειναν!",
	)},
}
func init()  {
	for _, e := range entries {
		tag := language.MustParse(e.tag)
		switch msg := e.msg.(type) {
		case string:
			message.SetString(tag, e.key, msg)
		case catalog.Message:
			message.Set(tag, e.key, msg)
		case []catalog.Message:
			message.Set(tag, e.key, msg...)
		}
	}
}

func main()  {
	p := message.NewPrinter(language.Greek)

	p.Printf("Hello World")
	p.Println()
	p.Printf("%d task(s) remaining!", 2)
	p.Println()

	p = message.NewPrinter(language.English)
	p.Printf("Hello World")
	p.Println()
	p.Printf("%d task(s) remaining!", 2)

}
