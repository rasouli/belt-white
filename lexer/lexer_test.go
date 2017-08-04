package lexer

import (
	"testing"

	"github.com/reezaras/belt/belt-white/token"
)

func TestNextToken(t *testing.T) {
	input := "=+(){},;"

	testCases := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	lexer := New(input)

	for i, tt := range testCases {

		tok := lexer.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("test [%d]: wrong token. expected token %q got %q", i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("test [%d]: wrong literal. expected literal %s got %s", i, tt.expectedLiteral, tok.Literal)
		}
	}
}
