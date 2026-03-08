import fs from 'node:fs/promises';
import process from 'node:process';

import antlr4 from 'antlr4';
import InterfacesLexer from '../generated/InterfacesLexer.js';
import InterfacesParser from '../generated/InterfacesParser.js';

class CollectingErrorListener extends antlr4.error.ErrorListener {
  constructor() {
    super();
    this.errors = [];
  }

  syntaxError(recognizer, offendingSymbol, line, column, msg, e) {
    this.errors.push(`line ${line}:${column} ${msg}`);
  }
}

function parseArgs(argv) {
  let filePath = '';
  let showTree = false;

  for (let i = 2; i < argv.length; i += 1) {
    const arg = argv[i];
    if (arg === '-file' && i + 1 < argv.length) {
      filePath = argv[i + 1];
      i += 1;
      continue;
    }
    if (arg === '-tree') {
      showTree = true;
      continue;
    }
    throw new Error('Usage: node src/main.js [-file <path|->] [-tree]');
  }

  return { filePath, showTree };
}

async function readInput(path) {
  if (!path || path === '-') {
    const chunks = [];
    for await (const chunk of process.stdin) {
      chunks.push(chunk);
    }
    return Buffer.concat(chunks.map((c) => (Buffer.isBuffer(c) ? c : Buffer.from(c)))).toString('utf8');
  }

  try {
    return await fs.readFile(path, 'utf8');
  } catch (err) {
    throw new Error(`read file "${path}": ${err.message}`);
  }
}

async function main() {
  const { filePath, showTree } = parseArgs(process.argv);
  const inputText = await readInput(filePath);

  const input = new antlr4.InputStream(inputText);
  const lexer = new InterfacesLexer(input);
  const lexerErrs = new CollectingErrorListener();
  lexer.removeErrorListeners();
  lexer.addErrorListener(lexerErrs);

  const tokens = new antlr4.CommonTokenStream(lexer);
  const parser = new InterfacesParser(tokens);
  const parserErrs = new CollectingErrorListener();
  parser.removeErrorListeners();
  parser.addErrorListener(parserErrs);

  const tree = parser.file();

  const allErrs = [...lexerErrs.errors, ...parserErrs.errors];
  if (allErrs.length > 0) {
    console.error('Syntax errors found:');
    for (const err of allErrs) {
      console.error(`  - ${err}`);
    }
    process.exitCode = 1;
    return;
  }

  console.log('Parse successful');
  if (showTree) {
    console.log(tree.toStringTree(parser.ruleNames));
  }
}

main().catch((err) => {
  console.error(err.message);
  process.exit(2);
});
