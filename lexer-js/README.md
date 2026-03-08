# Interfaces Lexer/Parser in JavaScript

This folder contains a JavaScript ANTLR4 lexer/parser implementation for `Interfaces.g4`.

## Prerequisites

- Node.js 18+
- Java runtime (required by ANTLR)
- ANTLR4 CLI available as `antlr4`

## Install Dependencies

```powershell
npm install
```

## Generate Parser Sources

From this `lexer-js` folder:

```powershell
.\generate.ps1
```

Or with npm script:

```powershell
npm run generate
```

## Run Parser

Parse file input:

```powershell
node src/main.js -file .\examples\interfaces.conf
```

Parse from stdin:

```powershell
Get-Content .\examples\interfaces.conf | node src/main.js -file -
```

Optional parse tree output:

```powershell
node src/main.js -file .\examples\interfaces.conf -tree
```
