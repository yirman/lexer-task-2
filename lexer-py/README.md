# Interfaces Lexer/Parser in Python

This folder contains a Python ANTLR4 lexer/parser implementation for `Interfaces.g4`.

## Prerequisites

- Python 3.10+
- Java runtime (required by ANTLR)
- ANTLR4 CLI available as `antlr4`

## Install Dependencies

From `lexer-py`:

```powershell
py -3 -m pip install -r requirements.txt
```

## Generate Parser Sources

From this `lexer-py` folder:

```powershell
.\generate.ps1
```

Or directly:

```powershell
antlr4 -Dlanguage=Python3 -o generated ..\Interfaces.g4
```

## Run Parser

Parse file input:

```powershell
py -3 src\main.py -file .\examples\interfaces.conf
```

Parse from stdin:

```powershell
Get-Content .\examples\interfaces.conf | py -3 src\main.py -file -
```

Optional parse tree output:

```powershell
py -3 src\main.py -file .\examples\interfaces.conf -tree
```
