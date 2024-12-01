package parser

import ( 
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

func (self *Parser) Parse(input string) {

}

func (self *Parser) Using(stacks []ParsingStack) { 

}