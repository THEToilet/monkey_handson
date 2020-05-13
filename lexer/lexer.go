package lexer

type Lexer struct {
	input         string
	position      int     // 入力における現在の位置（現在の文字を指し示す）
	readPosition  int     // これから読みこむ位置（現在の文字の次）
	ch            byte    // 現在検査中の文字
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	return l

}

func (L *Lexer) readChar() {
	if l.readPostion >= len(l.input) {
		l.ch = 0
	}
	else {
		l.ch = l.input[l.readPosition]
	}
	l.positon = l.readPosition
	l.readPositon += 1
}
