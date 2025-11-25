package tree_sitter_scaf_test

import (
	"testing"

	tree_sitter "github.com/tree-sitter/go-tree-sitter"
	tree_sitter_scaf "github.com/rlch/tree-sitter-scaf/bindings/go"
)

func TestCanLoadGrammar(t *testing.T) {
	language := tree_sitter.NewLanguage(tree_sitter_scaf.Language())
	if language == nil {
		t.Errorf("Error loading Scaf grammar")
	}
}
