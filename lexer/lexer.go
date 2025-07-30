package lexer

type lexer struct {
	input string
	position int 
	readPosition int
	ch byte
}

func New(input string) *Lexer {
	l := &LexerP{input: input}
	return l
}
