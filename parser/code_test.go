package parser

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/yourselfhosted/gomark/ast"
	"github.com/yourselfhosted/gomark/parser/tokenizer"
	"github.com/yourselfhosted/gomark/restore"
)

func TestCodeParser(t *testing.T) {
	tests := []struct {
		text string
		code ast.Node
	}{
		{
			text: "`Hello world!",
			code: nil,
		},
		{
			text: "`Hello world!`",
			code: &ast.Code{
				Content: "Hello world!",
			},
		},
		{
			text: "`Hello \nworld!`",
			code: nil,
		},
	}

	for _, test := range tests {
		tokens := tokenizer.Tokenize(test.text)
		node, _ := NewCodeParser().Match(tokens)
		require.Equal(t, restore.Restore([]ast.Node{test.code}), restore.Restore([]ast.Node{node}))
	}
}
