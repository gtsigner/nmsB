
#include "dll_main.h"

BOOL WINAPI DllMain(HINSTANCE hinstDLL, DWORD fdwReason, LPVOID lpReserved)
{
    switch (fdwReason)
    {
    case DLL_PROCESS_ATTACH:
        startProcessAttacheThread();
        break;
    case DLL_THREAD_ATTACH:
        //DisableThreadLibraryCalls(hinstDLL);
        //CreateThread(NULL, 0, (LPTHREAD_START_ROUTINE)internThreadAttached, NULL, 0, NULL);
        break;

    case DLL_THREAD_DETACH:
        // Do thread-specific cleanup.
        break;

    case DLL_PROCESS_DETACH:
        // Perform any necessary cleanup.
        break;
    }
    return TRUE; // Successful DLL_PROCESS_ATTACH.
}

void startProcessAttacheThread(){
    HANDLE thread = CreateThread(NULL, 0, (LPTHREAD_START_ROUTINE)&internProcessAttached, NULL, 0, NULL);
    CloseHandle(thread);
}

DWORD WINAPI internProcessAttached(LPVOID lpParameter)
{
    ProcessAttached();
    return 1;
}

DWORD WINAPI internThreadAttached(LPVOID lpParameter)
{
    ThreadAttached();
    return 1;
}
