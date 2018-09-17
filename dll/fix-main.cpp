#include "fix-main.hpp"
#include "_cgo_export.h"

BOOL WINAPI DllMain(HINSTANCE hinstDLL, DWORD fdwReason, LPVOID lpReserved) {
	Attach();
    return TRUE
}

EXPORT void hello(void) {
    printf ("Hello\n");
}
