package parser

type TokenType int

const (
	TOK_SCENE TokenType = iota
	TOK_TEXT
)

type Token struct {
	Type   TokenType
	Value  string
	Indent int
	Line   int
}
