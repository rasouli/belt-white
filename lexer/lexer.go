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

	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) NextToken() token.Token {

	var tok token.Token
	l.skipWhiteSpaces()

	switch l.ch {

	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.EQ, Literal: string(ch) + string(l.ch)}
		} else {
			tok = newToken(token.ASSIGN, l.ch)
		}
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '-':
		tok = newToken(token.MINUS, l.ch)
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.NOT_EQ, Literal: string(ch) + string(l.ch)}
		} else {
			tok = newToken(token.BANG, l.ch)
		}
	case '*':
		tok = newToken(token.ASTERISK, l.ch)
	case '<':
		tok = newToken(token.LT, l.ch)
	case '>':
		tok = newToken(token.GT, l.ch)
	case '/':
		tok = newToken(token.SLASH, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)

	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdentifier(tok.Literal)
			return tok
		} else if isNumeric(l.ch) {
			tok.Literal = l.readNumber()
			tok.Type = token.INT
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return tok
}

func (l *Lexer) readNumber() string {
	curPosition := l.position
	for isNumeric(l.ch) {
		l.readChar()
	}
	return l.input[curPosition:l.position]
}

func (l *Lexer) readIdentifier() string {
	curPosition := l.position
	for isLetter(l.ch) {
		l.readChar()
	}

	return l.input[curPosition:l.position]
}

func (l *Lexer) skipWhiteSpaces() {
	for isWhitespace(l.ch) {
		l.readChar()
	}
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

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

func isLetter(ch byte) bool {
	return ch >= 'a' && ch <= 'z' || ch >= 'A' && ch <= 'Z' || ch == '_'
}

func isWhitespace(ch byte) bool {
	return ch == ' ' || ch == '\t' || ch == '\r' || ch == '\n'
}

func isNumeric(ch byte) bool {
	return ch >= '0' && ch <= '9'
}
