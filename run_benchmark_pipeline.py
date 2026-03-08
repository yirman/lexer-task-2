#!/usr/bin/env python3
"""Run benchmark_parsers.py and plot_benchmark.py in a single command."""

from __future__ import annotations

import argparse
import subprocess
import sys
from pathlib import Path


def parse_args() -> argparse.Namespace:
    parser = argparse.ArgumentParser(
        description="Run parser benchmark and generate chart in one step."
    )
    parser.add_argument(
        "--data-dir",
        default="data",
        help="Directory (relative to project root) with .conf files.",
    )
    parser.add_argument(
        "--repeat",
        type=int,
        default=1,
        help="Number of benchmark rounds to execute (default: 1).",
    )
    parser.add_argument(
        "--show-output",
        action="store_true",
        help="Show parser stdout/stderr during benchmark runs.",
    )
    parser.add_argument(
        "--json-out",
        default="benchmark_results.json",
        help="JSON report output path (relative to project root).",
    )
    parser.add_argument(
        "--image-out",
        default="benchmark_times.png",
        help="Chart image output path (relative to project root).",
    )
    parser.add_argument(
        "--title",
        default="Parser Benchmark (Elapsed Time)",
        help="Chart title.",
    )
    return parser.parse_args()


def run_step(command: list[str], cwd: Path, label: str) -> int:
    print(f"[{label}] Running: {' '.join(command)}")
    result = subprocess.run(command, cwd=cwd, check=False)
    if result.returncode != 0:
        print(f"[{label}] Failed with exit code {result.returncode}", file=sys.stderr)
    return result.returncode


def main() -> int:
    args = parse_args()
    if args.repeat <= 0:
        print("--repeat must be greater than 0", file=sys.stderr)
        return 2

    root = Path(__file__).resolve().parent

    benchmark_cmd = [
        sys.executable,
        "benchmark_parsers.py",
        "--data-dir",
        args.data_dir,
        "--repeat",
        str(args.repeat),
        "--json-out",
        args.json_out,
    ]
    if args.show_output:
        benchmark_cmd.append("--show-output")

    benchmark_code = run_step(benchmark_cmd, cwd=root, label="benchmark")
    if benchmark_code != 0:
        return benchmark_code

    plot_cmd = [
        sys.executable,
        "plot_benchmark.py",
        "--input",
        args.json_out,
        "--output",
        args.image_out,
        "--title",
        args.title,
    ]
    plot_code = run_step(plot_cmd, cwd=root, label="plot")
    if plot_code != 0:
        return plot_code

    print("Pipeline finished successfully.")
    print(f"  JSON report: {(root / args.json_out).resolve()}")
    print(f"  Chart image: {(root / args.image_out).resolve()}")
    return 0


if __name__ == "__main__":
    raise SystemExit(main())
