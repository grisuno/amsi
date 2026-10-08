# Second Brain

*Last synthesized: 2026-10-07 | 3 files | 1 concept pages | offline, zero tokens*

> Raw sources -> readmenator wiki -> links (Karpathy LLM Wiki Pattern, deterministic).
> Start here, then open one community page. Prefer grep over full reads.

## Vault Overview

The codebase centres on `amsi.go`, `app.py`, `install.sh`. Architecturally it is 1 layers, dominant utility (3 files) across 1 import-based communities. Recorded risk surface: 0 security findings and 0 dependency cycles.

Communities are self-contained in the resolved import graph; no cross-boundary bridges were recorded.

Open work clusters around documentation (33% file coverage), 0 security findings, 0 taint paths, and 4 suggested exploration questions in `queries.md`.

## Stats

| Metric | Value |
|--------|-------|
| Files | 3 |
| Symbols | 2 |
| Resolved imports | 0 |
| Languages | go, py, sh |
| Communities | 1 |
| Doc coverage | 33% (1/3 files) |
| Security findings | 0 |
| Estimated read cost | ~544 tokens (chars/4, offline so $0) |

## Reading Order

1. Skim Stats and God Nodes below for blast radius.
2. Open the largest community page first, then follow Connections.
3. Use `queries.md` for the next question; log the answer there.

```
grep -rn '<keyword>' index.md community_*.md
readmenator query "<question>" --target readmenator_amsi_1kymaxo5
```

## Concept Wiki

- [root (3 files, cohesion 1.00)](./community_0_root.md)

## God Nodes

| File | Score |
|------|-------|
| `amsi.go` | 0.2 |
| `app.py` | 0.0 |
| `install.sh` | 0.0 |

## Strongest Connections

- No cross-community connections recorded.

## Navigation Tips

- Obsidian Graph View works: every community page links back here.
- `connections.json` is machine-readable for GraphRAG pipelines.
- `REPORT.md` states what was extracted vs inferred and current limits.
- Regenerate offline: `readmenator . --rebuild` (no network, no tokens).
