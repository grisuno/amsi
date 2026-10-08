# root

*Community 0 | 3 files | cohesion 1.00*

## Definition

This community groups 3 file(s) rooted at `root` with dominant language go (cohesion 1.00). Central symbols: `main`, `patchAMSI`. Core file: `amsi.go` (2 symbols). Documented purpose: Autor: Gris Iscomeback Correo electrónico: grisiscomeback[at]gmail[dot]com Fecha de creación: xx/xx/xxxx Licencia: GPL v3  Descripción:.

## Files

| File | Language | Layer | Symbols | Doc |
|------|----------|-------|---------|-----|
| `amsi.go` | go | utility | 2 | no |
| `app.py` | py | utility | 0 | yes |
| `install.sh` | sh | utility | 0 | no |

## Key Symbols

- `patchAMSI` (function, `amsi.go:11`) `func patchAMSI(`
- `main` (function, `amsi.go:79`) `func main(`

## Internal vs External Edges

- Internal resolved imports (EXTRACTED): 0
- Cross-boundary resolved imports (EXTRACTED): 0

## Connections

- No cross-community bridges recorded. This community is self-contained.

## Risks

- No scoped security, taint, cycle, or layer risks.

## Open Questions

- Why do 2 file(s) lack file-level docs (e.g. `amsi.go`)? What purpose do they serve?
- What would break if the most connected file in root changed?
- Should root be split, given cohesion 1.00?

## Sources

- `amsi.go`
- `app.py`
- `install.sh`
