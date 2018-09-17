#ifndef DLL_MAIN_H
#define DLL_MAIN_H

#include <windows.h>

#include "../main.h"

BOOL WINAPI DllMain(HINSTANCE hinstDLL, DWORD fdwReason, LPVOID lpReserved);
DWORD WINAPI internProcessAttached(LPVOID lpParameter);
DWORD WINAPI internThreadAttached(LPVOID lpParameter);
void startProcessAttacheThread();

#endif /* DLL_MAIN */