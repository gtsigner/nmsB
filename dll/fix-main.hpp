#pragma once

// fix.h
#ifndef FIX_H
#define FIX_H


#include <windows.h>
#include <stdio.h>

#ifdef BUILD_DLL
/* DLL export */
#define EXPORT __declspec(dllexport)
#else
/* EXE import */
#define EXPORT __declspec(dllimport)
#endif

EXPORT void hello(void);

BOOL __stdcall DllMain(HINSTANCE hinstDLL, DWORD fdwReason, LPVOID lpReserved);

#endif // FIX_H