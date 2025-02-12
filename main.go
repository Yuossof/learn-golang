package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	// number := 2
	// fmt.Printf("OOD: %v => %v", number % 2 == 0, number)

	//------ Compare ------//
	a := "str2"
	a2 := "str2"
	fmt.Println(strings.Compare(a, a2))

	//------ Contains ------//
	b := "Youssof"
	b2 := "Y"
	fmt.Println(strings.Contains(b, b2))

	//------ ContainsAny ------//
	c := "Ydddoussof"
	c2 := "uss"
	fmt.Println(strings.ContainsAny(c, c2))

	//------ ContainsFunc ------//
	f := func(r rune) bool {
		return r == 'a' || r == 'e' || r == 'i' || r == 'o' || r == 'u'
	}
	fmt.Println(strings.ContainsFunc("i", f))

	//------ ContainsRune ------//
	runeN := "ahmed"
	// a = 97
	fmt.Println(strings.ContainsRune(runeN, 97))

	//------ Count ------//
	fmt.Println(strings.Count("cheese", "e")) // 3 is number of e chars

	//------ Cut ------//
	before, after, found := strings.Cut("golang", "lan") // before(go) middle(lan) after(g)
	println(before, after, found)                        // go, g, true

	//------ CutPrefix ------//
	after2, found2 := strings.CutPrefix("golang", "go")
	after3, found3 := strings.CutPrefix("golang", "lan")
	println(after2, found2) // lang, true
	println(after3, found3) // golang, false

	//------ EqualFold ------//
	fmt.Println(strings.EqualFold("GO", "go")) // true , // you can use capital or small
	fmt.Println(strings.EqualFold("Og", "go")) // false

	// ------------------------------------------------------//
	// importants = *

	//------ Fields * ------//
	fmt.Printf("Fields are: %q", strings.Fields("  foo bar  baz   ")) // convert strings that including spaces to an array

	//------ FieldsFunc * ------//
	fnc := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c) // check if letter is anything else number or letter
	}
	fmt.Printf("Fields are: %q", strings.FieldsFunc("  foo1;bar2,baz3...", fnc)) // Fields are: ["foo1" "bar2" "baz3"]
	fmt.Println("")

	//------ HasPrefix  ------//
	fmt.Println(strings.HasPrefix("Gopher", "Go")) // true
	fmt.Println(strings.HasPrefix("Gopher", "ph")) // false

	//------ HasSuffix  ------//
	fmt.Println(strings.HasSuffix("Amigo", "go")) // true
	fmt.Println(strings.HasSuffix("Amigo", "mi")) // false
	fmt.Println(strings.HasSuffix("Amigo", "Am")) // false

	//------ IndexByte * ------//
	fmt.Println(strings.IndexByte("golang", 'g'))  // 0
	fmt.Println(strings.IndexByte("gophers", 'h')) // 3

	//------ IndexFunc * ------//
	checkLetter := func(c rune) bool {
		return unicode.Is(unicode.Arabic, c) // Check if this character belongs to the Arabic language
	}
	fmt.Println(strings.IndexFunc("Hello, ياعالم", checkLetter)) // if c = arabic char => return char index => return 7
	fmt.Println(strings.IndexFunc("Hello, world", checkLetter))  // => (-1) because this character is not an Arabic character, so it returns (-1)

	//------ IndexRune * ------//
	fmt.Println(strings.IndexRune("chicken", 'k')) // 4

	//------ Join * ------//
	s := []string{"foo", "bar", "baz"}
	fmt.Println(strings.Join(s, ",")) // "foo,bar,baz"

}
