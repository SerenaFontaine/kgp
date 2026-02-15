# KGP

Hugo-based documentation for the [KGP](../) (Kitty Graphics Protocol) Go bindings.

## Prerequisites

- [Hugo](https://gohugo.io/) Extended (0.146+)

## Setup

Initialize submodules from the repository root:

```bash
git submodule update --init --recursive
```

The theme is already configured at `docs/themes/kokoro`.

## Serve Locally

```bash
cd docs
hugo server --minify
```

Open http://localhost:1313

## Build

```bash
cd docs
hugo --minify
```

Output is in `docs/public/`.
