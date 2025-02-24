package token

import "fmt"

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

func (t *Token) Display() {
	fmt.Printf("TokenType: %v\n", t.Type)
	fmt.Printf("TokenLiteral: %v\n", t.Literal)
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	//Identifiers and literals
	IDENT = "IDENT"
	INT   = "INT"

	//Operators
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"

	LESS_THAN    = "<"
	GREATER_THAN = ">"

	//Delimiters
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"
	EQUAL     = "=="
	NOT_EQUAL = "!="

	//Keywords
	FUNCTION = "FN"
	LET      = "LET"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
)

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

func LookupIdentifier(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
