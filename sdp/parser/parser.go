package parser


type ParsingFunctionInput string 
type Signal string 
type ParsingFunction func(text ParsingFunctionInput) Signal 

func StringMatch() func(ParsingFunctionInput) Signal { 
	return func(input ParsingFunctionInput) Signal { 
		return "SAVE"
	}
}

func Parse([]ParsingFunction) {

}