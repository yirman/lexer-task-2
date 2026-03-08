#!/usr/bin/env python3
"""Run all parser implementations and report elapsed time per parser."""

from __future__ import annotations

import argparse
import json
import os
import statistics
import subprocess
import sys
import time
from dataclasses import dataclass
from pathlib import Path
from typing import Callable, Sequence


@dataclass
class ParserRunResult:
    name: str
    elapsed_seconds: float
    total_files: int
    failed_files: list[str]


ParserRunner = Callable[[Path, Sequence[Path], bool], ParserRunResult]


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
    parser.add_argument(
        "--repeat",
        type=int,
        default=1,
        help="Number of benchmark rounds to execute (default: 1).",
    )
    parser.add_argument(
        "--json-out",
        default="benchmark_results.json",
        help="Path (relative to project root) to write benchmark results JSON.",
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


def run_go(
    root: Path,
    files: Sequence[Path],
    show_output: bool,
    binary: Path,
) -> ParserRunResult:
    go_dir = root / "lexer-go"

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


def build_summary(name: str, results: Sequence[ParserRunResult]) -> dict[str, object]:
    times = [item.elapsed_seconds for item in results]
    total_failures = sum(len(item.failed_files) for item in results)
    return {
        "name": name,
        "times_seconds": times,
        "mean_seconds": statistics.mean(times),
        "min_seconds": min(times),
        "max_seconds": max(times),
        "stdev_seconds": statistics.stdev(times) if len(times) > 1 else 0.0,
        "total_failures": total_failures,
    }


def write_json_report(
    path: Path,
    data_dir: str,
    file_count: int,
    repeat: int,
    rounds: Sequence[dict[str, object]],
    summary: Sequence[dict[str, object]],
) -> None:
    report = {
        "data_dir": data_dir,
        "file_count": file_count,
        "repeat": repeat,
        "rounds": rounds,
        "summary": list(summary),
    }
    path.parent.mkdir(parents=True, exist_ok=True)
    path.write_text(json.dumps(report, indent=2), encoding="utf-8")


def main() -> int:
    args = parse_args()
    if args.repeat <= 0:
        print("--repeat must be greater than 0", file=sys.stderr)
        return 2

    root = project_root()

    try:
        files = discover_files(root, args.data_dir)
    except FileNotFoundError as ex:
        print(ex, file=sys.stderr)
        return 2

    print(
        f"Discovered {len(files)} files. Running parsers one by one "
        f"for {args.repeat} round(s)..."
    )

    go_binary = build_go_binary(root, show_output=args.show_output)
    runner_specs: list[tuple[str, ParserRunner]] = [
        ("Python", run_python),
        ("Go", lambda a, b, c: run_go(a, b, c, go_binary)),
        ("JavaScript", run_js),
    ]

    by_parser: dict[str, list[ParserRunResult]] = {name: [] for name, _ in runner_specs}
    rounds: list[dict[str, object]] = []

    overall_failures = 0
    for round_idx in range(1, args.repeat + 1):
        print(f"Round {round_idx}/{args.repeat}")
        round_results: list[dict[str, object]] = []
        for parser_name, runner in runner_specs:
            result = runner(root, files, args.show_output)
            by_parser[parser_name].append(result)
            print_result(result)
            overall_failures += len(result.failed_files)
            round_results.append(
                {
                    "name": result.name,
                    "elapsed_seconds": result.elapsed_seconds,
                    "total_files": result.total_files,
                    "failed_files": result.failed_files,
                }
            )
        rounds.append({"round": round_idx, "results": round_results})

    summary = [build_summary(name, by_parser[name]) for name, _ in runner_specs]

    json_path = (root / args.json_out).resolve()
    write_json_report(
        path=json_path,
        data_dir=args.data_dir,
        file_count=len(files),
        repeat=args.repeat,
        rounds=rounds,
        summary=summary,
    )
    print(f"Saved JSON report to: {json_path}")

    return 1 if overall_failures > 0 else 0


if __name__ == "__main__":
    raise SystemExit(main())
