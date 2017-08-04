package lexer

import (
	"github.com/reezaras/belt/belt-white/token"
)

type Lexer struct {
	input        string
	position     int // the position of the current character
	readPosition int // the next place we're gonna read
	ch           byte
}

func New(input string) *Lexer {
	return &Lexer{input: input}
}

func (l *Lexer) NextToken() token.Token {

	l.readChar()

	var x token.Token
	switch l.ch {

	case '=':
		x = newToken(token.ASSIGN, l.ch)
	case '+':
		x = newToken(token.PLUS, l.ch)
	case ',':
		x = newToken(token.COMMA, l.ch)
	case ';':
		x = newToken(token.SEMICOLON, l.ch)
	case '(':
		x = newToken(token.LPAREN, l.ch)
	case ')':
		x = newToken(token.RPAREN, l.ch)
	case '{':
		x = newToken(token.LBRACE, l.ch)
	case '}':
		x = newToken(token.RBRACE, l.ch)

	case 0:
		x.Literal = ""
		x.Type = token.EOF
	}

	return x
}

func newToken(tokenType token.TokenType, literal byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(literal)}
}

func (l *Lexer) readChar() {

	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition += 1
}
