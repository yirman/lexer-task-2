#!/usr/bin/env python3
"""Generate a benchmark chart from benchmark_parsers.py JSON output."""

from __future__ import annotations

import argparse
import json
from pathlib import Path


def parse_args() -> argparse.Namespace:
    parser = argparse.ArgumentParser(description="Plot parser benchmark times.")
    parser.add_argument(
        "--input",
        default="benchmark_results.json",
        help="Path to JSON produced by benchmark_parsers.py",
    )
    parser.add_argument(
        "--output",
        default="benchmark_times.png",
        help="Output image path (.png recommended)",
    )
    parser.add_argument(
        "--title",
        default="Parser Benchmark (Elapsed Time)",
        help="Chart title",
    )
    return parser.parse_args()


def load_report(path: Path) -> dict:
    if not path.exists():
        raise FileNotFoundError(f"Report not found: {path}")
    return json.loads(path.read_text(encoding="utf-8"))


def main() -> int:
    args = parse_args()
    root = Path(__file__).resolve().parent

    input_path = (root / args.input).resolve()
    output_path = (root / args.output).resolve()

    try:
        report = load_report(input_path)
    except FileNotFoundError as ex:
        print(ex)
        return 2

    try:
        import matplotlib.pyplot as plt
    except ModuleNotFoundError:
        print("matplotlib is required. Install with: py -3 -m pip install matplotlib")
        return 2

    summary = report.get("summary", [])
    if not summary:
        print("Invalid report: missing summary")
        return 2

    labels = [item["name"] for item in summary]
    means = [item["mean_seconds"] for item in summary]
    stdevs = [item.get("stdev_seconds", 0.0) for item in summary]

    colors = ["#1f77b4", "#ff7f0e", "#2ca02c"]

    fig, ax = plt.subplots(figsize=(8, 5))
    bars = ax.bar(labels, means, yerr=stdevs, capsize=6, color=colors[: len(labels)])

    ax.set_title(args.title)
    ax.set_ylabel("Seconds")
    ax.set_xlabel("Parser")
    ax.grid(axis="y", linestyle="--", alpha=0.35)

    for bar, value in zip(bars, means):
        ax.text(
            bar.get_x() + bar.get_width() / 2,
            bar.get_height(),
            f"{value:.3f}s",
            ha="center",
            va="bottom",
        )

    fig.tight_layout()
    output_path.parent.mkdir(parents=True, exist_ok=True)
    fig.savefig(output_path, dpi=150)
    print(f"Saved chart to: {output_path}")
    return 0


if __name__ == "__main__":
    raise SystemExit(main())
