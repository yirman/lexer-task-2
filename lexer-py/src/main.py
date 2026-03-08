import argparse
import sys
from pathlib import Path

from antlr4 import CommonTokenStream, FileStream, InputStream, StdinStream
from antlr4.error.ErrorListener import ErrorListener

ROOT_DIR = Path(__file__).resolve().parent.parent
GENERATED_DIR = ROOT_DIR / "generated"
if str(GENERATED_DIR) not in sys.path:
    sys.path.insert(0, str(GENERATED_DIR))

from InterfacesLexer import InterfacesLexer  # noqa: E402
from InterfacesParser import InterfacesParser  # noqa: E402


class CollectingErrorListener(ErrorListener):
    def __init__(self):
        super().__init__()
        self.errors = []

    def syntaxError(self, recognizer, offendingSymbol, line, column, msg, e):
        self.errors.append(f"line {line}:{column} {msg}")


def load_input(path: str):
    if not path or path == "-":
        return StdinStream(encoding="utf-8")
    return FileStream(path, encoding="utf-8")


def main() -> int:
    arg_parser = argparse.ArgumentParser()
    arg_parser.add_argument("-file", default="", help="Path to interfaces file. Use - for stdin")
    arg_parser.add_argument("-tree", action="store_true", help="Print parse tree when syntax is valid")
    args = arg_parser.parse_args()

    try:
        input_stream = load_input(args.file.strip())
    except OSError as ex:
        print(f"read file \"{args.file}\": {ex}", file=sys.stderr)
        return 2

    lexer = InterfacesLexer(input_stream)
    lexer_errors = CollectingErrorListener()
    lexer.removeErrorListeners()
    lexer.addErrorListener(lexer_errors)

    tokens = CommonTokenStream(lexer)
    parser = InterfacesParser(tokens)
    parser_errors = CollectingErrorListener()
    parser.removeErrorListeners()
    parser.addErrorListener(parser_errors)

    tree = parser.file_()

    all_errors = [*lexer_errors.errors, *parser_errors.errors]
    if all_errors:
        print("Syntax errors found:", file=sys.stderr)
        for err in all_errors:
            print(f"  - {err}", file=sys.stderr)
        return 1

    print("Parse successful")
    if args.tree:
        print(tree.toStringTree(recog=parser))

    return 0


if __name__ == "__main__":
    raise SystemExit(main())
