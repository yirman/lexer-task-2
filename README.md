# Interfaces Lexer/Parser in Go

This project uses ANTLR4 to generate a lexer and parser in Go from `Interfaces.g4`.

## Prerequisites

- Go 1.25+
- ANTLR4 CLI available as `antlr4`
- Java runtime (required by ANTLR)

## Generate Parser

From the project root:

```powershell
antlr4 -Dlanguage=Go -o internal/parser -package parser Interfaces.g4
```

## Install Dependencies

```powershell
go mod tidy
```

## Run Parser

Parse file input:

```powershell
go run ./cmd/interfaces-parser -file .\examples\interfaces.conf
```

Parse from stdin:

```powershell
Get-Content .\examples\interfaces.conf | go run ./cmd/interfaces-parser -file -
```

Optional parse tree output:

```powershell
go run ./cmd/interfaces-parser -file .\examples\interfaces.conf -tree
```
