#!/usr/bin/env python3
"""Generate random interfaces-style configuration files."""

from __future__ import annotations

import argparse
import random
from pathlib import Path

DNS_POOL = [
    "1.1.1.1",
    "8.8.8.8",
    "9.9.9.9",
    "208.67.222.222",
]

PROFILES = ["home", "office", "lab", "vpn"]
SCRIPTS = [
    "/usr/local/bin/get-network",
    "/usr/local/bin/select-profile",
    "/opt/net/profile-map",
]


def random_private_octets(rng: random.Random) -> tuple[int, int]:
    """Pick two octets from common private address spaces."""
    choice = rng.choice([10, 172, 192])
    if choice == 10:
        return 10, rng.randint(1, 254)
    if choice == 172:
        return 172, rng.randint(16, 31)
    return 192, 168


def build_static_options(rng: random.Random) -> list[str]:
    first, second = random_private_octets(rng)
    third = rng.randint(0, 254)
    host = rng.randint(2, 250)

    address = f"{first}.{second}.{third}.{host}"
    gateway = f"{first}.{second}.{third}.1"
    network = f"{first}.{second}.{third}.0"
    broadcast = f"{first}.{second}.{third}.255"

    lines = [
        f"    address {address}",
        "    netmask 255.255.255.0",
        f"    gateway {gateway}",
    ]

    if rng.random() < 0.7:
        lines.append(f"    network {network}")
    if rng.random() < 0.6:
        lines.append(f"    broadcast {broadcast}")

    dns_count = rng.randint(1, 3)
    dns_values = " ".join(rng.sample(DNS_POOL, k=dns_count))
    lines.append(f"    dns-nameservers {dns_values}")
    return lines


def build_iface_block(name: str, method: str, rng: random.Random) -> list[str]:
    lines = [f"iface {name} inet {method}"]
    if method == "static":
        lines.extend(build_static_options(rng))
    lines.append("")
    return lines


def build_mapping_block(name: str, rng: random.Random) -> list[str]:
    lines = [f"mapping {name}", f"    script {rng.choice(SCRIPTS)}"]
    map_lines = rng.randint(1, 3)
    for _ in range(map_lines):
        profile = rng.choice(PROFILES)
        target = rng.choice(["static", "dhcp", "loopback", "fallback"])
        lines.append(f"    map {profile} {target}")
    lines.append("")
    return lines


def generate_config(file_number: int, rng: random.Random) -> str:
    iface_total = rng.randint(1, 4)
    iface_names = [f"eth{i}" for i in range(iface_total)]

    lines: list[str] = [
        f"# Random interfaces configuration {file_number}",
        f"auto lo {' '.join(iface_names)}",
        "",
        "iface lo inet loopback",
        "",
    ]

    for iface_name in iface_names:
        method = rng.choice(["static", "static", "dhcp"])
        lines.extend(build_iface_block(iface_name, method, rng))

    mappings = rng.sample(iface_names, k=rng.randint(1, len(iface_names)))
    for iface_name in mappings:
        lines.extend(build_mapping_block(iface_name, rng))

    return "\n".join(lines).rstrip() + "\n"


def parse_args() -> argparse.Namespace:
    parser = argparse.ArgumentParser(description="Generate random interfaces-style config files.")
    parser.add_argument("--count", type=int, default=40, help="Number of files to generate (default: 40).")
    parser.add_argument("--output", default="data", help="Output folder relative to project root.")
    parser.add_argument("--seed", type=int, default=None, help="Optional deterministic random seed.")
    return parser.parse_args()


def main() -> int:
    args = parse_args()
    if args.count <= 0:
        raise ValueError("--count must be greater than 0")

    project_root = Path(__file__).resolve().parent
    output_dir = (project_root / args.output).resolve()
    output_dir.mkdir(parents=True, exist_ok=True)

    rng = random.Random(args.seed)

    for i in range(1, args.count + 1):
        destination = output_dir / f"interfaces_{i:02d}.conf"
        destination.write_text(generate_config(i, rng), encoding="utf-8")

    print(f"Generated {args.count} files in {output_dir}")
    return 0


if __name__ == "__main__":
    raise SystemExit(main())
