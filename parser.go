package alligopher

import (
	"github.com/alecthomas/repr"
	"github.com/celicoo/alligopher/internal/tag"
)

// NewParser constructs and returns a new Parser for the given grammar.
func NewParser(grammar interface{}) *Parser {
	//v := reflection.ValueOf(grammar)
	//if v.Kind() != reflection.Struct {
	//	s := fmt.Sprintf("cannot use grammar (type %[1]v) as type struct in argument to NewParser", v.Type().Name())
	//	panic(s)
	//}
	return &Parser{grammar}
}

type Parser struct {
	grammar interface{}
}

func (p *Parser) Parse(input string) interface{} {
	//Lexer.Init(input)
	//for {
	//	if Lexer.isAtEnd() {
	//		break
	//	}
	//	token := Lexer.Scan()
	//	// start at first field of grammar.
	//	// does it have an alligopher tag?
	//	//  no -> error
	//	//  yes ->
	//
	//	// for i in grammar.fields {
	//	//   field = grammar.fields[i]
	//	//   if field.HasAlligopherTag() {
	//	//     t := tag.New(field.Tag)
	//	//     if t.Match(token.Name) {
	//	//
	//	//     }
	//	//   } else {
	//	//     reportError(errFieldWithoutTag)
	//	//   }
	//	// }
	//	repr.Println(token)
	//}
	t := tag.New("|LET")
	repr.Print(t)
	return p.grammar
}
