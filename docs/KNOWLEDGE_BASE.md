# Polyglot Codebase Knowledge Graph

> Generated offline by **readmenator**. Supports C, C++, Python, Go, Rust, JS/TS, Java, C#, Shell, PHP, Dart, GDScript, Nim, ASM.
> No LLMs. No tokens. Pure static analysis.

**Total Files Parsed:** 3 | **Total Symbols Extracted:** 2 | **Total Imports:** 5

## Structural Knowledge Map
```mermaid
graph TD
    classDef mod fill:#1e1e1e,stroke:#ff6666,stroke-width:2px,color:#fff;
    classDef cls fill:#2d2d2d,stroke:#4ec9b0,stroke-width:2px,color:#fff;
    classDef fn fill:#333,stroke:#dcdcaa,stroke-width:1px,color:#dcdcaa;
    classDef ext fill:#111,stroke:#666,stroke-dasharray: 5 5,color:#aaa;
    amsi_go["amsi.go (go)"]
    class amsi_go mod;
    amsi_go_patchAMSI["patchAMSI"]
    class amsi_go_patchAMSI fn;
    amsi_go --> amsi_go_patchAMSI
    amsi_go_main["main"]
    class amsi_go_main fn;
    amsi_go --> amsi_go_main
    app_py["app.py (py)"]
    class app_py mod;
    install_sh["install.sh (sh)"]
    class install_sh mod;
    ext_fmt["fmt"]
    class ext_fmt ext;
    amsi_go -.->|imports| ext_fmt
    ext_syscall["syscall"]
    class ext_syscall ext;
    amsi_go -.->|imports| ext_syscall
    ext_unsafe["unsafe"]
    class ext_unsafe ext;
    amsi_go -.->|imports| ext_unsafe
    ext_golang_org_x_sys_windows["windows"]
    class ext_golang_org_x_sys_windows ext;
    amsi_go -.->|imports| ext_golang_org_x_sys_windows
    ext_os["os"]
    class ext_os ext;
    app_py -.->|imports| ext_os
```

---

## Architecture Reference

### GO (1 files)

#### `amsi.go`
**Path:** `amsi.go`

**Functions:**
- `patchAMSI` (line 11)
- `main` (line 79)

### PY (1 files)

#### `app.py`
**Path:** `app.py`

*No symbols extracted*

### SH (1 files)

#### `install.sh`
**Path:** `install.sh`

*No symbols extracted*
