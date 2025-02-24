package lexer

import (
	"fmt"
	"monkey/token"
)

type Lexer struct {
	input        string
	inputSize    int
	position     int  //current position in input. Points to current character
	readPosition int  //current reading readPosition. Points after current character
	ch           byte //current character
}

func (l *Lexer) Display() {
	fmt.Printf("%+v\n", l)
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= l.inputSize {
		return 0
	}
	return l.input[l.readPosition]
}

func New(input string) *Lexer {
	l := &Lexer{input: input, inputSize: len(input)}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= l.inputSize {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{
		Type:    tokenType,
		Literal: string(ch),
	}
}

func isLetter(char byte) bool {
	return 'a' <= char && char <= 'z' || 'A' <= char && char <= 'Z' || char == '_'
}

func isDigit(char byte) bool {
	return '0' <= char && char <= '9'
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) NextToken() token.Token {
	var t token.Token
	l.skipWhitespace()
	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			t = token.Token{Type: token.EQUAL, Literal: "=="}
			l.readChar()
			break
		}
		t = newToken(token.ASSIGN, l.ch)
	case ';':
		t = newToken(token.SEMICOLON, l.ch)
	case '(':
		t = newToken(token.LPAREN, l.ch)
	case ')':
		t = newToken(token.RPAREN, l.ch)
	case ',':
		t = newToken(token.COMMA, l.ch)
	case '+':
		t = newToken(token.PLUS, l.ch)
	case '{':
		t = newToken(token.LBRACE, l.ch)
	case '}':
		t = newToken(token.RBRACE, l.ch)
	case '!':
		if l.peekChar() == '=' {
			t = token.Token{Type: token.NOT_EQUAL, Literal: "!="}
			l.readChar()
			break
		}
		t = newToken(token.BANG, l.ch)
	case '*':
		t = newToken(token.ASTERISK, l.ch)
	case '<':
		t = newToken(token.LESS_THAN, l.ch)
	case '>':
		t = newToken(token.GREATER_THAN, l.ch)
	case '-':
		t = newToken(token.MINUS, l.ch)
	case '/':
		t = newToken(token.SLASH, l.ch)
	case 0:
		t.Literal = ""
		t.Type = token.EOF
	default:
		if isLetter(l.ch) {
			t.Literal = l.readIdentifier()
			t.Type = token.LookupIdentifier(t.Literal)
			return t
		}
		if isDigit(l.ch) {
			t.Type = token.INT
			t.Literal = l.readNumber()
			return t
		}
		t = newToken(token.ILLEGAL, l.ch)
	}
	l.readChar()
	return t
}
