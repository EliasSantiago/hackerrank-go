package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	// S or C: split or combine
	// C:Classe, M:Method, V:Variable

	//Enter your code here. Read input from STDIN. Print output to STDOUT
	var input string

	fmt.Scanln(&input)
	parts := strings.Split(input, ";")
	if parts[0] == "S" {
		s := split(parts)
		fmt.Println(s)
	} else {
		s := combine(parts)
		fmt.Println(s)
	}
}

func split(parts []string) string {
	var output string
	fmt.Printf(parts[0])

	if parts[1] == "M" {
		for i, c := range parts[2] {
			if i > 0 && unicode.IsUpper(c) && !unicode.IsUpper(rune(parts[2][i-1])) {
				output += " "
			}
			t := string(c)
			output += strings.ToLower(t)
		}
		length := len(output)
		return output[:length-2]
	} else if parts[1] == "C" {
		for i, c := range parts[2] {
			if i != 0 && unicode.IsUpper(c) && !unicode.IsUpper(rune(parts[2][i-1])) {
				output += " "
			}
			t := string(c)
			output += strings.ToLower(t)
		}
		return output
	} else if parts[1] == "V" {
		for i, c := range parts[2] {
			if i > 0 && unicode.IsUpper(c) && !unicode.IsUpper(rune(parts[2][i-1])) {
				output += " "
			}
			t := string(c)
			output += strings.ToLower(t)
		}
		return output
	}
	return ""
}

func combine(parts []string) string {
	//var output string
	/*
			if parts[1] == "M" {
		    for i, c := range parts[2] {
		      if i > 0 && unicode.IsUpper(c) && !unicode.IsUpper(rune(parts[2][i-1])) {
		        output += " "
		      }
		      t := string(c)
		      output += strings.ToLower(t)
		    }
		    length := len(output)
		    return output[:length-2]
		  } else if parts[1] == "C" {
		    for i, c := range parts[2] {
		      if i != 0 && unicode.IsUpper(c) && !unicode.IsUpper(rune(parts[2][i-1])) {
		        output += " "
		      }
		      t := string(c)
		      output += strings.ToLower(t)
		    }
		    return output
		  } else if parts[1] == "V" {
		      for i, c := range parts[2] {
		      if i > 0 && unicode.IsUpper(c) && !unicode.IsUpper(rune(parts[2][i-1])) {
		        output += " "
		      }
		      t := string(c)
		      output += strings.ToLower(t)
		    }
		    return output
		  }
		  return ""
	*/
	return ""
}
