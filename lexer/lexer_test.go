package lexer

import (
	"github.com/ash9991win/Interpreter/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `
	let f = 10;
	let fnTest = fn add(x,y){};
	if i >= 5 {
		add(i,i)
	} else {

	}
	`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LET, "let"},
		{token.IDENT, "f"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "fnTest"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fn"},
		{token.IDENT, "add"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},
		{token.IF, "if"},
		{token.IDENT, "i"},
		{token.GT_EQ, ">="},
		{token.INT, "5"},
		{token.LBRACE, "{"},
		{token.IDENT, "add"},
		{token.LPAREN, "("},
		{token.IDENT, "i"},
		{token.COMMA, ","},
		{token.IDENT, "i"},
		{token.RPAREN, ")"},
		{token.RBRACE, "}"},
		{token.ELSE, "else"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
	}

	l := New(input)
	for i, tt := range tests {
		tok := l.NextToken()
		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokenttype wrong. expected = %q, got = %q", i, tt.expectedType, tok.Type)
		}
		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - tokenliteral wrong. expected = %q, got = %q", i, tt.expectedLiteral, tok.Literal)
		}
	}
}
