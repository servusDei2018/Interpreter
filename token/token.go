package token

//TokenType holds the actual type of the tokens
type TokenType string

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
	"true":   TRUE,
	"false":  FALSE,
}

//Checks if a passed string is a user defined variable or a keyword
func LookUpIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	//Identifiers (var names) + Literals
	IDENT = "IDENT"
	INT   = "INT"
	FLOAT = "FLOAT"

	//Operators
	ASSIGN  = "="
	PLUS    = "+"
	MINUS   = "-"
	STAR    = "*"
	SLASH   = "/"
	MODULUS = "%"
	LT      = "<"
	GT      = ">"
	LT_EQ   = "<="
	GT_EQ   = ">="
	NOT_EQ  = "!="
	EQ      = "=="
	NOT     = "!"

	//Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	//Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
	GROUP    = "GROUP"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
)

//Token struct holds the token data
type Token struct {
	Type       TokenType
	Literal    string
	LineNumber int
	ColNumber  int
}
