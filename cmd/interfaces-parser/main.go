package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"lexer-task-2/internal/parser"

	"github.com/antlr4-go/antlr/v4"
)

type collectingErrorListener struct {
	*antlr.DefaultErrorListener
	errors []string
}

func (l *collectingErrorListener) SyntaxError(_ antlr.Recognizer, _ interface{}, line, column int, msg string, _ antlr.RecognitionException) {
	l.errors = append(l.errors, fmt.Sprintf("line %d:%d %s", line, column, msg))
}

func loadInput(path string) (antlr.CharStream, error) {
	if path == "" || path == "-" {
		bytes, err := io.ReadAll(os.Stdin)
		if err != nil {
			return nil, fmt.Errorf("read stdin: %w", err)
		}
		return antlr.NewInputStream(string(bytes)), nil
	}

	stream, err := antlr.NewFileStream(path)
	if err != nil {
		return nil, fmt.Errorf("read file %q: %w", path, err)
	}
	return stream, nil
}

func main() {
	filePath := flag.String("file", "", "Path to interfaces file. Use - for stdin")
	showTree := flag.Bool("tree", false, "Print parse tree when syntax is valid")
	flag.Parse()

	input, err := loadInput(strings.TrimSpace(*filePath))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}

	lexer := parser.NewInterfacesLexer(input)
	lexerErrs := &collectingErrorListener{DefaultErrorListener: antlr.NewDefaultErrorListener()}
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(lexerErrs)

	tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewInterfacesParser(tokens)
	parserErrs := &collectingErrorListener{DefaultErrorListener: antlr.NewDefaultErrorListener()}
	p.RemoveErrorListeners()
	p.AddErrorListener(parserErrs)

	tree := p.File()

	allErrs := append(append([]string{}, lexerErrs.errors...), parserErrs.errors...)
	if len(allErrs) > 0 {
		fmt.Fprintln(os.Stderr, "Syntax errors found:")
		for _, parseErr := range allErrs {
			fmt.Fprintf(os.Stderr, "  - %s\n", parseErr)
		}
		os.Exit(1)
	}

	fmt.Println("Parse successful")
	if *showTree {
		fmt.Println(tree.ToStringTree(nil, p))
	}
}
