// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

//go:generate gotext -srclang=en update -out=catalog/catalog.go -lang=en,el

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	_ "go-internationalization/examples/dynamic/catalog"
)

func main() {
	p := message.NewPrinter(language.Greek)
	p.Printf("Hello world!")
	p.Println()

	p.Printf("Hello", "world!")
	p.Println()

	person := "Alex"
	place := "Utah"

	p.Printf("Hello ", person, " in ", place, "!")
	p.Println()

	// Greet everyone.
	p.Printf("Hello world!")
	p.Println()

	city := "Munich"
	p.Printf("Hello %s!", city)
	p.Println()

	// Person visiting a place.
	p.Printf("%s is visiting %s!",
		person,
		place)
	p.Println()

	// Double arguments.
	miles := 1.2345
	p.Printf("%.2[1]f miles traveled (%[1]f)", miles)
}
