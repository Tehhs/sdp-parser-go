package main

import (
	"fmt"
	parser "sdp/parser"
)

func main() { 
	fmt.Println("Program Starting")
	parser.Parse([]parser.ParsingFunction{
		func(text parser.ParsingFunctionInput) parser.Signal {
			if text == "Something" { 
				return "DONE"
			}
			return "N/A"
		},
	})
	
	
} 	