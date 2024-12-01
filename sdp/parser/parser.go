package parser

import ( 
	"fmt"
)


type ParsingStack struct { 
	functions []ParsingFunction
}


type ParsingFunctionInput string 
type Signal string 
type ParsingFunction func(text ParsingFunctionInput) Signal 

func StringMatch() func(ParsingFunctionInput) Signal { 
	return func(input ParsingFunctionInput) Signal { 
		return "SAVE"
	}
}


type Parser struct { 
	stacks []ParsingStack 
}

func (parser *Parser) Using(stacks []ParsingStack) { 
	parser.stacks = stacks; 
}

func (parser *Parser) Parse(input string) {


	var startIndex int = 0
	var endIndex int = 0
	for { 

		var sight string = input[startIndex:endIndex]

		fmt.Println(">", sight)

		endIndex++;

		if endIndex >= len(input) {
			break;
		}
	}

}

