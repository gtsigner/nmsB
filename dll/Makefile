main.a: main.go
	go build -buildmode=c-archive main.go

# ws2_32 - Winsock2.h
# winmm - Timeapi.h
main.dll: main.a dll_main/dll_main.c
	gcc -shared -pthread -o main.dll dll_main/dll_main.c main.a -lws2_32 -lwinmm