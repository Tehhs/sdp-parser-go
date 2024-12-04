package main

import (
	"fmt"
	parser "sdp/parser"
)

func main() { 
	fmt.Println("Program Starting")
	// parser.Parse([]parser.ParsingFunction{
	// 	func(text parser.ParsingFunctionInput) parser.Signal {
	// 		if text == "Something" { 
	// 			return "DONE"
	// 		}
	// 		return "N/A"
	// 	},
	// })
	
	p := parser.Parser{}
	p.Using([]parser.ParsingStack{
		// parser.ParsingStackUtil.New(
		// 	// func(text parser.ParsingFunctionInput) parser.Signal {
		// 	// 	return "YES"
		// 	// },
		// ),
	})
		
	
} 	