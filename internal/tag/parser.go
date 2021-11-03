package tag

import "fmt"

var Parser parser

// parser implements the following grammar using a recursive approach:
//   tag        → expression* EOF
//   expression → primary? meta? ('|' expression)*
//   primary    → String
//              | Identifier
//              | At
//              | '(' expression ')'
//   meta       → '*' | '+' | '?'
type parser struct {}

func (p *parser) Parse() Tag {
	var t Tag
	for {
		if Lexer.NextIs(EOF) {
			break
		}
		e := p.expression()
		t.AppendExpression(*e)
	}
	return t
}

// expression → primary? meta? ('|' expression)*.
func (p *parser) expression() *Expression {
	var e Expression
	e, e.Options = p.primary(), p.meta()
	for Lexer.NextIs(VerticalBar) {
		Lexer.Next()
		e.Alternative = p.expression()
	}
	return &e
}

// primary → String | Identifier | At | '(' expression ')'.
func (p *parser) primary() Expression {
	var e Expression
	if Lexer.NextIs(String, Identifier) {
		var v Value
		e.Kind, v = Lexer.Next()
		e.AppendValue(v)
	} else if Lexer.NextIs(At) {
		Lexer.Next()
		e.AppendValue(At)
	} else if Lexer.NextIs(LeftParenthesis) {
		for {
			sube := p.expression()
			e.AppendValue(sube)
			if Lexer.isAtEnd() || Lexer.NextIs(RightParenthesis) {
				break
			}
		}
		if Lexer.NextIs(RightParenthesis) {
			Lexer.Next()
		} else {
			s := fmt.Sprintf("SyntaxError: error parsing tag: missing closing ): %q", Lexer.Input)
			panic(s)
		}
	}
	return e
}

// meta → '*' | '+' | '?'.
func (p *parser) meta() Options {
	var o Options
	if Lexer.NextIs(Plus, Star, Question) {
		o.Kind, _ = Lexer.Next()
	}
	return o
}
