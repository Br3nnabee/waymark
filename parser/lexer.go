package parser

import (
	"fmt"
	"strings"
)

func Tokenize(src string) ([]Token, error) {
	var tokens []Token
	lines := strings.Split(src, "\n")

	for i, raw := range lines {
		// Strip \r for Windows line endings
		raw = strings.TrimRight(raw, "\r")

		trimmed := strings.TrimSpace(raw)

		// Skip blanks and comments
		if trimmed == "" || strings.HasPrefix(trimmed, "#") {
			continue
		}

		indent, err := countIndent(raw)
		if err != nil {
			return nil, fmt.Errorf("line %d: %w", i+1, err)
		}

		tok := Token{Indent: indent, Line: i + 1}

		switch {
		case strings.HasPrefix(trimmed, "scene "):
			tok.Type, tok.Value = TOK_SCENE, after(trimmed, "scene ")
		case strings.HasPrefix(trimmed, "var "):
			tok.Type, tok.Value = TOK_VAR, after(trimmed, "var ")
		default:
			tok.Type, tok.Value = TOK_TEXT, trimmed
		}

		tokens = append(tokens, tok)
	}

	return tokens, nil
}

// Count the number of indents in a line. Only accepts tabs, not spaces
func countIndent(line string) (int, error) {
	if len(line) == 0 {
		return 0, nil
	}

	// Detect if the line uses spaces for indenting
	if line[0] == ' ' {
		spaces := 0
		for _, ch := range line {
			if ch != ' ' {
				break
			}
			spaces++
		}
		return 0, fmt.Errorf(
			"indent uses spaces (%d found) — Waymark requires tabs. "+
				"Check your editor's 'indent with tabs' setting", spaces,
		)
	}

	indent := 0
	for _, ch := range line {
		if ch != '\t' {
			break
		}
		indent++
	}
	return indent, nil
}

// Wrapper for trimming prefix from line
func after(s, prefix string) string {
	return strings.TrimSpace(strings.TrimPrefix(s, prefix))
}
