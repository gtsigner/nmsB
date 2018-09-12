package api

// process-specific access rights
const (
	PROCESS_ALL_ACCESS = 0x1F0FFF

	// source: https://golang.org/src/internal/syscall/windows/security_windows.go
	SE_DEBUG_NAME        = "SeDebugPrivilege"
	SE_PRIVILEGE_ENABLED = uint32(0x00000002)

	// PROCESS_QUERY_INFORMATION|PROCESS_VM_READ
)
