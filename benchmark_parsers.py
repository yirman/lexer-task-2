#!/usr/bin/env python3
"""Run all parser implementations and report elapsed time per parser."""

from __future__ import annotations

import argparse
import os
import subprocess
import sys
import time
from dataclasses import dataclass
from pathlib import Path
from typing import Sequence


@dataclass
class ParserRunResult:
    name: str
    elapsed_seconds: float
    total_files: int
    failed_files: list[str]


def parse_args() -> argparse.Namespace:
    parser = argparse.ArgumentParser(
        description="Benchmark lexer-go, lexer-js and lexer-py against data/*.conf"
    )
    parser.add_argument(
        "--data-dir",
        default="data",
        help="Directory (relative to project root) with .conf files.",
    )
    parser.add_argument(
        "--show-output",
        action="store_true",
        help="Show parser stdout/stderr instead of suppressing it.",
    )
    return parser.parse_args()


def project_root() -> Path:
    return Path(__file__).resolve().parent


def discover_files(root: Path, data_dir: str) -> list[Path]:
    directory = (root / data_dir).resolve()
    files = sorted(directory.glob("*.conf"))
    if not files:
        raise FileNotFoundError(f"No .conf files found in {directory}")
    return files


def run_command(
    command: Sequence[str],
    cwd: Path,
    show_output: bool,
) -> subprocess.CompletedProcess[str]:
    if show_output:
        return subprocess.run(command, cwd=cwd, check=False, text=True)
    return subprocess.run(
        command,
        cwd=cwd,
        check=False,
        text=True,
        stdout=subprocess.DEVNULL,
        stderr=subprocess.DEVNULL,
    )


def build_go_binary(root: Path, show_output: bool) -> Path:
    go_dir = root / "lexer-go"
    binary_name = "interfaces_parser.exe" if os.name == "nt" else "interfaces_parser"
    binary_path = go_dir / binary_name

    build_cmd = ["go", "build", "-o", str(binary_path), "./cmd/interfaces-parser"]
    result = run_command(build_cmd, cwd=go_dir, show_output=show_output)
    if result.returncode != 0:
        raise RuntimeError("Failed to build Go parser binary")

    return binary_path


def run_go(root: Path, files: Sequence[Path], show_output: bool) -> ParserRunResult:
    go_dir = root / "lexer-go"
    binary = build_go_binary(root, show_output=show_output)

    started = time.perf_counter()
    failures: list[str] = []
    for file_path in files:
        result = run_command(
            [str(binary), "-file", str(file_path)],
            cwd=go_dir,
            show_output=show_output,
        )
        if result.returncode != 0:
            failures.append(file_path.name)

    elapsed = time.perf_counter() - started
    return ParserRunResult("Go", elapsed, len(files), failures)


def run_js(root: Path, files: Sequence[Path], show_output: bool) -> ParserRunResult:
    js_dir = root / "lexer-js"

    started = time.perf_counter()
    failures: list[str] = []
    for file_path in files:
        result = run_command(
            ["node", "src/main.js", "-file", str(file_path)],
            cwd=js_dir,
            show_output=show_output,
        )
        if result.returncode != 0:
            failures.append(file_path.name)

    elapsed = time.perf_counter() - started
    return ParserRunResult("JavaScript", elapsed, len(files), failures)


def run_python(root: Path, files: Sequence[Path], show_output: bool) -> ParserRunResult:
    py_dir = root / "lexer-py"

    started = time.perf_counter()
    failures: list[str] = []
    for file_path in files:
        result = run_command(
            [sys.executable, "src/main.py", "-file", str(file_path)],
            cwd=py_dir,
            show_output=show_output,
        )
        if result.returncode != 0:
            failures.append(file_path.name)

    elapsed = time.perf_counter() - started
    return ParserRunResult("Python", elapsed, len(files), failures)


def print_result(result: ParserRunResult) -> None:
    print(
        f"{result.name}: elapsed={result.elapsed_seconds:.3f}s, "
        f"files={result.total_files}, failed={len(result.failed_files)}"
    )
    if result.failed_files:
        preview = ", ".join(result.failed_files[:10])
        print(f"  failing files (up to 10): {preview}")


def main() -> int:
    args = parse_args()
    root = project_root()

    try:
        files = discover_files(root, args.data_dir)
    except FileNotFoundError as ex:
        print(ex, file=sys.stderr)
        return 2

    runners = [run_python, run_go, run_js]
    print(f"Discovered {len(files)} files. Running parsers one by one...")

    overall_failures = 0
    for runner in runners:
        result = runner(root, files, args.show_output)
        print_result(result)
        overall_failures += len(result.failed_files)

    return 1 if overall_failures > 0 else 0


if __name__ == "__main__":
    raise SystemExit(main())
