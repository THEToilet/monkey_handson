package lexer

import(
	"testing"
	"../token"
	)


func TestNextToken(t *testing.T) {
	input := `=+(){},;`

	tests := [] struct {
		expectedType   token.TokenType
		expectedLitera string
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

	}

	l := New(input)

	for i,tt := range test{
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tokType)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expextedLiteral, tok.Litera)
		}
	}
}
