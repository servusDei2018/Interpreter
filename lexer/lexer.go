package lexer

import (
	"github.com/ash9991win/Interpreter/token"
	"regexp"
)

type Lexer struct {
	input        string //The input string that the lexer is working on
	position     int    //The current position of the lexer in the string (points to the current char)
	readPosition int    //The current reading position after the current char
	ch           byte   //The current char
	line         int    //The line number
	col          int    //The col number
}

var floatRegexp *regexp.Regexp = regexp.MustCompile("^[0-9]*[.][0-9]+$")

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

//Read the character of the lexer, increment the position and readPosition
func (l *Lexer) readChar() {
	//If the next position is greater than the length, reset the char back to 0
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		//If not, assign the next char to the current char
		l.ch = l.input[l.readPosition]
	}
	//Move the position pointer by 1
	l.position = l.readPosition
	//Increment the readPosition
	l.readPosition += 1
}

//SkipChar skips the position
func (l *Lexer) skipChars(count int) {
	l.readPosition += count - 1
	l.readChar()
}

//PeekChar peeks at the next character
func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}
func (l *Lexer) skipWhiteSpaces() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		switch l.ch {
		case ' ':
			l.col++
		case '\n':
			l.line++
		case '\t':
			l.col += 8
		}
		l.readChar()
	}
}
func (l *Lexer) NextToken() token.Token {
	var tok token.Token
	l.skipWhiteSpaces()
	switch l.ch {

	case '=':
		if l.peekChar() == '=' {
			//The operator is ==
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.EQ, Literal: literal, LineNumber: l.line, ColNumber: l.col}
		} else {
			tok = newToken(token.ASSIGN, l.ch, l.line, l.col)
		}
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.NOT_EQ, Literal: literal, LineNumber: l.line, ColNumber: l.col}
		} else {
			tok = newToken(token.NOT, l.ch, l.line, l.col)
		}
	case '>':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.GT_EQ, Literal: literal, LineNumber: l.line, ColNumber: l.col}
		} else {
			tok = newToken(token.GT, l.ch, l.line, l.col)
		}
	case '<':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.LT_EQ, Literal: literal, LineNumber: l.line, ColNumber: l.col}
		} else {
			tok = newToken(token.LT, l.ch, l.line, l.col)
		}
	case ';':
		tok = newToken(token.SEMICOLON, l.ch, l.line, l.col)
	case ',':
		tok = newToken(token.COMMA, l.ch, l.line, l.col)
	case '{':
		tok = newToken(token.LBRACE, l.ch, l.line, l.col)
	case '}':
		tok = newToken(token.RBRACE, l.ch, l.line, l.col)
	case '(':
		tok = newToken(token.LPAREN, l.ch, l.line, l.col)
	case ')':
		tok = newToken(token.RPAREN, l.ch, l.line, l.col)
	case '+':
		tok = newToken(token.PLUS, l.ch, l.line, l.col)
	case '-':
		tok = newToken(token.MINUS, l.ch, l.line, l.col)
	case '/':
		tok = newToken(token.SLASH, l.ch, l.line, l.col)
	case '*':
		tok = newToken(token.STAR, l.ch, l.line, l.col)
	case '%':
		tok = newToken(token.MODULUS, l.ch, l.line, l.col)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookUpIdent(tok.Literal)
			tok.LineNumber = l.line
			tok.ColNumber = l.col
			l.col += len(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			peekedWord := l.peekWord()
			if floatRegexp.MatchString(peekedWord) {
				tok.Type = token.FLOAT
				tok.Literal = peekedWord
				l.skipChars(len(peekedWord))
			} else {
				tok.Type = token.INT
				tok.Literal = l.readNumber()
			}
			tok.LineNumber = l.line
			tok.ColNumber = l.col
			l.col += len(tok.Literal)
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch, l.line, l.col)
		}
	}
	l.readChar()
	return tok
}

func newToken(tokenType token.TokenType, ch byte, lineNumber int, columnNo int) token.Token {
	return token.Token{
		Type:       tokenType,
		Literal:    string(ch),
		LineNumber: lineNumber,
		ColNumber:  columnNo,
	}
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isBoundary(ch byte) bool {
	return ch == ' ' || ch == '\n' || ch == ';' || ch == 0
}
func (l *Lexer) readIdentifier() string {
	var result string
	for isLetter(l.ch) {
		result += string(l.ch)
		l.readChar()
	}
	return result
}

//Peeks till it encounters a space
func (l *Lexer) peekWord() string {
	startPosition := l.position
	counterPosition := startPosition
	for counterPosition < len(l.input) && !isBoundary(l.input[counterPosition]) {
		counterPosition++
	}
	return l.input[startPosition:counterPosition]
}

func (l *Lexer) peekChars(count int) string {
	startPosition := l.position
	counterPosition := startPosition
	for i := 0; counterPosition < len(l.input) && i < count; i++ {
		counterPosition++
	}
	return l.input[startPosition:counterPosition]
}

//ReadWord reads the whole word
func (l *Lexer) readWord() string {
	var result string
	for !isBoundary(l.ch) {
		result += string(l.ch)
		l.readChar()
	}
	return result
}
func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}
